package main

import "testing"

func clientOp(verb, tmpl, opID string) ClientOp {
	return ClientOp{Verb: verb, PathValue: tmpl, OperationID: opID, Norm: normalizeClientPath(tmpl, nil)}
}

func specOp(verb, path, opID string) SpecOp {
	return SpecOp{Verb: verb, RawPath: path, OperationID: opID, Norm: normalizeSpecPath(path)}
}

func TestPairRenamesOperationID(t *testing.T) {
	// RBS reset: same operationId, but segment counts differ
	// (.../credentials/reset is 2 segments, .../reset_credentials is 1) — only
	// operationId pairing can catch this.
	missing := []SpecOp{
		specOp("POST", "/v1/remote_block_storage/volumes/{id}/reset_credentials", "ResetCredentialsForAnRbsVolume"),
	}
	stale := []ClientOp{
		clientOp("POST", "/remote_block_storage/volumes/%s/credentials/reset", "ResetCredentialsForAnRbsVolume"),
	}
	renames, remMissing, remStale := pairRenames(missing, stale)
	if len(renames) != 1 {
		t.Fatalf("renames = %d, want 1", len(renames))
	}
	if renames[0].Reason != "operationId" {
		t.Errorf("reason = %q, want operationId", renames[0].Reason)
	}
	if len(remMissing) != 0 || len(remStale) != 0 {
		t.Errorf("leftovers: missing=%d stale=%d, want 0/0", len(remMissing), len(remStale))
	}
}

func TestPairRenamesPathSimilarity(t *testing.T) {
	// No operationId annotation: fall back to a single-literal-segment rename.
	missing := []SpecOp{
		specOp("POST", "/v1/cloud_computing/instances/{id}/switch_off", ""),
	}
	stale := []ClientOp{
		clientOp("POST", "/cloud_computing/instances/%s/switch_power_off", ""),
	}
	renames, remMissing, remStale := pairRenames(missing, stale)
	if len(renames) != 1 {
		t.Fatalf("renames = %d, want 1", len(renames))
	}
	if renames[0].Reason != "path-similarity" {
		t.Errorf("reason = %q, want path-similarity", renames[0].Reason)
	}
	if len(remMissing) != 0 || len(remStale) != 0 {
		t.Errorf("leftovers: missing=%d stale=%d, want 0/0", len(remMissing), len(remStale))
	}
}

func TestPairRenamesNoFalsePositive(t *testing.T) {
	// Unrelated ops (different verb, different path, no shared operationId) must not pair.
	missing := []SpecOp{specOp("GET", "/v1/metrics/hosts", "ListHostsMetrics")}
	stale := []ClientOp{clientOp("DELETE", "/ssh_keys/%s", "")}
	renames, remMissing, remStale := pairRenames(missing, stale)
	if len(renames) != 0 {
		t.Fatalf("renames = %d, want 0", len(renames))
	}
	if len(remMissing) != 1 || len(remStale) != 1 {
		t.Errorf("leftovers: missing=%d stale=%d, want 1/1", len(remMissing), len(remStale))
	}
}
