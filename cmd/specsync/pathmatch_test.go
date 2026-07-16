package main

import "testing"

func TestNormalizeSpecPathStripsVersion(t *testing.T) {
	got := normalizeSpecPath("/v1/ssh_keys/{fingerprint}")
	if got.Raw != "/ssh_keys/{}" {
		t.Fatalf("Raw = %q, want %q", got.Raw, "/ssh_keys/{}")
	}
}

func TestPathsMatch(t *testing.T) {
	cases := []struct {
		name        string
		clientTmpl  string
		clientSubst []string
		specPath    string
		want        bool
	}{
		{
			name:       "single id matches param",
			clientTmpl: "/ssh_keys/%s",
			specPath:   "/v1/ssh_keys/{fingerprint}",
			want:       true,
		},
		{
			name:       "concrete sub-resource does NOT cover a parameterized path",
			clientTmpl: "/l2_segments/location_groups",
			specPath:   "/v1/l2_segments/{l2_segment_id}",
			want:       false,
		},
		{
			name:       "parameterized single segment matches param",
			clientTmpl: "/l2_segments/%s",
			specPath:   "/v1/l2_segments/{l2_segment_id}",
			want:       true,
		},
		{
			name:       "load balancer l4 vs l4",
			clientTmpl: "/load_balancers/l4/%s",
			specPath:   "/v1/load_balancers/l4/{id}",
			want:       true,
		},
		{
			name:       "load balancer l4 must not cover l7",
			clientTmpl: "/load_balancers/l4/%s",
			specPath:   "/v1/load_balancers/l7/{id}",
			want:       false,
		},
		{
			name:        "generic host path substituted to its concrete type",
			clientTmpl:  "/hosts/%s/%s/networks",
			clientSubst: []string{"dedicated_servers", ""},
			specPath:    "/v1/hosts/dedicated_servers/{server_id}/networks",
			want:        true,
		},
		{
			name:        "substituted generic host path must not cover a different host type",
			clientTmpl:  "/hosts/%s/%s/networks",
			clientSubst: []string{"dedicated_servers", ""},
			specPath:    "/v1/hosts/sbm_servers/{server_id}/networks",
			want:        false,
		},
		{
			name:       "client wildcard covers a concrete spec literal (feature name)",
			clientTmpl: "/hosts/dedicated_servers/%s/features/%s/activate",
			specPath:   "/v1/hosts/dedicated_servers/{server_id}/features/no_private_ip/activate",
			want:       true,
		},
		{
			name:       "deeply nested %d order options",
			clientTmpl: "/locations/%d/order_options/server_models/%d",
			specPath:   "/v1/locations/{location_id}/order_options/server_models/{server_model_id}",
			want:       true,
		},
		{
			name:       "different segment counts never match",
			clientTmpl: "/hosts/dedicated_servers/%s",
			specPath:   "/v1/hosts/dedicated_servers/{server_id}/networks",
			want:       false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := normalizeClientPath(tc.clientTmpl, tc.clientSubst)
			spec := normalizeSpecPath(tc.specPath)
			if got := pathsMatch(client, spec); got != tc.want {
				t.Errorf("pathsMatch(%q, %q) = %v, want %v", client.Raw, spec.Raw, got, tc.want)
			}
		})
	}
}
