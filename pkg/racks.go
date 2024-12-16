package serverscom

import (
	"context"
	"encoding/json"
)

const (
	rackPath       = "/racks"
	rackPathWithID = "/racks/%s"
)

// RacksService is an interface for interfacing with Rack endpoints
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Rack
type RacksService interface {
	// Primary collection
	Collection() Collection[Rack]

	// Generic operations
	Get(ctx context.Context, id string) (*Rack, error)
	Update(ctx context.Context, id string, input RackUpdateInput) (*Rack, error)
}

// RacksHandler handles operations around hosts
type RacksHandler struct {
	client *Client
}

// Collection builds a new Collection[Rack] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Rack/operation/ListRacks
func (h *RacksHandler) Collection() Collection[Rack] {
	return NewCollection[Rack](h.client, rackPath)
}

// GetRack returns a rack
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Rack/operation/GetARack
func (h *RacksHandler) Get(ctx context.Context, id string) (*Rack, error) {
	url := h.client.buildURL(rackPathWithID, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	rack := new(Rack)

	if err := json.Unmarshal(body, &rack); err != nil {
		return nil, err
	}

	return rack, nil
}

// UpdateRack updates rack
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Rack/operation/UpdateARack
func (h *RacksHandler) Update(ctx context.Context, id string, input RackUpdateInput) (*Rack, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(rackPathWithID, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	rack := new(Rack)

	if err := json.Unmarshal(body, &rack); err != nil {
		return nil, err
	}

	return rack, nil
}
