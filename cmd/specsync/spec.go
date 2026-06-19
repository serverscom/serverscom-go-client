package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// openAPI is the minimal subset of an OpenAPI 3.0 document we need.
type openAPI struct {
	Paths      map[string]*pathItem `json:"paths"`
	Components struct {
		Schemas map[string]*Schema `json:"schemas"`
	} `json:"components"`
}

// pathItem only declares the HTTP-method operations; non-operation keys
// (parameters, summary, description, servers, $ref) are ignored on purpose so
// they don't fail to decode into *operation.
type pathItem struct {
	Get    *operation `json:"get"`
	Post   *operation `json:"post"`
	Put    *operation `json:"put"`
	Delete *operation `json:"delete"`
	Patch  *operation `json:"patch"`
}

func (p *pathItem) byVerb() map[string]*operation {
	m := map[string]*operation{}
	if p.Get != nil {
		m["GET"] = p.Get
	}
	if p.Post != nil {
		m["POST"] = p.Post
	}
	if p.Put != nil {
		m["PUT"] = p.Put
	}
	if p.Delete != nil {
		m["DELETE"] = p.Delete
	}
	if p.Patch != nil {
		m["PATCH"] = p.Patch
	}
	return m
}

type operation struct {
	OperationID string               `json:"operationId"`
	Summary     string               `json:"summary"`
	Tags        []string             `json:"tags"`
	RequestBody *requestBody         `json:"requestBody"`
	Responses   map[string]*response `json:"responses"`
}

type requestBody struct {
	Content map[string]*mediaType `json:"content"`
}

type response struct {
	Content map[string]*mediaType `json:"content"`
}

type mediaType struct {
	Schema *Schema `json:"schema"`
}

// Schema is a permissive view of an OpenAPI schema node. The servers.com spec
// sometimes attaches several shapes to one node (inline properties AND a $ref AND
// items); resolution below imposes a clear precedence.
type Schema struct {
	Ref        string             `json:"$ref"`
	Type       string             `json:"type"`
	Format     string             `json:"format"`
	Enum       []interface{}      `json:"enum"`
	Properties map[string]*Schema `json:"properties"`
	Items      *Schema            `json:"items"`
	Required   []string           `json:"required"`
	AllOf      []*Schema          `json:"allOf"`
	Nullable   bool               `json:"nullable"`
	Deprecated bool               `json:"deprecated"`
}

const jsonMediaType = "application/json"

// loadSpec parses an OpenAPI document and flattens it into SpecOps. The source may be a
// local file path or an http(s) URL (fetched with a short timeout).
func loadSpec(src string) (*openAPI, []SpecOp, error) {
	raw, err := readSpec(src)
	if err != nil {
		return nil, nil, err
	}

	var doc openAPI
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, nil, fmt.Errorf("parse spec: %w", err)
	}

	var ops []SpecOp
	for rawPath, item := range doc.Paths {
		if item == nil {
			continue
		}
		for verb, op := range item.byVerb() {
			so := SpecOp{
				Verb:        verb,
				RawPath:     rawPath,
				Norm:        normalizeSpecPath(rawPath),
				OperationID: op.OperationID,
				Summary:     op.Summary,
				Tags:        op.Tags,
			}
			if op.RequestBody != nil {
				if mt := op.RequestBody.Content[jsonMediaType]; mt != nil {
					so.HasReqBody = true
					so.ReqBodySchema = mt.Schema
				}
			}
			so.SuccessCode, so.SuccessSchema = pickSuccessSchema(op.Responses)
			ops = append(ops, so)
		}
	}

	sort.Slice(ops, func(i, j int) bool {
		if ops[i].RawPath != ops[j].RawPath {
			return ops[i].RawPath < ops[j].RawPath
		}
		return ops[i].Verb < ops[j].Verb
	})
	return &doc, ops, nil
}

// readSpec reads the spec bytes from a local path or an http(s) URL.
func readSpec(src string) ([]byte, error) {
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Get(src)
		if err != nil {
			return nil, fmt.Errorf("fetch spec %q: %w", src, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("fetch spec %q: HTTP %d", src, resp.StatusCode)
		}
		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read spec body %q: %w", src, err)
		}
		return raw, nil
	}
	raw, err := os.ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("read spec: %w", err)
	}
	return raw, nil
}

// pickSuccessSchema returns the lowest 2xx response with a JSON schema.
func pickSuccessSchema(responses map[string]*response) (string, *Schema) {
	codes := make([]string, 0, len(responses))
	for code := range responses {
		if strings.HasPrefix(code, "2") {
			codes = append(codes, code)
		}
	}
	sort.Strings(codes)
	for _, code := range codes {
		resp := responses[code]
		if resp == nil {
			continue
		}
		if mt := resp.Content[jsonMediaType]; mt != nil && mt.Schema != nil {
			return code, mt.Schema
		}
	}
	return "", nil
}

const schemaRefPrefix = "#/components/schemas/"

// resolved is the outcome of reducing a schema node to a comparable property set.
type resolved struct {
	props     map[string]*Schema
	required  []string
	name      string // component name, or "(inline)"
	isList    bool   // came through an array/items wrapper
	lowConf   bool
	lowReason string
}

// resolveProps reduces a schema node to its effective object property set.
//
// Precedence (matches the spec's quirks):
//  1. $ref wins over any sibling inline props (the spec attaches stray message/code
//     props next to a $ref on some single-entity responses).
//  2. array / items -> resolve the item schema, flag as a list.
//  3. inline properties.
//  4. allOf -> union of members' properties.
//
// ok is false when there is nothing meaningful to compare (nil node, or a free-form
// object with no declared properties, e.g. a labels map).
func (o *openAPI) resolveProps(s *Schema) (resolved, bool) {
	return o.resolve(s, map[string]bool{})
}

func (o *openAPI) resolve(s *Schema, seen map[string]bool) (resolved, bool) {
	if s == nil {
		return resolved{}, false
	}

	if s.Ref != "" {
		name := strings.TrimPrefix(s.Ref, schemaRefPrefix)
		if seen[name] {
			return resolved{}, false
		}
		seen[name] = true
		target := o.Components.Schemas[name]
		r, ok := o.resolve(target, seen)
		if ok && r.name == "(inline)" {
			r.name = name
		}
		return r, ok
	}

	if s.Type == "array" || s.Items != nil {
		r, ok := o.resolve(s.Items, seen)
		r.isList = true
		return r, ok
	}

	if len(s.Properties) > 0 {
		return resolved{props: s.Properties, required: s.Required, name: "(inline)"}, true
	}

	if len(s.AllOf) > 0 {
		merged := map[string]*Schema{}
		var required []string
		any := false
		for _, member := range s.AllOf {
			r, ok := o.resolve(member, seen)
			if !ok {
				continue
			}
			any = true
			for k, v := range r.props {
				merged[k] = v
			}
			required = append(required, r.required...)
		}
		if any {
			return resolved{props: merged, required: required, name: "(inline)"}, true
		}
	}

	return resolved{}, false
}

// resolveFields reduces a schema node to ordered, typed fields. refToGo maps a spec
// component schema name to its known Go struct name (for $ref fields). Fields are sorted
// by name for deterministic output.
func (o *openAPI) resolveFields(s *Schema, refToGo map[string]string) (name string, fields []FieldInfo, isList, ok bool) {
	r, rok := o.resolveProps(s)
	if !rok || len(r.props) == 0 {
		return "", nil, r.isList, false
	}
	reqSet := make(map[string]bool, len(r.required))
	for _, k := range r.required {
		reqSet[k] = true
	}
	for prop, ps := range r.props {
		if ps != nil && ps.Deprecated {
			continue
		}
		gt, note := goType(ps, refToGo)
		fields = append(fields, FieldInfo{Name: prop, GoType: gt, Required: reqSet[prop], Note: note})
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	return r.name, fields, r.isList, true
}

// goType maps a single schema node to a Go type and a short note (format / enum values /
// nesting). It is intentionally single-level: nested objects collapse to "object" and
// arrays of objects to "[]object", leaving the Go shape for a human/agent to refine.
// A nullable scalar or struct field is reported as a pointer (e.g. *string).
func goType(s *Schema, refToGo map[string]string) (gotype, note string) {
	gt, note := baseGoType(s, refToGo)
	if s != nil && s.Nullable && pointerable(gt) {
		gt = "*" + gt
	}
	return gt, note
}

// pointerable reports whether prefixing "*" is meaningful. Slices, maps and
// interface{} are already nil-able in Go, and "object" is a placeholder we
// don't refine, so they are left as-is.
func pointerable(gt string) bool {
	switch {
	case strings.HasPrefix(gt, "[]"),
		strings.HasPrefix(gt, "map["),
		gt == "interface{}",
		gt == "object":
		return false
	}
	return true
}

// baseGoType maps a schema node to its Go type without considering nullability.
func baseGoType(s *Schema, refToGo map[string]string) (gotype, note string) {
	if s == nil {
		return "interface{}", ""
	}
	if s.Ref != "" {
		ref := strings.TrimPrefix(s.Ref, schemaRefPrefix)
		if g, found := refToGo[ref]; found {
			return g, ""
		}
		return "object", "schema " + ref
	}
	switch s.Type {
	case "string":
		if len(s.Enum) > 0 {
			return "string", "enum: " + enumValues(s.Enum)
		}
		if s.Format != "" {
			return "string", s.Format
		}
		return "string", ""
	case "integer":
		return "int", ""
	case "number":
		return "float64", ""
	case "boolean":
		return "bool", ""
	case "array":
		et, n := goType(s.Items, refToGo)
		return "[]" + et, n
	case "object":
		if len(s.Properties) > 0 {
			return "object", "nested object"
		}
		return "map[string]string", "free-form object"
	}
	if len(s.Properties) > 0 {
		return "object", "nested object"
	}
	return "interface{}", ""
}

func enumValues(enum []interface{}) string {
	parts := make([]string, 0, len(enum))
	for _, e := range enum {
		parts = append(parts, fmt.Sprintf("%v", e))
	}
	return strings.Join(parts, ", ")
}
