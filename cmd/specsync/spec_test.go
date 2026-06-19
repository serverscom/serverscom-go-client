package main

import "testing"

func TestGoType(t *testing.T) {
	refToGo := map[string]string{"v1-cloud-computing-instances-entity": "CloudComputingInstance"}
	cases := []struct {
		name     string
		schema   *Schema
		wantType string
		wantNote string
	}{
		{"string", &Schema{Type: "string"}, "string", ""},
		{"string with format", &Schema{Type: "string", Format: "date-time"}, "string", "date-time"},
		{"string enum", &Schema{Type: "string", Enum: []interface{}{"a", "b"}}, "string", "enum: a, b"},
		{"integer", &Schema{Type: "integer"}, "int", ""},
		{"number", &Schema{Type: "number"}, "float64", ""},
		{"boolean", &Schema{Type: "boolean"}, "bool", ""},
		{"array of string", &Schema{Type: "array", Items: &Schema{Type: "string"}}, "[]string", ""},
		{"free-form object", &Schema{Type: "object"}, "map[string]string", "free-form object"},
		{"nested object", &Schema{Type: "object", Properties: map[string]*Schema{"x": {Type: "string"}}}, "object", "nested object"},
		{"known ref", &Schema{Ref: schemaRefPrefix + "v1-cloud-computing-instances-entity"}, "CloudComputingInstance", ""},
		{"unknown ref", &Schema{Ref: schemaRefPrefix + "v1-mystery"}, "object", "schema v1-mystery"},
		{"nil", nil, "interface{}", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gt, note := goType(tc.schema, refToGo)
			if gt != tc.wantType || note != tc.wantNote {
				t.Errorf("goType = (%q, %q), want (%q, %q)", gt, note, tc.wantType, tc.wantNote)
			}
		})
	}
}
