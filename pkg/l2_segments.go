package serverscom

import (
	"context"
	"encoding/json"
)

const (
	l2SegmentListPath           = "/l2_segments"
	l2SegmentPath               = "/l2_segments/%s"
	l2SegmentCreatePath         = "/l2_segments"
	l2SegmentUpdatePath         = "/l2_segments/%s"
	l2SegmentDeletePath         = "/l2_segments/%s"
	l2SegmentChangeNetworksPath = "/l2_segments/%s/networks"

	l2MemberListPath = "/l2_segments/%s/members"

	l2NetworksListPath = "/l2_segments/%s/networks"

	l2LocationGroupListPath = "/l2_segments/location_groups"
)

// L2SegmentsService is an interface for interfacing with Host, Dedicated Server endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/L2-Segment
type L2SegmentsService interface {
	// Primary collection
	Collection() Collection[L2Segment]

	// Extra collections
	LocationGroups() Collection[L2LocationGroup]

	// Generic operations
	Get(ctx context.Context, segmentID string) (*L2Segment, error)
	Create(ctx context.Context, input L2SegmentCreateInput) (*L2Segment, error)
	Update(ctx context.Context, segmentID string, input L2SegmentUpdateInput) (*L2Segment, error)
	Delete(ctx context.Context, segmentID string) error

	// Additional operations
	ChangeNetworks(ctx context.Context, segmentID string, input L2SegmentChangeNetworksInput) (*L2Segment, error)

	// Additional collections
	Members(segmentID string) Collection[L2Member]
	Networks(segmentID string) Collection[Network]
}

// L2SegmentsHandler handles  operatings around l2 segments
type L2SegmentsHandler struct {
	client *Client
}

// Collection builds a new Collection[L2Segment] interface
func (h *L2SegmentsHandler) Collection() Collection[L2Segment] {
	return NewCollection[L2Segment](h.client, l2SegmentListPath)
}

// Get l2 segment
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingL2Segment
func (h *L2SegmentsHandler) Get(ctx context.Context, segmentID string) (*L2Segment, error) {
	url := h.client.buildURL(l2SegmentPath, []interface{}{segmentID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	l2Segment := new(L2Segment)

	if err := json.Unmarshal(body, &l2Segment); err != nil {
		return nil, err
	}

	return l2Segment, nil
}

// Create l2 segment
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateANewL2Segment
func (h *L2SegmentsHandler) Create(ctx context.Context, input L2SegmentCreateInput) (*L2Segment, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l2SegmentCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	l2Segment := new(L2Segment)

	if err := json.Unmarshal(body, &l2Segment); err != nil {
		return nil, err
	}

	return l2Segment, nil
}

// Update l2 segment
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateAnExistingL2Segment
func (h *L2SegmentsHandler) Update(ctx context.Context, segmentID string, input L2SegmentUpdateInput) (*L2Segment, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l2SegmentUpdatePath, []interface{}{segmentID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	l2Segment := new(L2Segment)

	if err := json.Unmarshal(body, &l2Segment); err != nil {
		return nil, err
	}

	return l2Segment, nil
}

// Delete l2 segment
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteAnExistingL2Segment
func (h *L2SegmentsHandler) Delete(ctx context.Context, segmentID string) error {
	url := h.client.buildURL(l2SegmentDeletePath, []interface{}{segmentID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// LocationGroups builds a new Collection[L2LocationGroup] interface
func (h *L2SegmentsHandler) LocationGroups() Collection[L2LocationGroup] {
	return NewCollection[L2LocationGroup](h.client, l2LocationGroupListPath)
}

// Members builds a new Collection[L2Member] interface
func (h *L2SegmentsHandler) Members(segmentID string) Collection[L2Member] {
	path := h.client.buildPath(l2MemberListPath, []interface{}{segmentID}...)

	return NewCollection[L2Member](h.client, path)
}

// Networks builds a new L2NetworksCollection interface
func (h *L2SegmentsHandler) Networks(segmentID string) Collection[Network] {
	path := h.client.buildPath(l2NetworksListPath, []interface{}{segmentID}...)

	return NewCollection[Network](h.client, path)
}

// ChangeNetworks changes networks set
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateAnExistingL2SegmentNetworks
func (h *L2SegmentsHandler) ChangeNetworks(ctx context.Context, segmentID string, input L2SegmentChangeNetworksInput) (*L2Segment, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l2SegmentChangeNetworksPath, []interface{}{segmentID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	l2Segment := new(L2Segment)

	if err := json.Unmarshal(body, &l2Segment); err != nil {
		return nil, err
	}

	return l2Segment, nil
}
