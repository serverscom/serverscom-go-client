package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

func hasTag(op SpecOp, tag string) bool {
	for _, t := range op.Tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

func filterMissingByTag(missing []SpecOp, tag string) []SpecOp {
	if tag == "" {
		return missing
	}
	var out []SpecOp
	for _, m := range missing {
		if hasTag(m, tag) {
			out = append(out, m)
		}
	}
	return out
}

func renderText(w io.Writer, res *MatchResult, cm *clientModel, opt options) {
	fmt.Fprintln(w, "specsync — OpenAPI ↔ Go client coverage")
	if opt.tagFilter != "" {
		fmt.Fprintf(w, "(filtered to tag %q)\n", opt.tagFilter)
	}
	fmt.Fprintln(w)

	renderMissing(w, filterMissingByTag(res.Missing, opt.tagFilter))
	renderRenames(w, res.Renames, opt.tagFilter)
	renderStale(w, res.Stale, opt.tagFilter)
	renderFieldDiffs(w, res.FieldDiffs, opt.tagFilter)
	renderSummary(w, res, opt.tagFilter)

	if opt.debug {
		renderDebug(w, res, cm)
	}
}

func filterRenamesByTag(renames []Rename, tag string) []Rename {
	if tag == "" {
		return renames
	}
	var out []Rename
	for _, r := range renames {
		if hasTag(r.Spec, tag) {
			out = append(out, r)
		}
	}
	return out
}

func renderRenames(w io.Writer, renames []Rename, tagFilter string) {
	renames = filterRenamesByTag(renames, tagFilter)
	fmt.Fprintf(w, "B. RENAMED / MOVED — same operation, path changed in the spec (%d)\n", len(renames))
	if len(renames) == 0 {
		fmt.Fprintln(w, "   (none)")
		fmt.Fprintln(w)
		return
	}
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	for _, r := range renames {
		fmt.Fprintf(tw, "   %s\t%s → %s\t%s\t%s.%s\t%s:%d\n",
			r.Spec.Verb, r.Client.Norm.Raw, r.Spec.Norm.Raw,
			r.Spec.OperationID, r.Client.RecvType, r.Client.Method, r.Client.File, r.Client.Line)
	}
	tw.Flush()
	fmt.Fprintln(w)
}

func renderMissing(w io.Writer, missing []SpecOp) {
	fmt.Fprintf(w, "A. MISSING — in spec, not implemented in client (%d)\n", len(missing))
	if len(missing) == 0 {
		fmt.Fprintln(w, "   (none)")
		fmt.Fprintln(w)
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
		fmt.Fprintf(w, "   [%s]\n", t)
		tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
		for _, op := range byTag[t] {
			summary := op.Summary
			if summary != "" {
				summary = " — " + summary
			}
			fmt.Fprintf(tw, "      %s\t%s\t%s%s\n", op.Verb, op.RawPath, op.OperationID, summary)
		}
		tw.Flush()
	}
	fmt.Fprintln(w)
}

func renderStale(w io.Writer, stale []ClientOp, tagFilter string) {
	fmt.Fprintf(w, "C. STALE — in client, not found in spec and not a rename (%d)\n", len(stale))
	if tagFilter != "" {
		fmt.Fprintln(w, "   (note: stale detection is global; not affected by -tag)")
	}
	if len(stale) == 0 {
		fmt.Fprintln(w, "   (none)")
		fmt.Fprintln(w)
		return
	}
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	for _, c := range stale {
		fmt.Fprintf(tw, "   %s\t%s\t%s.%s\t%s:%d\n",
			c.Verb, c.PathValue, c.RecvType, c.Method, c.File, c.Line)
	}
	tw.Flush()
	fmt.Fprintln(w)
}

func renderFieldDiffs(w io.Writer, diffs []FieldDiff, tagFilter string) {
	if tagFilter != "" {
		var out []FieldDiff
		for _, d := range diffs {
			if hasTag(d.Op, tagFilter) {
				out = append(out, d)
			}
		}
		diffs = out
	}

	fmt.Fprintf(w, "D. FIELD DIFFS — advisory, matched operations only (%d)\n", len(diffs))
	if len(diffs) == 0 {
		fmt.Fprintln(w, "   (none)")
		fmt.Fprintln(w)
		return
	}
	for _, d := range diffs {
		conf := ""
		if d.LowConf {
			conf = fmt.Sprintf("  (low-confidence: %s)", d.LowReason)
		}
		fmt.Fprintf(w, "   %s %s  [%s]  %s ↔ %s%s\n",
			d.Op.Verb, d.Op.RawPath, d.Kind, d.GoType, d.SpecRef, conf)
		if len(d.MissingReqd) > 0 {
			fmt.Fprintf(w, "      ! required by spec, missing from %s: %s\n", d.GoType, fieldList(d.MissingReqd))
		}
		if len(d.MissingInGo) > 0 {
			fmt.Fprintf(w, "      + in spec, missing from %s: %s\n", d.GoType, fieldList(d.MissingInGo))
		}
		if len(d.MissingInSpec) > 0 {
			fmt.Fprintf(w, "      - in %s, not in spec: %s\n", d.GoType, strings.Join(d.MissingInSpec, ", "))
		}
	}
	fmt.Fprintln(w)
}

// fieldList renders typed fields as "name (GoType), name2 (GoType2 — note)".
func fieldList(fields []FieldInfo) string {
	parts := make([]string, 0, len(fields))
	for _, f := range fields {
		s := f.Name + " (" + f.GoType
		if f.Note != "" {
			s += " — " + f.Note
		}
		s += ")"
		parts = append(parts, s)
	}
	return strings.Join(parts, ", ")
}

func renderSummary(w io.Writer, res *MatchResult, tagFilter string) {
	if tagFilter != "" {
		fmt.Fprintf(w, "E. SUMMARY (tag: %s)\n", tagFilter)
	} else {
		fmt.Fprintln(w, "E. SUMMARY")
	}

	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	var total, covered int
	matched := false
	for _, c := range res.PerTag {
		if tagFilter != "" && !strings.EqualFold(c.Tag, tagFilter) {
			continue
		}
		matched = true
		total += c.Total
		covered += c.Covered
		mark := ""
		if c.Covered < c.Total {
			mark = "  <"
		}
		fmt.Fprintf(tw, "   %s\t%d/%d%s\n", c.Tag, c.Covered, c.Total, mark)
	}
	tw.Flush()

	fmt.Fprintln(w, "   "+strings.Repeat("-", 50))
	if tagFilter != "" {
		if !matched {
			fmt.Fprintf(w, "   (no spec operations tagged %q)\n", tagFilter)
		}
		pct := 0.0
		if total > 0 {
			pct = 100 * float64(covered) / float64(total)
		}
		// stale is reported globally: client ops carry no spec tag to filter by.
		fmt.Fprintf(w, "   spec ops=%d  covered=%d (%.1f%%)  missing=%d  renamed=%d  stale=%d (global)\n",
			total, covered, pct,
			len(filterMissingByTag(res.Missing, tagFilter)), len(filterRenamesByTag(res.Renames, tagFilter)), len(res.Stale))
		fmt.Fprintln(w)
		return
	}

	pct := 0.0
	if res.SpecOpCount > 0 {
		pct = 100 * float64(res.CoveredCount) / float64(res.SpecOpCount)
	}
	fmt.Fprintf(w, "   spec ops=%d  client ops=%d  covered=%d (%.1f%%)  missing=%d  renamed=%d  stale=%d\n",
		res.SpecOpCount, res.ClientOpCount, res.CoveredCount, pct,
		len(res.Missing), len(res.Renames), len(res.Stale))
	fmt.Fprintln(w)
}

func renderDebug(w io.Writer, res *MatchResult, cm *clientModel) {
	fmt.Fprintf(w, "DEBUG — unresolved handler methods (%d)\n", len(cm.unresolved))
	for _, u := range cm.unresolved {
		fmt.Fprintf(w, "   %s\n", u)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "DEBUG — coverage detail (spec op → covering client const)")
	keys := make([]string, 0, len(res.CoveredBy))
	for k := range res.CoveredBy {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	for _, k := range keys {
		var by []string
		for _, c := range res.CoveredBy[k] {
			by = append(by, fmt.Sprintf("%s.%s", c.RecvType, c.Method))
		}
		fmt.Fprintf(tw, "   %s\t%s\n", k, strings.Join(by, ", "))
	}
	tw.Flush()
	fmt.Fprintln(w)
}

// jsonView is the serializable shape emitted by -json.
type jsonView struct {
	Summary struct {
		SpecOps   int     `json:"spec_ops"`
		ClientOps int     `json:"client_ops"`
		Covered   int     `json:"covered"`
		Missing   int     `json:"missing"`
		Renamed   int     `json:"renamed"`
		Stale     int     `json:"stale"`
		Percent   float64 `json:"percent"`
	} `json:"summary"`
	PerTag  []Coverage      `json:"per_tag"`
	Missing []jsonSpecOp    `json:"missing"`
	Renames []jsonRename    `json:"renames"`
	Stale   []jsonClientOp  `json:"stale"`
	Fields  []jsonFieldDiff `json:"field_diffs"`
}

type jsonRename struct {
	Verb        string `json:"verb"`
	ClientPath  string `json:"client_path"`
	SpecPath    string `json:"spec_path"`
	OperationID string `json:"operation_id"`
	Method      string `json:"method"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Reason      string `json:"reason"`
}

type jsonField struct {
	Name     string `json:"name"`
	GoType   string `json:"go_type"`
	Required bool   `json:"required,omitempty"`
	Note     string `json:"note,omitempty"`
}

type jsonFieldDiff struct {
	Verb          string      `json:"verb"`
	Path          string      `json:"path"`
	Kind          string      `json:"kind"`
	GoType        string      `json:"go_type"`
	SpecSchema    string      `json:"spec_schema"`
	LowConfidence bool        `json:"low_confidence"`
	MissingInGo   []jsonField `json:"missing_in_go,omitempty"`
	MissingReqd   []jsonField `json:"missing_required,omitempty"`
	MissingInSpec []string    `json:"missing_in_spec,omitempty"`
}

func jsonFields(fields []FieldInfo) []jsonField {
	if len(fields) == 0 {
		return nil
	}
	out := make([]jsonField, 0, len(fields))
	for _, f := range fields {
		out = append(out, jsonField{Name: f.Name, GoType: f.GoType, Required: f.Required, Note: f.Note})
	}
	return out
}

type jsonSpecOp struct {
	Verb        string   `json:"verb"`
	Path        string   `json:"path"`
	OperationID string   `json:"operation_id"`
	Summary     string   `json:"summary"`
	Tags        []string `json:"tags"`
}

type jsonClientOp struct {
	Verb     string `json:"verb"`
	Path     string `json:"path"`
	Template string `json:"template"`
	Method   string `json:"method"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func renderJSON(w io.Writer, res *MatchResult) error {
	var v jsonView
	v.Summary.SpecOps = res.SpecOpCount
	v.Summary.ClientOps = res.ClientOpCount
	v.Summary.Covered = res.CoveredCount
	v.Summary.Missing = len(res.Missing)
	v.Summary.Renamed = len(res.Renames)
	v.Summary.Stale = len(res.Stale)
	if res.SpecOpCount > 0 {
		v.Summary.Percent = 100 * float64(res.CoveredCount) / float64(res.SpecOpCount)
	}
	v.PerTag = res.PerTag
	for _, m := range res.Missing {
		v.Missing = append(v.Missing, jsonSpecOp{m.Verb, m.RawPath, m.OperationID, m.Summary, m.Tags})
	}
	for _, r := range res.Renames {
		v.Renames = append(v.Renames, jsonRename{
			Verb: r.Spec.Verb, ClientPath: r.Client.Norm.Raw, SpecPath: r.Spec.Norm.Raw,
			OperationID: r.Spec.OperationID, Method: r.Client.RecvType + "." + r.Client.Method,
			File: r.Client.File, Line: r.Client.Line, Reason: r.Reason,
		})
	}
	for _, c := range res.Stale {
		v.Stale = append(v.Stale, jsonClientOp{c.Verb, c.Norm.Raw, c.PathValue, c.RecvType + "." + c.Method, c.File, c.Line})
	}
	for _, d := range res.FieldDiffs {
		v.Fields = append(v.Fields, jsonFieldDiff{
			Verb: d.Op.Verb, Path: d.Op.RawPath, Kind: d.Kind,
			GoType: d.GoType, SpecSchema: d.SpecRef, LowConfidence: d.LowConf,
			MissingInGo: jsonFields(d.MissingInGo), MissingReqd: jsonFields(d.MissingReqd), MissingInSpec: d.MissingInSpec,
		})
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
