// Command specsync compares the servers.com Go client against an OpenAPI spec and
// reports what is missing, stale, or field-mismatched. It is a read-only report tool:
// it never edits the client or the spec.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

type options struct {
	specPath      string
	pkgDir        string
	jsonOut       bool
	tagFilter     string
	failOnMissing bool
	failOnStale   bool
	debug         bool
}

func main() {
	var opt options
	flag.StringVar(&opt.specPath, "spec", "./openapi_spec.json", "path to the OpenAPI spec JSON")
	flag.StringVar(&opt.pkgDir, "pkg", "./pkg", "path to the Go client package directory")
	flag.BoolVar(&opt.jsonOut, "json", false, "emit the full result as JSON")
	flag.StringVar(&opt.tagFilter, "tag", "", "restrict the report to a single spec tag")
	flag.BoolVar(&opt.failOnMissing, "fail-on-missing", false, "exit 1 if any spec operation is missing from the client")
	flag.BoolVar(&opt.failOnStale, "fail-on-stale", false, "exit 1 if any client operation is absent from the spec")
	flag.BoolVar(&opt.debug, "debug", false, "show unresolved methods and per-op coverage detail")
	flag.Parse()

	doc, specOps, err := loadSpec(opt.specPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "specsync:", err)
		os.Exit(2)
	}
	cm, err := parseClient(opt.pkgDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "specsync:", err)
		os.Exit(2)
	}

	res := buildResult(specOps, cm)
	res.FieldDiffs = computeFieldDiffs(doc, specOps, res, cm)

	if opt.jsonOut {
		if err := renderJSON(os.Stdout, res); err != nil {
			fmt.Fprintln(os.Stderr, "specsync:", err)
			os.Exit(2)
		}
	} else {
		renderText(os.Stdout, res, cm, opt)
	}

	exit := 0
	if opt.failOnMissing && len(filterMissingByTag(res.Missing, opt.tagFilter)) > 0 {
		exit = 1
	}
	if opt.failOnStale && len(res.Stale) > 0 {
		exit = 1
	}
	os.Exit(exit)
}

// buildResult matches the client against the spec and rolls up coverage.
func buildResult(specOps []SpecOp, cm *clientModel) *MatchResult {
	missing, stale, coveredBy := computeCoverage(specOps, cm.ops)
	res := &MatchResult{
		Missing:       missing,
		Stale:         stale,
		CoveredBy:     coveredBy,
		SpecOpCount:   len(specOps),
		ClientOpCount: len(cm.ops),
		CoveredCount:  len(specOps) - len(missing),
	}

	tagTotals := map[string]int{}
	tagCovered := map[string]int{}
	for _, so := range specOps {
		t := so.PrimaryTag()
		tagTotals[t]++
		if _, ok := coveredBy[so.matchKey()]; ok {
			tagCovered[t]++
		}
	}
	tags := make([]string, 0, len(tagTotals))
	for t := range tagTotals {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	for _, t := range tags {
		res.PerTag = append(res.PerTag, Coverage{Tag: t, Total: tagTotals[t], Covered: tagCovered[t]})
	}
	return res
}
