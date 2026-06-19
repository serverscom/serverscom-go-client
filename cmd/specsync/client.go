package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// clientModel is the result of statically analyzing the Go client package.
type clientModel struct {
	rawConsts  map[string]ast.Expr        // const name -> value expression (pre-fold)
	consts     map[string]string          // const name -> folded string value
	structDefs map[string]*ast.StructType // struct type name -> definition
	ops        []ClientOp
	unresolved []string // handler methods that look like API calls but didn't parse

	// helpers indexes path-building methods by "Recv.Method" so that a thin wrapper
	// which only delegates (e.g. ActivateOobPublicAccessFeature -> activateFeature)
	// can be resolved to a concrete path: the wrapper's literal arguments are
	// substituted into the helper's path-parameter wildcards.
	helpers   map[string]helperInfo
	opKeys    map[string]bool // "Recv.Method" that already produced a direct op
	delegated map[string]bool // helper keys a wrapper delegated to (their own op is dropped)
}

// helperInfo captures what a path-building method needs in order to be resolved from a
// delegating caller: the path template, and for each path wildcard the helper parameter
// that fills it (so a caller passing a literal there pins that segment).
type helperInfo struct {
	verb       string
	template   string
	isList     bool
	params     []string // flattened parameter names in signature order (ctx first)
	baseSubsts []string // the helper's own resolved substitutions (consts pinned, params "")
	fillBy     []string // fillBy[i] = parameter name filling path-wildcard i, or "" if not a param
}

// jsonFields returns the json field names of a struct, recursing one level into
// embedded (anonymous) structs. Fields tagged "-" or with no json tag are skipped.
func (m *clientModel) jsonFields(name string) []string {
	return m.collectFields(name, map[string]bool{})
}

func (m *clientModel) collectFields(name string, seen map[string]bool) []string {
	st, ok := m.structDefs[name]
	if !ok || seen[name] {
		return nil
	}
	seen[name] = true

	var out []string
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 { // embedded field
			if emb := baseTypeName(f.Type); emb != "" {
				out = append(out, m.collectFields(emb, seen)...)
			}
			continue
		}
		tag := jsonTagName(f.Tag)
		if tag == "" || tag == "-" {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// parseClient analyzes every non-test *.go file in pkgDir.
func parseClient(pkgDir string) (*clientModel, error) {
	m := &clientModel{
		rawConsts:  map[string]ast.Expr{},
		consts:     map[string]string{},
		structDefs: map[string]*ast.StructType{},
		helpers:    map[string]helperInfo{},
		opKeys:     map[string]bool{},
		delegated:  map[string]bool{},
	}
	fset := token.NewFileSet()

	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("read pkg dir %q: %w", pkgDir, err)
	}

	type parsedFile struct {
		path string
		file *ast.File
	}
	var files []parsedFile
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") {
			continue
		}
		if strings.HasSuffix(name, "_test.go") || name == "test.go" {
			continue
		}
		full := filepath.Join(pkgDir, name)
		f, err := parser.ParseFile(fset, full, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", full, err)
		}
		files = append(files, parsedFile{full, f})
	}

	// Pass 1: collect const expressions and struct definitions.
	for _, pf := range files {
		m.collectDecls(pf.file)
	}
	m.foldConsts()

	// Pass 2: extract operations from handler methods, and register every
	// path-building method as a potential delegation target.
	for _, pf := range files {
		m.extractOps(pf.file, pf.path, fset)
	}

	// Pass 3: resolve thin wrappers that only delegate to a path-building helper,
	// pinning the helper's path wildcards with the wrapper's literal arguments. Then
	// drop the helper's own (wildcard) op so it no longer spuriously covers every
	// literal a wrapper enumerates.
	for _, pf := range files {
		m.resolveDelegations(pf.file, pf.path, fset)
	}
	if len(m.delegated) > 0 {
		kept := m.ops[:0]
		for _, op := range m.ops {
			if m.delegated[op.RecvType+"."+op.Method] && !ast.IsExported(op.Method) {
				continue
			}
			kept = append(kept, op)
		}
		m.ops = kept
	}

	sort.Slice(m.ops, func(i, j int) bool {
		if m.ops[i].Norm.Raw != m.ops[j].Norm.Raw {
			return m.ops[i].Norm.Raw < m.ops[j].Norm.Raw
		}
		return m.ops[i].Verb < m.ops[j].Verb
	})
	return m, nil
}

func (m *clientModel) collectDecls(f *ast.File) {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		switch gd.Tok {
		case token.CONST:
			for _, spec := range gd.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for i, name := range vs.Names {
					if i < len(vs.Values) {
						m.rawConsts[name.Name] = vs.Values[i]
					}
				}
			}
		case token.TYPE:
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if st, ok := ts.Type.(*ast.StructType); ok {
					m.structDefs[ts.Name.Name] = st
				}
			}
		}
	}
}

// foldConsts evaluates string-valued const expressions to a fixpoint, so chained
// concatenations (A = "x"; B = A + "/%s"; C = B + "/nodes") all resolve.
func (m *clientModel) foldConsts() {
	for {
		changed := false
		for name, expr := range m.rawConsts {
			if _, done := m.consts[name]; done {
				continue
			}
			if v, ok := m.evalStr(expr); ok {
				m.consts[name] = v
				changed = true
			}
		}
		if !changed {
			return
		}
	}
}

// evalStr folds a string expression: literals, references to known consts, "+"
// concatenation, and a "*.baseURL" selector (which contributes nothing to the path
// template since the client base URL already carries the version prefix).
func (m *clientModel) evalStr(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			if v, err := strconv.Unquote(e.Value); err == nil {
				return v, true
			}
		}
	case *ast.Ident:
		if v, ok := m.consts[e.Name]; ok {
			return v, true
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			l, lok := m.evalStr(e.X)
			r, rok := m.evalStr(e.Y)
			if lok && rok {
				return l + r, true
			}
		}
	case *ast.SelectorExpr:
		if e.Sel.Name == "baseURL" {
			return "", true
		}
	case *ast.ParenExpr:
		return m.evalStr(e.X)
	}
	return "", false
}

func (m *clientModel) extractOps(f *ast.File, file string, fset *token.FileSet) {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 || fn.Body == nil {
			continue
		}
		recv := recvTypeName(fn.Recv.List[0].Type)
		if !strings.HasSuffix(recv, "Handler") || recv == "CollectionHandler" {
			continue
		}
		m.extractMethod(fn, recv, file, fset)
	}
}

func (m *clientModel) extractMethod(fn *ast.FuncDecl, recv, file string, fset *token.FileSet) {
	locals := map[string]ast.Expr{} // single-ident assignments: name -> RHS
	var reqCall *ast.CallExpr       // buildAndExecRequest[WithResponse] call
	var collPathExpr ast.Expr       // NewCollection's path argument
	collFound := false

	ast.Inspect(fn.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) == 1 && len(node.Rhs) == 1 {
				if id, ok := node.Lhs[0].(*ast.Ident); ok {
					locals[id.Name] = node.Rhs[0]
				}
			}
		case *ast.CallExpr:
			switch fexpr := node.Fun.(type) {
			case *ast.SelectorExpr:
				if reqCall == nil &&
					(fexpr.Sel.Name == "buildAndExecRequest" || fexpr.Sel.Name == "buildAndExecRequestWithResponse") {
					reqCall = node
				}
			case *ast.IndexExpr:
				if id, ok := fexpr.X.(*ast.Ident); ok && id.Name == "NewCollection" && !collFound {
					collFound = true
					collPathExpr = secondArg(node)
				}
			case *ast.IndexListExpr:
				if id, ok := fexpr.X.(*ast.Ident); ok && id.Name == "NewCollection" && !collFound {
					collFound = true
					collPathExpr = secondArg(node)
				}
			}
		}
		return true
	})

	var (
		verb      string
		template  string
		valueArgs []ast.Expr
		ok        bool
		isList    bool
	)

	switch {
	case reqCall != nil:
		if len(reqCall.Args) >= 3 {
			if lit, isLit := reqCall.Args[1].(*ast.BasicLit); isLit && lit.Kind == token.STRING {
				verb, _ = strconv.Unquote(lit.Value)
			}
			template, valueArgs, ok = m.resolvePath(reqCall.Args[2], locals)
		}
	case collFound:
		verb = "GET"
		isList = true
		template, valueArgs, ok = m.resolvePath(collPathExpr, locals)
	}

	line := fset.Position(fn.Pos()).Line

	if !ok || verb == "" || template == "" {
		// Flag only methods that clearly attempt an API call but didn't fully parse.
		if reqCall != nil || collFound {
			m.unresolved = append(m.unresolved,
				fmt.Sprintf("%s.%s (%s:%d)", recv, fn.Name.Name, file, line))
		}
		return
	}

	substs := m.substsOf(valueArgs)
	key := recv + "." + fn.Name.Name
	m.opKeys[key] = true
	m.helpers[key] = helperInfo{
		verb:       verb,
		template:   template,
		isList:     isList,
		params:     paramNames(fn),
		baseSubsts: substs,
		fillBy:     fillByParams(valueArgs, fn),
	}
	m.ops = append(m.ops, ClientOp{
		Verb:        verb,
		PathValue:   template,
		Substs:      substs,
		Norm:        normalizeClientPath(template, substs),
		RecvType:    recv,
		Method:      fn.Name.Name,
		File:        file,
		Line:        line,
		OperationID: operationID(fn.Doc),
		ReturnType:  returnEntity(fn),
		InputType:   m.inputStruct(fn),
		IsList:      isList,
	})
}

// resolveDelegations synthesizes an op for each handler method that builds no path of
// its own but delegates to a registered path-building helper on the same receiver. The
// wrapper's literal arguments are substituted into the helper's path-parameter wildcards,
// so e.g. ActivateOobPublicAccessFeature resolves to .../features/oob_public_access/activate
// instead of inheriting the helper's catch-all wildcard.
func (m *clientModel) resolveDelegations(f *ast.File, file string, fset *token.FileSet) {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 || fn.Body == nil {
			continue
		}
		recv := recvTypeName(fn.Recv.List[0].Type)
		if !strings.HasSuffix(recv, "Handler") || recv == "CollectionHandler" {
			continue
		}
		if m.opKeys[recv+"."+fn.Name.Name] {
			continue // already extracted a direct op
		}
		recvVar := recvVarName(fn)
		if recvVar == "" {
			continue
		}

		var (
			call *ast.CallExpr
			info helperInfo
			hkey string
		)
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok || call != nil {
				return call == nil
			}
			sel, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			id, ok := sel.X.(*ast.Ident)
			if !ok || id.Name != recvVar {
				return true
			}
			if hi, found := m.helpers[recv+"."+sel.Sel.Name]; found {
				call, info, hkey = ce, hi, recv+"."+sel.Sel.Name
			}
			return call == nil
		})
		if call == nil {
			continue
		}

		substs := make([]string, len(info.baseSubsts))
		copy(substs, info.baseSubsts)
		for i, pname := range info.fillBy {
			if pname == "" {
				continue // const-pinned or runtime in the helper; inherit baseSubsts[i]
			}
			idx := indexOf(info.params, pname)
			if idx < 0 || idx >= len(call.Args) {
				continue
			}
			if lit, ok := call.Args[idx].(*ast.BasicLit); ok && lit.Kind == token.STRING {
				if v, err := strconv.Unquote(lit.Value); err == nil {
					substs[i] = v
				}
			}
		}

		m.delegated[hkey] = true
		m.ops = append(m.ops, ClientOp{
			Verb:        info.verb,
			PathValue:   info.template,
			Substs:      substs,
			Norm:        normalizeClientPath(info.template, substs),
			RecvType:    recv,
			Method:      fn.Name.Name,
			File:        file,
			Line:        fset.Position(fn.Pos()).Line,
			OperationID: operationID(fn.Doc),
			ReturnType:  returnEntity(fn),
			InputType:   m.inputStruct(fn),
			IsList:      info.isList,
		})
	}
}

// operationID extracts the spec operationId from a method's doc comment, which the
// client annotates as ".../operation/<ID>". Returns "" when absent.
func operationID(doc *ast.CommentGroup) string {
	if doc == nil {
		return ""
	}
	const marker = "operation/"
	text := doc.Text()
	i := strings.Index(text, marker)
	if i < 0 {
		return ""
	}
	rest := text[i+len(marker):]
	end := 0
	for end < len(rest) && isOpIDChar(rest[end]) {
		end++
	}
	return rest[:end]
}

func isOpIDChar(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' || b == '_'
}

// resolvePath reduces a path expression to a format template and its substitution
// arguments. It traces local variables, folds string concatenations (including the
// baseURL+const form), and unpacks buildURL/buildPath calls.
func (m *clientModel) resolvePath(expr ast.Expr, locals map[string]ast.Expr) (template string, valueArgs []ast.Expr, ok bool) {
	switch e := expr.(type) {
	case *ast.CallExpr:
		if sel, isSel := e.Fun.(*ast.SelectorExpr); isSel && len(e.Args) > 0 {
			switch sel.Sel.Name {
			case "buildURL", "buildPath":
				tmpl, tok := m.evalStr(e.Args[0])
				return tmpl, valueArgsOf(e), tok
			case "applyParams":
				// applyParams(endpointURL, params) — resolve the wrapped URL.
				return m.resolvePath(e.Args[0], locals)
			}
		}
	case *ast.Ident:
		if v, isConst := m.consts[e.Name]; isConst {
			return v, nil, true
		}
		if rhs, isLocal := locals[e.Name]; isLocal {
			return m.resolvePath(rhs, locals)
		}
	case *ast.BinaryExpr, *ast.BasicLit, *ast.SelectorExpr, *ast.ParenExpr:
		if v, vok := m.evalStr(expr); vok {
			return v, nil, true
		}
	}
	return "", nil, false
}

// valueArgsOf returns the substitution arguments of a buildURL/buildPath call,
// flattening a spread composite literal: f(const, []interface{}{a, b}...) -> [a, b].
func valueArgsOf(call *ast.CallExpr) []ast.Expr {
	rest := call.Args[1:]
	if call.Ellipsis != token.NoPos && len(rest) == 1 {
		if cl, ok := rest[0].(*ast.CompositeLit); ok {
			return cl.Elts
		}
	}
	return rest
}

// substsOf resolves positional path arguments: a value arg that is a known string
// const becomes a literal substitution; anything else (a runtime id) stays "".
func (m *clientModel) substsOf(valueArgs []ast.Expr) []string {
	substs := make([]string, len(valueArgs))
	for i, a := range valueArgs {
		if id, ok := a.(*ast.Ident); ok {
			if v, found := m.consts[id.Name]; found {
				substs[i] = v
			}
		}
	}
	return substs
}

// paramNames returns the method's parameter names flattened in signature order,
// expanding grouped declarations ("serverID, feature string" -> two names). An unnamed
// parameter contributes "" to keep positions aligned with the call's argument list.
func paramNames(fn *ast.FuncDecl) []string {
	if fn.Type.Params == nil {
		return nil
	}
	var names []string
	for _, p := range fn.Type.Params.List {
		if len(p.Names) == 0 {
			names = append(names, "")
			continue
		}
		for _, n := range p.Names {
			names = append(names, n.Name)
		}
	}
	return names
}

// fillByParams maps each path wildcard to the parameter name that supplies it: a value
// arg that is a plain parameter ident yields that name; anything else (a const, a runtime
// expression) yields "" so the helper's own resolution is kept.
func fillByParams(valueArgs []ast.Expr, fn *ast.FuncDecl) []string {
	params := map[string]bool{}
	for _, n := range paramNames(fn) {
		if n != "" {
			params[n] = true
		}
	}
	out := make([]string, len(valueArgs))
	for i, a := range valueArgs {
		if id, ok := a.(*ast.Ident); ok && params[id.Name] {
			out[i] = id.Name
		}
	}
	return out
}

// recvVarName returns the receiver variable name (e.g. "h" in "func (h *X) ..."), or ""
// when the receiver is unnamed.
func recvVarName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 || len(fn.Recv.List[0].Names) == 0 {
		return ""
	}
	return fn.Recv.List[0].Names[0].Name
}

func indexOf(ss []string, want string) int {
	for i, s := range ss {
		if s == want {
			return i
		}
	}
	return -1
}

func secondArg(call *ast.CallExpr) ast.Expr {
	if len(call.Args) >= 2 {
		return call.Args[1]
	}
	return nil
}

// returnEntity returns the first non-error result type, unwrapped to its base name
// (handles *T, []T, Collection[T]).
func returnEntity(fn *ast.FuncDecl) string {
	if fn.Type.Results == nil {
		return ""
	}
	for _, r := range fn.Type.Results.List {
		bt := baseTypeName(r.Type)
		if bt == "" || bt == "error" {
			continue
		}
		return bt
	}
	return ""
}

// inputStruct returns the last parameter whose type is a struct defined in the package.
func (m *clientModel) inputStruct(fn *ast.FuncDecl) string {
	if fn.Type.Params == nil {
		return ""
	}
	last := ""
	for _, p := range fn.Type.Params.List {
		bt := baseTypeName(p.Type)
		if bt == "" {
			continue
		}
		if _, ok := m.structDefs[bt]; ok {
			last = bt
		}
	}
	return last
}

// baseTypeName unwraps *T, []T and generic T[Arg] to the meaningful element name.
// For an index expression (Collection[Network]) it returns the type argument.
func baseTypeName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return baseTypeName(t.X)
	case *ast.ArrayType:
		return baseTypeName(t.Elt)
	case *ast.SelectorExpr:
		return t.Sel.Name
	case *ast.IndexExpr:
		return baseTypeName(t.Index)
	case *ast.IndexListExpr:
		if len(t.Indices) > 0 {
			return baseTypeName(t.Indices[0])
		}
	}
	return ""
}

// recvTypeName returns the base (generic) name of a receiver type: for *Handler it
// is "Handler"; for *CollectionHandler[K] it is "CollectionHandler".
func recvTypeName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.StarExpr:
		return recvTypeName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.IndexExpr:
		return recvTypeName(t.X)
	case *ast.IndexListExpr:
		return recvTypeName(t.X)
	}
	return ""
}

func jsonTagName(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	unq, err := strconv.Unquote(tag.Value)
	if err != nil {
		return ""
	}
	v := reflect.StructTag(unq).Get("json")
	if v == "" {
		return ""
	}
	return strings.Split(v, ",")[0]
}
