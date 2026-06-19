package main

import (
	"os"
	"path/filepath"
	"testing"
)

// A thin wrapper that only delegates to an internal path-building helper, passing a
// literal for the feature segment, must resolve to a concrete path (not inherit the
// helper's catch-all wildcard) — and the helper's own wildcard op must be dropped.
const delegationSrc = `package fake

type FooHandler struct{}

const featActivatePath = "/things/%s/features/%s/activate"

func (h *FooHandler) activateFeature(ctx interface{}, id, feature string) interface{} {
	url := h.client.buildURL(featActivatePath, id, feature)
	return h.client.buildAndExecRequest(ctx, "POST", url, nil)
}

// ActivateAlphaFeature activates alpha.
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Foo/operation/ActivateAlphaForAThing
func (h *FooHandler) ActivateAlphaFeature(ctx interface{}, id string) interface{} {
	return h.activateFeature(ctx, id, "alpha")
}

// ActivateBetaFeature activates beta.
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Foo/operation/ActivateBetaForAThing
func (h *FooHandler) ActivateBetaFeature(ctx interface{}, id string) interface{} {
	return h.activateFeature(ctx, id, "beta")
}
`

func TestResolveDelegations(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "fake.go"), []byte(delegationSrc), 0o600); err != nil {
		t.Fatal(err)
	}

	m, err := parseClient(dir)
	if err != nil {
		t.Fatal(err)
	}

	byMethod := map[string]ClientOp{}
	for _, op := range m.ops {
		byMethod[op.Method] = op
	}

	// The unexported helper's wildcard op is suppressed once wrappers resolve it.
	if _, ok := byMethod["activateFeature"]; ok {
		t.Error("helper activateFeature should be dropped after delegation resolution")
	}

	// Each wrapper resolves to its own literal feature segment + carries its operationId.
	want := map[string]struct{ raw, opID string }{
		"ActivateAlphaFeature": {"/things/{}/features/alpha/activate", "ActivateAlphaForAThing"},
		"ActivateBetaFeature":  {"/things/{}/features/beta/activate", "ActivateBetaForAThing"},
	}
	for method, w := range want {
		op, ok := byMethod[method]
		if !ok {
			t.Errorf("missing resolved op for %s", method)
			continue
		}
		if op.Norm.Raw != w.raw {
			t.Errorf("%s path = %q, want %q", method, op.Norm.Raw, w.raw)
		}
		if op.OperationID != w.opID {
			t.Errorf("%s operationId = %q, want %q", method, op.OperationID, w.opID)
		}
		if op.Verb != "POST" {
			t.Errorf("%s verb = %q, want POST", method, op.Verb)
		}
	}

	if len(m.ops) != 2 {
		t.Errorf("ops = %d, want 2 (two resolved wrappers, helper dropped)", len(m.ops))
	}
}
