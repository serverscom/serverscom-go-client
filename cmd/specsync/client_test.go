package main

import (
	"go/ast"
	"testing"
)

func commentGroup(lines ...string) *ast.CommentGroup {
	cg := &ast.CommentGroup{}
	for _, l := range lines {
		cg.List = append(cg.List, &ast.Comment{Text: l})
	}
	return cg
}

func TestOperationID(t *testing.T) {
	cases := []struct {
		name string
		doc  *ast.CommentGroup
		want string
	}{
		{
			name: "endpoint annotation",
			doc:  commentGroup("// PowerOff cloud instance", "// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/PowerOffACloudInstance"),
			want: "PowerOffACloudInstance",
		},
		{
			name: "tag-only annotation, no operation",
			doc:  commentGroup("// https://developers.servers.com/api-documentation/v1/#tag/Account"),
			want: "",
		},
		{
			name: "plain doc comment",
			doc:  commentGroup("// Collection builds a collection of hosts"),
			want: "",
		},
		{
			name: "nil doc",
			doc:  nil,
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := operationID(tc.doc); got != tc.want {
				t.Errorf("operationID = %q, want %q", got, tc.want)
			}
		})
	}
}
