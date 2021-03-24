package serverscom

import (
	"context"
	"encoding/json"
)

const (
	networkPoolPath = "/network_pools/%s"

	subnetworkCreatePath = "/network_pools/%s/subnetworks"
	subnetworkPath       = "/network_pools/%s/subnetworks/%s"
)

// NetworkPoolsService is an interface to interfacing with the Network Pool endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Network-Pool
type NetworkPoolsService interface {
	// Primary collection
	Collection() NetworkPoolsCollection

	// Generic operations
	Get(ctx context.Context, id string) (*NetworkPool, error)
	Update(ctx context.Context, id string, input NetworkPoolInput) (*NetworkPool, error)

	CreateSubnetwork(ctx context.Context, networkPoolID string, input SubnetworkCreateInput) (*Subnetwork, error)
	GetSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) (*Subnetwork, error)
	UpdateSubnetwork(ctx context.Context, networkPoolID, subnetworkID string, input SubnetworkUpdateInput) (*Subnetwork, error)
	DeleteSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) error

	// Additional collections
	Subnetworks(networkPoolID string) SubnetworksCollection
}

// NetworkPoolsHandler handles operations around network pools
type NetworkPoolsHandler struct {
	client *Client
}

// Collection builds a new NetworkPoolsCollection interface
func (resource *NetworkPoolsHandler) Collection() NetworkPoolsCollection {
	return NewNetworkPoolsCollection(resource.client)
}

// Get returns a network pool
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingNetworkPool
func (resource *NetworkPoolsHandler) Get(ctx context.Context, id string) (*NetworkPool, error) {
	url := resource.client.buildURL(networkPoolPath, []interface{}{id}...)

	body, err := resource.client.buildAndExecRequest(ctx, "GET", url, nil)
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
func (resource *NetworkPoolsHandler) Update(ctx context.Context, id string, input NetworkPoolInput) (*NetworkPool, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := resource.client.buildURL(networkPoolPath, []interface{}{id}...)

	body, err := resource.client.buildAndExecRequest(ctx, "PUT", url, payload)
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
func (resource *NetworkPoolsHandler) CreateSubnetwork(ctx context.Context, networkPoolID string, input SubnetworkCreateInput) (*Subnetwork, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := resource.client.buildURL(subnetworkCreatePath, []interface{}{networkPoolID}...)

	body, err := resource.client.buildAndExecRequest(ctx, "POST", url, payload)
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
func (resource *NetworkPoolsHandler) GetSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) (*Subnetwork, error) {
	url := resource.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	body, err := resource.client.buildAndExecRequest(ctx, "GET", url, nil)
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
func (resource *NetworkPoolsHandler) UpdateSubnetwork(ctx context.Context, networkPoolID, subnetworkID string, input SubnetworkUpdateInput) (*Subnetwork, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := resource.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	body, err := resource.client.buildAndExecRequest(ctx, "PUT", url, payload)
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
func (resource *NetworkPoolsHandler) DeleteSubnetwork(ctx context.Context, networkPoolID, subnetworkID string) error {
	url := resource.client.buildURL(subnetworkPath, []interface{}{networkPoolID, subnetworkID}...)

	_, err := resource.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// Subnetworks builds a new SubnetworksCollection interface
func (resource *NetworkPoolsHandler) Subnetworks(networkPoolID string) SubnetworksCollection {
	return NewSubnetworksCollection(resource.client, networkPoolID)
}
