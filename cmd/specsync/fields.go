package main

import "sort"

// computeFieldDiffs produces advisory field-level comparisons for matched operations.
// Each diff is traceable to a single operation and its resolved struct/schema pair.
// Diffs are deduplicated by (kind, Go struct, spec schema) since many operations share
// the same entity type.
func computeFieldDiffs(doc *openAPI, specOps []SpecOp, res *MatchResult, cm *clientModel) []FieldDiff {
	var diffs []FieldDiff
	seen := map[string]bool{}

	add := func(d FieldDiff, ok bool) {
		if !ok || d.empty() {
			return
		}
		key := d.Kind + "|" + d.GoType + "|" + d.SpecRef
		if seen[key] {
			return
		}
		seen[key] = true
		diffs = append(diffs, d)
	}

	for _, so := range specOps {
		covers := res.CoveredBy[so.matchKey()]
		if len(covers) == 0 {
			continue
		}
		if so.SuccessSchema != nil {
			if c, ok := pickClientWithReturn(covers, cm); ok {
				add(responseDiff(doc, so, c, cm))
			}
		}
		if so.ReqBodySchema != nil {
			if c, ok := pickClientWithInput(covers, cm); ok {
				add(requestDiff(doc, so, c, cm))
			}
		}
	}

	sort.Slice(diffs, func(i, j int) bool {
		if diffs[i].Op.RawPath != diffs[j].Op.RawPath {
			return diffs[i].Op.RawPath < diffs[j].Op.RawPath
		}
		return diffs[i].Kind < diffs[j].Kind
	})
	return diffs
}

func pickClientWithReturn(covers []ClientOp, cm *clientModel) (ClientOp, bool) {
	for _, c := range covers {
		if _, ok := cm.structDefs[c.ReturnType]; ok {
			return c, true
		}
	}
	return ClientOp{}, false
}

func pickClientWithInput(covers []ClientOp, cm *clientModel) (ClientOp, bool) {
	for _, c := range covers {
		if c.InputType == "" {
			continue
		}
		if _, ok := cm.structDefs[c.InputType]; ok {
			return c, true
		}
	}
	return ClientOp{}, false
}

func responseDiff(doc *openAPI, so SpecOp, c ClientOp, cm *clientModel) (FieldDiff, bool) {
	r, ok := doc.resolveProps(so.SuccessSchema)
	if !ok || len(r.props) == 0 {
		return FieldDiff{}, false
	}
	goFields := cm.jsonFields(c.ReturnType)
	if len(goFields) == 0 {
		return FieldDiff{}, false
	}
	return buildDiff(so, c, "response", c.ReturnType, r, goFields), true
}

func requestDiff(doc *openAPI, so SpecOp, c ClientOp, cm *clientModel) (FieldDiff, bool) {
	r, ok := doc.resolveProps(so.ReqBodySchema)
	if !ok || len(r.props) == 0 {
		return FieldDiff{}, false
	}
	goFields := cm.jsonFields(c.InputType)
	if len(goFields) == 0 {
		return FieldDiff{}, false
	}
	return buildDiff(so, c, "request", c.InputType, r, goFields), true
}

// buildDiff compares a resolved schema property set against a struct's json field set.
// missingInGo and missingReqd are disjoint: a missing spec property goes to missingReqd
// if the spec marks it required, otherwise to missingInGo.
func buildDiff(so SpecOp, c ClientOp, kind, goType string, r resolved, goFields []string) FieldDiff {
	specSet := make(map[string]bool, len(r.props))
	for k := range r.props {
		specSet[k] = true
	}
	reqSet := make(map[string]bool, len(r.required))
	for _, k := range r.required {
		reqSet[k] = true
	}
	goSet := make(map[string]bool, len(goFields))
	for _, f := range goFields {
		goSet[f] = true
	}

	var missingInGo, missingReqd, missingInSpec []string
	for k := range specSet {
		if goSet[k] {
			continue
		}
		if reqSet[k] {
			missingReqd = append(missingReqd, k)
		} else {
			missingInGo = append(missingInGo, k)
		}
	}
	for f := range goSet {
		if !specSet[f] {
			missingInSpec = append(missingInSpec, f)
		}
	}
	sort.Strings(missingInGo)
	sort.Strings(missingReqd)
	sort.Strings(missingInSpec)

	d := FieldDiff{
		Op:            so,
		ClientOp:      c,
		Kind:          kind,
		GoType:        goType,
		SpecRef:       r.name,
		MissingInGo:   missingInGo,
		MissingReqd:   missingReqd,
		MissingInSpec: missingInSpec,
	}
	if r.isList && r.name == "(inline)" {
		d.LowConf = true
		d.LowReason = "inline list items"
	}
	return d
}
