package serverscom

import (
	"context"
	"encoding/json"
)

const (
	loadBalancerListPath = "/load_balancers"

	l4LoadBalancerCreatePath = "/load_balancers/l4"
	l4LoadBalancerPath       = "/load_balancers/l4/%s"
	l4LoadBalancerUpdatePath = "/load_balancers/l4/%s"
	l4LoadBalancerDeletePath = "/load_balancers/l4/%s"
	l7LoadBalancerCreatePath = "/load_balancers/l7"
	l7LoadBalancerPath       = "/load_balancers/l7/%s"
	l7LoadBalancerUpdatePath = "/load_balancers/l7/%s"
	l7LoadBalancerDeletePath = "/load_balancers/l7/%s"
)

// LoadBalancersService is an interface for interfacing with Load balancers endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/LoadBalancers
type LoadBalancersService interface {
	// Primary collection
	Collection() Collection[LoadBalancer]

	// Generic operations
	GetL4LoadBalancer(ctx context.Context, id string) (*L4LoadBalancer, error)
	CreateL4LoadBalancer(ctx context.Context, input L4LoadBalancerCreateInput) (*L4LoadBalancer, error)
	UpdateL4LoadBalancer(ctx context.Context, id string, input L4LoadBalancerUpdateInput) (*L4LoadBalancer, error)
	DeleteL4LoadBalancer(ctx context.Context, id string) error
	GetL7LoadBalancer(ctx context.Context, id string) (*L7LoadBalancer, error)
	CreateL7LoadBalancer(ctx context.Context, input L7LoadBalancerCreateInput) (*L7LoadBalancer, error)
	UpdateL7LoadBalancer(ctx context.Context, id string, input L7LoadBalancerUpdateInput) (*L7LoadBalancer, error)
	DeleteL7LoadBalancer(ctx context.Context, id string) error
}

// LoadBalancersHandler handles operations around hosts
type LoadBalancersHandler struct {
	client *Client
}

// Collection builds a new Collection[LoadBalancer] interface
func (h *LoadBalancersHandler) Collection() Collection[LoadBalancer] {
	return NewCollection[LoadBalancer](h.client, loadBalancerListPath)
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

// DeleteL4LoadBalancer deletes l4 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/DeleteAnExistingL4LoadBalancer
func (h *LoadBalancersHandler) DeleteL4LoadBalancer(ctx context.Context, id string) error {
	url := h.client.buildURL(l4LoadBalancerDeletePath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// GetL7LoadBalancer returns a l7 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/RetrieveAnExistingL7LoadBalancer
func (h *LoadBalancersHandler) GetL7LoadBalancer(ctx context.Context, id string) (*L7LoadBalancer, error) {
	url := h.client.buildURL(l7LoadBalancerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L7LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// CreateL7LoadBalancer creates a l7 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/CreateANewL7LoadBalancer
func (h *LoadBalancersHandler) CreateL7LoadBalancer(ctx context.Context, input L7LoadBalancerCreateInput) (*L7LoadBalancer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l7LoadBalancerCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L7LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// UpdateL7LoadBalancer updates l7 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/UpdateAnExistingL7LoadBalancer
func (h *LoadBalancersHandler) UpdateL7LoadBalancer(ctx context.Context, id string, input L7LoadBalancerUpdateInput) (*L7LoadBalancer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(l7LoadBalancerUpdatePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	loadBalancer := new(L7LoadBalancer)

	if err := json.Unmarshal(body, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// DeleteL7LoadBalancer deletes l7 load balancer
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Load-Balancer/operation/DeleteAnExistingL7LoadBalancer
func (h *LoadBalancersHandler) DeleteL7LoadBalancer(ctx context.Context, id string) error {
	url := h.client.buildURL(l7LoadBalancerDeletePath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
