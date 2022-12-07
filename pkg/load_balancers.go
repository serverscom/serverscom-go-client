package serverscom

import (
	"context"
	"encoding/json"
)

const (
	l4LoadBalancerCreatePath = "/load_balancers/l4"
	l4LoadBalancerPath       = "/load_balancers/l4/%s"
	l4LoadBalancerUpdatePath = "/load_balancers/l4/%s"
	l4LoadBalancerDeletePath = "/load_balancers/l4/%s"
)

// LoadBalancersService is an interface for interfacing with Load balancers endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/LoadBalancers
type LoadBalancersService interface {
	// Primary collection
	Collection() LoadBalancersCollection

	// Generic operations
	GetL4LoadBalancer(ctx context.Context, id string) (*L4LoadBalancer, error)
	CreateL4LoadBalancer(ctx context.Context, input L4LoadBalancerCreateInput) (*L4LoadBalancer, error)
	UpdateL4LoadBalancer(ctx context.Context, id string, input L4LoadBalancerUpdateInput) (*L4LoadBalancer, error)
	DeleteL4LoadBalancer(ctx context.Context, id string) error
}

// LoadBalancersHandler handles operations around hosts
type LoadBalancersHandler struct {
	client *Client
}

// Collection builds a new HostsCollection interface
func (h *LoadBalancersHandler) Collection() LoadBalancersCollection {
	return NewLoadBalancersCollection(h.client)
}

// GetL4LoadBalancer returns a l4 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/RetrieveAnExisitingL4LoadBalancer
func (h *LoadBalancersHandler) GetL4LoadBalancer(ctx context.Context, id string) (*L4LoadBalancer, error) {
	url := h.client.buildURL(l4LoadBalancerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L4LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// CreateL4LoadBalancer creates a l4 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/CreateANewL4LoadBalancer
func (h *LoadBalancersHandler) CreateL4LoadBalancer(ctx context.Context, input L4LoadBalancerCreateInput) (*L4LoadBalancer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l4LoadBalancerCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L4LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// UpdateL4LoadBalancer updates l4 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/UpdateAnExisitingL4LoadBalancer
func (h *LoadBalancersHandler) UpdateL4LoadBalancer(ctx context.Context, id string, input L4LoadBalancerUpdateInput) (*L4LoadBalancer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l4LoadBalancerUpdatePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L4LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// DeleteL4LoadBalancer deltes l4 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/DeleteAnExistingL4LoadBalancer
func (h *LoadBalancersHandler) DeleteL4LoadBalancer(ctx context.Context, id string) error {
	url := h.client.buildURL(l4LoadBalancerDeletePath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
