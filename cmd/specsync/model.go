package main

// Segment kinds for a normalized path.
const (
	segLiteral  = "literal"
	segWildcard = "wildcard"
)

// Segment is one '/'-delimited piece of a normalized path.
//
// A literal segment carries its concrete text (e.g. "hosts", "dedicated_servers",
// "power_on"). A wildcard segment is a path parameter — "%s"/"%d" on the client side or
// "{server_id}" on the spec side — and its Text is kept only for display.
type Segment struct {
	Kind string
	Text string
}

// NormPath is a path reduced to a canonical, comparable form. Raw keeps the
// human-readable rendering used in reports.
type NormPath struct {
	Raw  string
	Segs []Segment
}

// ClientOp is a single API operation discovered in the Go client by static analysis.
type ClientOp struct {
	Verb      string // GET / POST / PUT / DELETE
	PathValue string // raw template value, e.g. "/hosts/%s/%s/networks"

	// Substs holds positional literal substitutions for the path's wildcards.
	// Substs[i] is the resolved literal for the i-th wildcard (e.g. "dedicated_servers"),
	// or "" when the wildcard is filled by a runtime value (an id).
	Substs []string

	Norm NormPath

	RecvType string // receiver type, e.g. "HostsHandler"
	Method   string // method name, e.g. "GetDedicatedServer"
	File     string
	Line     int

	OperationID string // spec operationId from the method's "operation/<ID>" doc comment, or ""

	ReturnType string // base entity type name (pointer/slice/Collection unwrapped), or ""
	InputType  string // request input struct name, or ""
	IsList     bool   // discovered via NewCollection[T]
}

// SpecOp is a single API operation declared in the OpenAPI spec.
type SpecOp struct {
	Verb        string
	RawPath     string // with leading /v1
	Norm        NormPath
	OperationID string
	Summary     string
	Tags        []string

	HasReqBody    bool
	ReqBodySchema *Schema // request body application/json schema node (may be nil)
	SuccessSchema *Schema // chosen 2xx application/json schema node (may be nil)
	SuccessCode   string
}

// PrimaryTag returns the first tag, or "(untagged)".
func (s SpecOp) PrimaryTag() string {
	if len(s.Tags) > 0 {
		return s.Tags[0]
	}
	return "(untagged)"
}

// matchKey is a stable identity for a spec operation, used as a map key.
func (s SpecOp) matchKey() string { return s.Verb + " " + s.RawPath }

// FieldInfo is a spec property reduced to a Go-facing shape: its json name, the mapped
// Go type, whether the spec marks it required, and a short note (format/enum/nesting).
type FieldInfo struct {
	Name     string
	GoType   string
	Required bool
	Note     string
}

// Rename pairs a stale client op with a missing spec op that is the same operation under
// a changed path (the client const points at the old path).
type Rename struct {
	Client ClientOp
	Spec   SpecOp
	Reason string // "operationId" or "path-similarity"
}

// FieldDiff is the advisory field-level comparison for one matched operation.
type FieldDiff struct {
	Op       SpecOp
	ClientOp ClientOp

	Kind      string // "response" or "request"
	GoType    string // Go struct compared
	SpecRef   string // schema name or "(inline)"
	LowConf   bool   // resolution was heuristic (e.g. inline list items)
	LowReason string

	MissingInGo   []FieldInfo // properties in spec, absent from the Go struct
	MissingInSpec []string    // json fields in the Go struct, absent from the spec
	MissingReqd   []FieldInfo // required spec properties absent from the Go struct
}

func (d FieldDiff) empty() bool {
	return len(d.MissingInGo) == 0 && len(d.MissingInSpec) == 0 && len(d.MissingReqd) == 0
}

// Coverage is a per-tag rollup for the summary section.
type Coverage struct {
	Tag     string `json:"tag"`
	Total   int    `json:"total"`
	Covered int    `json:"covered"`
}

// MatchResult is the full comparison outcome, also the shape emitted by -json.
type MatchResult struct {
	Missing    []SpecOp    `json:"-"` // spec ops with no client match (renames removed)
	Stale      []ClientOp  `json:"-"` // client ops with no spec match (renames removed)
	Renames    []Rename    `json:"-"` // stale↔missing pairs that are the same op, path changed
	FieldDiffs []FieldDiff `json:"-"`

	// CoveredBy maps a spec op key to the client ops that cover it (for -debug/audit).
	CoveredBy map[string][]ClientOp `json:"-"`

	SpecOpCount   int
	ClientOpCount int
	CoveredCount  int

	PerTag []Coverage
}
