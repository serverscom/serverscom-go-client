package serverscom

import (
	"context"
	"encoding/json"
)

const (
	kubernetesClusterPath           = "/kubernetes_clusters"
	kubernetesClusterPathWithID     = kubernetesClusterPath + "/%s"
	kubernetesClusterNodePath       = kubernetesClusterPathWithID + "/nodes"
	kubernetesClusterNodePathWithID = kubernetesClusterNodePath + "/%s"

	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes
	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes/{node_id}
)

// KubernetesClustersService is an interface for interfacing with Kubernetes Cluster endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster
type KubernetesClustersService interface {
	// Primary collection
	Collection() Collection[KubernetesCluster]

	// Generic operations
	Get(ctx context.Context, id string) (*KubernetesCluster, error)
	GetNode(ctx context.Context, clusterID string, nodeID string) (*KubernetesClusterNode, error)
	Update(ctx context.Context, id string, input KubernetesClusterUpdateInput) (*KubernetesCluster, error)

	// Additional collections
	Nodes(id string) Collection[KubernetesClusterNode]
}

// KubernetesClustersHandler handles operations around kubernetes clusters
type KubernetesClustersHandler struct {
	client *Client
}

// Collection builds a new Collection[KubernetesCluster] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ListKubernetesClusters
func (h *KubernetesClustersHandler) Collection() Collection[KubernetesCluster] {
	return NewCollection[KubernetesCluster](h.client, kubernetesClusterPath)
}

// Nodes builds a new Collection[KubernetesClusterNode] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ListNodesForAKubernetesCluster
func (h *KubernetesClustersHandler) Nodes(id string) Collection[KubernetesClusterNode] {
	path := h.client.buildPath(kubernetesClusterNodePath, []interface{}{id}...)

	return NewCollection[KubernetesClusterNode](h.client, path)
}

// Get a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/GetAKubernetesCluster
func (h *KubernetesClustersHandler) Get(ctx context.Context, id string) (*KubernetesCluster, error) {
	url := h.client.buildURL(kubernetesClusterPathWithID, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var cluster KubernetesCluster
	if err := json.Unmarshal(body, &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}

// Get a node for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/GetANodeForAKubernetesCluster
func (h *KubernetesClustersHandler) GetNode(ctx context.Context, clusterID string, nodeID string) (*KubernetesClusterNode, error) {
	url := h.client.buildURL(kubernetesClusterNodePathWithID, []interface{}{clusterID, nodeID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var node KubernetesClusterNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// Update a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/UpdateAKubernetesCluster
func (h *KubernetesClustersHandler) Update(ctx context.Context, id string, input KubernetesClusterUpdateInput) (*KubernetesCluster, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterPathWithID, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var cluster KubernetesCluster
	if err := json.Unmarshal(body, &cluster); err != nil {
		return nil, err
	}

	return &cluster, nil
}
