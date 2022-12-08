package serverscom

import (
	"context"
	"encoding/json"
)

const (
	networkPoolListPath = "/network_pools"
	networkPoolPath     = "/network_pools/%s"

	subnetworkListPath   = "/network_pools/%s/subnetworks"
	subnetworkCreatePath = "/network_pools/%s/subnetworks"
	subnetworkPath       = "/network_pools/%s/subnetworks/%s"
)

// NetworkPoolsService is an interface to interfacing with the Network Pool endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Network-Pool
type NetworkPoolsService interface {
	// Primary collection
	Collection() Collection[NetworkPool]

	// Generic operations
	Get(ctx context.Context, id string) (*NetworkPool, error)
	Update(ctx context.Context, id string, input NetworkPoolInput) (*NetworkPool, error)

	CreateSubnetwork(ctx context.Context, networkPoolID string, input SubnetworkCreateInput) (*Subnetwork, error)
	GetSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) (*Subnetwork, error)
	UpdateSubnetwork(ctx context.Context, networkPoolID, subnetworkID string, input SubnetworkUpdateInput) (*Subnetwork, error)
	DeleteSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) error

	// Additional collections
	Subnetworks(networkPoolID string) Collection[Subnetwork]
}

// NetworkPoolsHandler handles operations around network pools
type NetworkPoolsHandler struct {
	client *Client
}

// Collection builds a new Collection[NetworkPool] interface
func (h *NetworkPoolsHandler) Collection() Collection[NetworkPool] {
	return NewCollection[NetworkPool](h.client, networkPoolListPath)
}

// Get returns a network pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingNetworkPool
func (h *NetworkPoolsHandler) Get(ctx context.Context, id string) (*NetworkPool, error) {
	url := h.client.buildURL(networkPoolPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	networkPool := new(NetworkPool)

	if err := json.Unmarshal(body, &networkPool); err != nil {
		return nil, err
	}

	return networkPool, nil
}

// Update returns updated network pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateAnExistingNetworkPool
func (h *NetworkPoolsHandler) Update(ctx context.Context, id string, input NetworkPoolInput) (*NetworkPool, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(networkPoolPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)
	if err != nil {
		return nil, err
	}

	networkPool := new(NetworkPool)

	if err := json.Unmarshal(body, &networkPool); err != nil {
		return nil, err
	}

	return networkPool, nil
}

// CreateSubnetwork returns created subnetwork from the pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateOrAllocateSubnetworkFromTheNetworkPool
func (h *NetworkPoolsHandler) CreateSubnetwork(ctx context.Context, networkPoolID string, input SubnetworkCreateInput) (*Subnetwork, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(subnetworkCreatePath, []interface{}{networkPoolID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}

	subnetwork := new(Subnetwork)

	if err := json.Unmarshal(body, &subnetwork); err != nil {
		return nil, err
	}

	return subnetwork, nil
}

// GetSubnetwork returns subnetwork from the pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingSubnetwork
func (h *NetworkPoolsHandler) GetSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) (*Subnetwork, error) {
	url := h.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	subnetwork := new(Subnetwork)

	if err := json.Unmarshal(body, &subnetwork); err != nil {
		return nil, err
	}

	return subnetwork, nil
}

// UpdateSubnetwork returns subnetwork from the pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateAnExistingSubnetwork
func (h *NetworkPoolsHandler) UpdateSubnetwork(ctx context.Context, networkPoolID, subnetworkID string, input SubnetworkUpdateInput) (*Subnetwork, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)
	if err != nil {
		return nil, err
	}

	subnetwork := new(Subnetwork)

	if err := json.Unmarshal(body, &subnetwork); err != nil {
		return nil, err
	}

	return subnetwork, nil
}

// DeleteSubnetwork delete subnetwork
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteAnExistingSubnetwork
func (h *NetworkPoolsHandler) DeleteSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) error {
	url := h.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// Subnetworks builds a new Collection[Subnetwork] interface
func (h *NetworkPoolsHandler) Subnetworks(networkPoolID string) Collection[Subnetwork] {
	path := h.client.buildPath(subnetworkListPath, []interface{}{networkPoolID}...)

	return NewCollection[Subnetwork](h.client, path)
}
