package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// renderInstructions emits a Markdown brief an AI agent (or human) can act on: renames to
// fix, missing endpoints to add (with typed request/response fields), struct fields to
// add, and genuinely-stale ops to investigate. It honors the -tag filter.
func renderInstructions(w io.Writer, doc *openAPI, res *MatchResult, refToGo map[string]string, opt options) {
	missing := filterMissingByTag(res.Missing, opt.tagFilter)
	renames := filterRenamesByTag(res.Renames, opt.tagFilter)

	fmt.Fprintln(w, "# specsync — sync instructions")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Spec: `%s`\n\n", opt.specPath)
	if opt.tagFilter != "" {
		fmt.Fprintf(w, "Filtered to tag **%s**.\n\n", opt.tagFilter)
	}
	pct := 0.0
	if res.SpecOpCount > 0 {
		pct = 100 * float64(res.CoveredCount) / float64(res.SpecOpCount)
	}
	fmt.Fprintf(w, "Coverage: %d/%d (%.1f%%) — missing=%d, renamed=%d, stale=%d.\n\n",
		res.CoveredCount, res.SpecOpCount, pct, len(res.Missing), len(res.Renames), len(res.Stale))

	instrRenames(w, renames)
	instrMissing(w, doc, missing, refToGo)
	instrStructFields(w, res.FieldDiffs, opt.tagFilter)
	instrStale(w, res.Stale)
}

func instrRenames(w io.Writer, renames []Rename) {
	fmt.Fprintf(w, "## 1. Renames to fix (%d)\n\n", len(renames))
	if len(renames) == 0 {
		fmt.Fprint(w, "_None._\n\n")
		return
	}
	fmt.Fprint(w, "The client path constant points at the old path; update it to the spec path. No other change is needed.\n\n")
	for _, r := range renames {
		fmt.Fprintf(w, "- **%s** — `%s`\n", r.Spec.OperationID, r.Spec.Verb)
		fmt.Fprintf(w, "  - client: `%s:%d` (`%s.%s`)\n", r.Client.File, r.Client.Line, r.Client.RecvType, r.Client.Method)
		fmt.Fprintf(w, "  - change path: `%s` → `%s`\n", r.Client.Norm.Raw, r.Spec.Norm.Raw)
		fmt.Fprintf(w, "  - matched by: %s\n", r.Reason)
	}
	fmt.Fprintln(w)
}

func instrMissing(w io.Writer, doc *openAPI, missing []SpecOp, refToGo map[string]string) {
	fmt.Fprintf(w, "## 2. Missing endpoints to add (%d)\n\n", len(missing))
	if len(missing) == 0 {
		fmt.Fprint(w, "_None._\n\n")
		return
	}

	byTag := map[string][]SpecOp{}
	for _, op := range missing {
		t := op.PrimaryTag()
		byTag[t] = append(byTag[t], op)
	}
	tags := make([]string, 0, len(byTag))
	for t := range byTag {
		tags = append(tags, t)
	}
	sort.Strings(tags)

	for _, t := range tags {
		fmt.Fprintf(w, "### %s\n\n", t)
		for _, op := range byTag[t] {
			fmt.Fprintf(w, "- **%s** — `%s %s`\n", op.OperationID, op.Verb, op.RawPath)
			if op.Summary != "" {
				fmt.Fprintf(w, "  - summary: %s\n", op.Summary)
			}
			if params := pathParams(op.RawPath); len(params) > 0 {
				fmt.Fprintf(w, "  - path params: %s\n", strings.Join(params, ", "))
			}
			instrBody(w, doc, "request body", op.ReqBodySchema, refToGo)
			instrBody(w, doc, "response", op.SuccessSchema, refToGo)
		}
		fmt.Fprintln(w)
	}
}

// instrBody renders one request/response schema as a typed field list, or names the known
// Go struct when the schema maps to an existing entity.
func instrBody(w io.Writer, doc *openAPI, label string, s *Schema, refToGo map[string]string) {
	if s == nil {
		return
	}
	name, fields, isList, ok := doc.resolveFields(s, refToGo)
	if !ok {
		return
	}
	prefix := ""
	if isList {
		prefix = "[]"
	}
	if name != "" && name != "(inline)" {
		if g, found := refToGo[name]; found {
			fmt.Fprintf(w, "  - %s: `%s%s` (existing struct)\n", label, prefix, g)
			return
		}
	}
	fmt.Fprintf(w, "  - %s fields:\n", label)
	for _, f := range fields {
		fmt.Fprintf(w, "    - `%s`  %s%s\n", f.Name, f.GoType, reqNote(f))
	}
}

func reqNote(f FieldInfo) string {
	var extra []string
	if f.Required {
		extra = append(extra, "required")
	}
	if f.Note != "" {
		extra = append(extra, f.Note)
	}
	if len(extra) == 0 {
		return ""
	}
	return "  (" + strings.Join(extra, "; ") + ")"
}

func instrStructFields(w io.Writer, diffs []FieldDiff, tagFilter string) {
	if tagFilter != "" {
		var out []FieldDiff
		for _, d := range diffs {
			if hasTag(d.Op, tagFilter) {
				out = append(out, d)
			}
		}
		diffs = out
	}
	// Only diffs that add fields to a Go struct are actionable here.
	var actionable []FieldDiff
	for _, d := range diffs {
		if len(d.MissingInGo) > 0 || len(d.MissingReqd) > 0 {
			actionable = append(actionable, d)
		}
	}

	fmt.Fprintf(w, "## 3. Struct fields to add (%d)\n\n", len(actionable))
	if len(actionable) == 0 {
		fmt.Fprint(w, "_None._\n\n")
		return
	}
	fmt.Fprint(w, "Spec properties present on matched operations but absent from the Go struct.\n\n")
	for _, d := range actionable {
		conf := ""
		if d.LowConf {
			conf = fmt.Sprintf(" _(low-confidence: %s)_", d.LowReason)
		}
		fmt.Fprintf(w, "- **%s** (%s, via `%s %s`)%s\n", d.GoType, d.Kind, d.Op.Verb, d.Op.RawPath, conf)
		for _, f := range append(append([]FieldInfo{}, d.MissingReqd...), d.MissingInGo...) {
			fmt.Fprintf(w, "    - `%s %s` `json:\"%s\"`%s\n", goFieldName(f.Name), f.GoType, f.Name, reqNote(f))
		}
	}
	fmt.Fprintln(w)
}

func instrStale(w io.Writer, stale []ClientOp) {
	fmt.Fprintf(w, "## 4. Stale to investigate / remove (%d)\n\n", len(stale))
	if len(stale) == 0 {
		fmt.Fprint(w, "_None._\n\n")
		return
	}
	fmt.Fprint(w, "Client operations with no spec match and no rename pair — confirm whether the endpoint was removed.\n\n")
	for _, c := range stale {
		fmt.Fprintf(w, "- `%s %s` — `%s.%s` (`%s:%d`)\n", c.Verb, c.Norm.Raw, c.RecvType, c.Method, c.File, c.Line)
	}
	fmt.Fprintln(w)
}

// pathParams returns the {param} names in a raw spec path.
func pathParams(rawPath string) []string {
	var out []string
	for _, seg := range splitPath(rawPath) {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			out = append(out, strings.TrimSuffix(strings.TrimPrefix(seg, "{"), "}"))
		}
	}
	return out
}

// goFieldName converts a snake_case json name to an exported Go field name (best-effort).
func goFieldName(json string) string {
	parts := strings.Split(json, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		b.WriteString(p[1:])
	}
	if b.Len() == 0 {
		return json
	}
	return b.String()
}
