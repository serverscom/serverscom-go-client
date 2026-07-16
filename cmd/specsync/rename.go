package main

import "strings"

// pairRenames matches leftover stale client ops against missing spec ops that are the
// same operation under a changed path. It returns the pairs plus the still-unpaired
// missing and stale ops. Pairing is one-to-one and greedy: operationId matches first
// (segment-count independent), then a path-similarity fallback for unannotated methods.
func pairRenames(missing []SpecOp, stale []ClientOp) (renames []Rename, remMissing []SpecOp, remStale []ClientOp) {
	missClaimed := make([]bool, len(missing))
	staleClaimed := make([]bool, len(stale))

	// Tier 1: operationId equality.
	specByOpID := map[string]int{}
	for i, so := range missing {
		if so.OperationID != "" {
			if _, exists := specByOpID[so.OperationID]; !exists {
				specByOpID[so.OperationID] = i
			}
		}
	}
	for si, c := range stale {
		if c.OperationID == "" {
			continue
		}
		if mi, ok := specByOpID[c.OperationID]; ok && !missClaimed[mi] {
			renames = append(renames, Rename{Client: c, Spec: missing[mi], Reason: "operationId"})
			missClaimed[mi] = true
			staleClaimed[si] = true
		}
	}

	// Tier 2: path-similarity fallback (same verb, exactly one renamed literal segment).
	for si, c := range stale {
		if staleClaimed[si] {
			continue
		}
		for mi, so := range missing {
			if missClaimed[mi] {
				continue
			}
			if !strings.EqualFold(c.Verb, so.Verb) {
				continue
			}
			if oneLiteralRename(c.Norm, so.Norm) {
				renames = append(renames, Rename{Client: c, Spec: so, Reason: "path-similarity"})
				missClaimed[mi] = true
				staleClaimed[si] = true
				break
			}
		}
	}

	for i, so := range missing {
		if !missClaimed[i] {
			remMissing = append(remMissing, so)
		}
	}
	for i, c := range stale {
		if !staleClaimed[i] {
			remStale = append(remStale, c)
		}
	}
	return renames, remMissing, remStale
}

// oneLiteralRename reports whether two normalized paths are identical except for exactly
// one segment, where that segment is a literal on both sides (a plain rename).
func oneLiteralRename(a, b NormPath) bool {
	if len(a.Segs) != len(b.Segs) {
		return false
	}
	diffs := 0
	diffLiteral := false
	for i := range a.Segs {
		if segEqual(a.Segs[i], b.Segs[i]) {
			continue
		}
		diffs++
		diffLiteral = a.Segs[i].Kind == segLiteral && b.Segs[i].Kind == segLiteral
	}
	return diffs == 1 && diffLiteral
}

func segEqual(a, b Segment) bool {
	if a.Kind == segWildcard && b.Kind == segWildcard {
		return true
	}
	return a.Kind == segLiteral && b.Kind == segLiteral && a.Text == b.Text
}
