package main

import "strings"

// splitPath splits a URL path into non-empty segments.
func splitPath(p string) []string {
	var out []string
	for _, s := range strings.Split(p, "/") {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// renderNorm renders normalized segments back to a readable path, with wildcards
// shown uniformly as "{}" so client and spec paths render comparably.
func renderNorm(segs []Segment) string {
	if len(segs) == 0 {
		return "/"
	}
	var b strings.Builder
	for _, s := range segs {
		b.WriteByte('/')
		if s.Kind == segWildcard {
			b.WriteString("{}")
		} else {
			b.WriteString(s.Text)
		}
	}
	return b.String()
}

// normalizeSpecPath canonicalizes a spec path: the leading "v1" version segment is
// dropped (the client base URL already includes it) and "{param}" segments become
// wildcards.
func normalizeSpecPath(rawPath string) NormPath {
	parts := splitPath(rawPath)
	if len(parts) > 0 && parts[0] == "v1" {
		parts = parts[1:]
	}
	segs := make([]Segment, 0, len(parts))
	for _, p := range parts {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			segs = append(segs, Segment{Kind: segWildcard, Text: p})
		} else {
			segs = append(segs, Segment{Kind: segLiteral, Text: p})
		}
	}
	return NormPath{Raw: renderNorm(segs), Segs: segs}
}

// normalizeClientPath canonicalizes a client path template ("/hosts/%s/%s/networks").
// "%s"/"%d" segments are wildcards, except where a positional literal substitution is
// known (e.g. the host-type prefix "dedicated_servers"), which collapses the wildcard
// to that concrete literal so a generic const matches exactly one spec path.
func normalizeClientPath(template string, substs []string) NormPath {
	parts := splitPath(template)
	segs := make([]Segment, 0, len(parts))
	wc := 0
	for _, p := range parts {
		if strings.Contains(p, "%") {
			if wc < len(substs) && substs[wc] != "" {
				segs = append(segs, Segment{Kind: segLiteral, Text: substs[wc]})
			} else {
				segs = append(segs, Segment{Kind: segWildcard, Text: p})
			}
			wc++
		} else {
			segs = append(segs, Segment{Kind: segLiteral, Text: p})
		}
	}
	return NormPath{Raw: renderNorm(segs), Segs: segs}
}

// pathsMatch reports whether a client path covers a spec path, segment by segment.
// A client wildcard ("%s"/"%d", or an unsubstituted host-type prefix) matches any spec
// segment; a client literal matches ONLY an equal spec literal. A client literal does
// NOT match a spec path-parameter — otherwise a concrete sub-resource like
// "/l2_segments/location_groups" would spuriously cover "/l2_segments/{id}".
func pathsMatch(client, spec NormPath) bool {
	if len(client.Segs) != len(spec.Segs) {
		return false
	}
	for i := range client.Segs {
		c, s := client.Segs[i], spec.Segs[i]
		if c.Kind == segWildcard {
			continue // client wildcard matches any spec segment
		}
		if s.Kind != segLiteral || c.Text != s.Text {
			return false
		}
	}
	return true
}

// opMatch reports whether a client op covers a spec op (verb + path).
func opMatch(c ClientOp, s SpecOp) bool {
	return strings.EqualFold(c.Verb, s.Verb) && pathsMatch(c.Norm, s.Norm)
}

// computeCoverage matches client ops against spec ops in both directions.
func computeCoverage(specOps []SpecOp, clientOps []ClientOp) (missing []SpecOp, stale []ClientOp, coveredBy map[string][]ClientOp) {
	coveredBy = map[string][]ClientOp{}
	clientMatched := make([]bool, len(clientOps))

	for _, so := range specOps {
		var covers []ClientOp
		for i := range clientOps {
			if opMatch(clientOps[i], so) {
				covers = append(covers, clientOps[i])
				clientMatched[i] = true
			}
		}
		if len(covers) == 0 {
			missing = append(missing, so)
		} else {
			coveredBy[so.matchKey()] = covers
		}
	}

	for i, c := range clientOps {
		if !clientMatched[i] {
			stale = append(stale, c)
		}
	}
	return missing, stale, coveredBy
}
