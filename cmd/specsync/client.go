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
		f, err := parser.ParseFile(fset, full, nil, 0)
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

	// Pass 2: extract operations from handler methods.
	for _, pf := range files {
		m.extractOps(pf.file, pf.path, fset)
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
	m.ops = append(m.ops, ClientOp{
		Verb:       verb,
		PathValue:  template,
		Substs:     substs,
		Norm:       normalizeClientPath(template, substs),
		RecvType:   recv,
		Method:     fn.Name.Name,
		File:       file,
		Line:       line,
		ReturnType: returnEntity(fn),
		InputType:  m.inputStruct(fn),
		IsList:     isList,
	})
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
