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
	fields := actionableFieldDiffs(res.FieldDiffs, opt.tagFilter)

	// Sections with nothing to do are skipped entirely; the remaining ones are
	// numbered sequentially so the headings stay 1..N with no gaps.
	n := 0
	if len(renames) > 0 {
		n++
		instrRenames(w, n, renames)
	}
	if len(missing) > 0 {
		n++
		instrMissing(w, n, doc, missing, refToGo)
	}
	if len(fields) > 0 {
		n++
		instrStructFields(w, n, fields)
	}
	if len(res.Stale) > 0 {
		n++
		instrStale(w, n, res.Stale)
	}
	if n == 0 {
		fmt.Fprintln(w, "_Nothing to do — client is in sync with the spec._")
	}
}

func instrRenames(w io.Writer, n int, renames []Rename) {
	fmt.Fprintf(w, "## %d. Renames to fix (%d)\n\n", n, len(renames))
	fmt.Fprint(w, "The client path constant points at the old path; update it to the spec path. No other change is needed.\n\n")
	for _, r := range renames {
		fmt.Fprintf(w, "- **%s** — `%s`\n", r.Spec.OperationID, r.Spec.Verb)
		fmt.Fprintf(w, "  - client: `%s:%d` (`%s.%s`)\n", r.Client.File, r.Client.Line, r.Client.RecvType, r.Client.Method)
		fmt.Fprintf(w, "  - change path: `%s` → `%s`\n", r.Client.Norm.Raw, r.Spec.Norm.Raw)
		fmt.Fprintf(w, "  - matched by: %s\n", r.Reason)
	}
	fmt.Fprintln(w)
}

func instrMissing(w io.Writer, n int, doc *openAPI, missing []SpecOp, refToGo map[string]string) {
	fmt.Fprintf(w, "## %d. Missing endpoints to add (%d)\n\n", n, len(missing))

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

// actionableFieldDiffs returns the field diffs that add fields to a Go struct,
// honoring the tag filter. These are the only struct-field diffs worth emitting.
func actionableFieldDiffs(diffs []FieldDiff, tagFilter string) []FieldDiff {
	var actionable []FieldDiff
	for _, d := range diffs {
		if tagFilter != "" && !hasTag(d.Op, tagFilter) {
			continue
		}
		if len(d.MissingInGo) > 0 || len(d.MissingReqd) > 0 {
			actionable = append(actionable, d)
		}
	}
	return actionable
}

func instrStructFields(w io.Writer, n int, actionable []FieldDiff) {
	fmt.Fprintf(w, "## %d. Struct fields to add (%d)\n\n", n, len(actionable))
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

func instrStale(w io.Writer, n int, stale []ClientOp) {
	fmt.Fprintf(w, "## %d. Stale to investigate / remove (%d)\n\n", n, len(stale))
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
