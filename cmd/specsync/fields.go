package main

import "sort"

// computeFieldDiffs produces advisory field-level comparisons for matched operations.
// Each diff is traceable to a single operation and its resolved struct/schema pair.
// Diffs are deduplicated by (kind, Go struct, spec schema) since many operations share
// the same entity type.
func computeFieldDiffs(doc *openAPI, specOps []SpecOp, res *MatchResult, cm *clientModel, refToGo map[string]string) []FieldDiff {
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
				add(responseDiff(doc, so, c, cm, refToGo))
			}
		}
		if so.ReqBodySchema != nil {
			if c, ok := pickClientWithInput(covers, cm); ok {
				add(requestDiff(doc, so, c, cm, refToGo))
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

func responseDiff(doc *openAPI, so SpecOp, c ClientOp, cm *clientModel, refToGo map[string]string) (FieldDiff, bool) {
	name, fields, isList, ok := doc.resolveFields(so.SuccessSchema, refToGo)
	if !ok {
		return FieldDiff{}, false
	}
	goFields := cm.jsonFields(c.ReturnType)
	if len(goFields) == 0 {
		return FieldDiff{}, false
	}
	return buildDiff(so, c, "response", c.ReturnType, name, fields, isList, goFields), true
}

func requestDiff(doc *openAPI, so SpecOp, c ClientOp, cm *clientModel, refToGo map[string]string) (FieldDiff, bool) {
	name, fields, isList, ok := doc.resolveFields(so.ReqBodySchema, refToGo)
	if !ok {
		return FieldDiff{}, false
	}
	goFields := cm.jsonFields(c.InputType)
	if len(goFields) == 0 {
		return FieldDiff{}, false
	}
	return buildDiff(so, c, "request", c.InputType, name, fields, isList, goFields), true
}

// buildDiff compares a resolved schema's typed fields against a struct's json field set.
// missingInGo and missingReqd are disjoint: a missing spec property goes to missingReqd
// if the spec marks it required, otherwise to missingInGo.
func buildDiff(so SpecOp, c ClientOp, kind, goType, specName string, fields []FieldInfo, isList bool, goFields []string) FieldDiff {
	specSet := make(map[string]bool, len(fields))
	for _, f := range fields {
		specSet[f.Name] = true
	}
	goSet := make(map[string]bool, len(goFields))
	for _, f := range goFields {
		goSet[f] = true
	}

	var missingInGo, missingReqd []FieldInfo
	for _, f := range fields {
		if goSet[f.Name] {
			continue
		}
		if f.Required {
			missingReqd = append(missingReqd, f)
		} else {
			missingInGo = append(missingInGo, f)
		}
	}
	var missingInSpec []string
	for f := range goSet {
		if !specSet[f] {
			missingInSpec = append(missingInSpec, f)
		}
	}
	sort.Strings(missingInSpec)

	d := FieldDiff{
		Op:            so,
		ClientOp:      c,
		Kind:          kind,
		GoType:        goType,
		SpecRef:       specName,
		MissingInGo:   missingInGo,
		MissingReqd:   missingReqd,
		MissingInSpec: missingInSpec,
	}
	if isList && specName == "(inline)" {
		d.LowConf = true
		d.LowReason = "inline list items"
	}
	return d
}

// buildRefIndex maps a spec component-schema name to the Go struct a covered operation
// returns for it, so $ref fields in instructions/diffs can name the real Go type.
func buildRefIndex(doc *openAPI, specOps []SpecOp, res *MatchResult, cm *clientModel) map[string]string {
	idx := map[string]string{}
	for _, so := range specOps {
		if so.SuccessSchema == nil {
			continue
		}
		covers := res.CoveredBy[so.matchKey()]
		if len(covers) == 0 {
			continue
		}
		r, ok := doc.resolveProps(so.SuccessSchema)
		if !ok || r.name == "" || r.name == "(inline)" {
			continue
		}
		if _, exists := idx[r.name]; exists {
			continue
		}
		if c, ok := pickClientWithReturn(covers, cm); ok {
			idx[r.name] = c.ReturnType
		}
	}
	return idx
}
