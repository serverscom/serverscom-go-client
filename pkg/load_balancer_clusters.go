package serverscom

import (
	"context"
	"encoding/json"
)

const (
	loadBalancerClusterListPath = "/load_balancer_clusters"

	LoadBalancerClusterPath = "/load_balancer_clusters/%s"
)

// LoadBalancersService is an interface for interfacing with Load balancers endpoints
type LoadBalancerClustersService interface {
	// Primary collection
	Collection() Collection[LoadBalancerCluster]

	// Generic operations
	GetLoadBalancerCluster(ctx context.Context, id string) (*LoadBalancerCluster, error)
}

// LoadBalancersHandler handles operations around hosts
type LoadBalancerClustersHandler struct {
	client *Client
}

// Collection builds a new Collection[LoadBalancer] interface
func (h *LoadBalancerClustersHandler) Collection() Collection[LoadBalancerCluster] {
	return NewCollection[LoadBalancerCluster](h.client, loadBalancerClusterListPath)
}

// GetLoadBalancerCluster returns a load balancer cluster
func (h *LoadBalancerClustersHandler) GetLoadBalancerCluster(ctx context.Context, id string) (*LoadBalancerCluster, error) {
	url := h.client.buildURL(LoadBalancerClusterPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	loadBalancerCluster := new(LoadBalancerCluster)

	if err := json.Unmarshal(body, &loadBalancerCluster); err != nil {
		return nil, err
	}

	return loadBalancerCluster, nil
}
