package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
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
	Properties map[string]*Schema `json:"properties"`
	Items      *Schema            `json:"items"`
	Required   []string           `json:"required"`
	AllOf      []*Schema          `json:"allOf"`
}

const jsonMediaType = "application/json"

// loadSpec parses an OpenAPI document and flattens it into SpecOps.
func loadSpec(path string) (*openAPI, []SpecOp, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("read spec: %w", err)
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
