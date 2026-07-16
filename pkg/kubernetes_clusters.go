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
	kubernetesClusterNodeMovePath   = kubernetesClusterNodePath + "/move"

	kubernetesClusterNodeGroupPath       = kubernetesClusterPathWithID + "/node_groups"
	kubernetesClusterNodeGroupPathWithID = kubernetesClusterNodeGroupPath + "/%s"

	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes
	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes/{node_id}
	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/nodes/move
	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/node_groups
	// /v1/kubernetes_clusters/{kubernetes_cluster_id}/node_groups/{node_group_id}
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

	// node operations
	MoveNodes(ctx context.Context, clusterID string, input KubernetesClusterMoveNodesInput) ([]KubernetesClusterNode, error)
	UpdateNode(ctx context.Context, clusterID string, nodeID string, input KubernetesClusterNodeUpdateInput) (*KubernetesClusterNode, error)

	// node group operations
	CreateNodeGroup(ctx context.Context, clusterID string, input KubernetesClusterNodeGroupCreateInput) (*KubernetesClusterNodeGroup, error)
	UpdateNodeGroup(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterNodeGroupUpdateInput) (*KubernetesClusterNodeGroup, error)
	DeleteNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) error

	// Additional collections
	Nodes(id string) Collection[KubernetesClusterNode]
	NodeGroups(id string) Collection[KubernetesClusterNodeGroup]
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

// NodeGroups builds a new Collection[KubernetesClusterNodeGroup] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ListNodeGroupsForAKubernetesCluster
func (h *KubernetesClustersHandler) NodeGroups(id string) Collection[KubernetesClusterNodeGroup] {
	path := h.client.buildPath(kubernetesClusterNodeGroupPath, []interface{}{id}...)

	return NewCollection[KubernetesClusterNodeGroup](h.client, path)
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

// MoveNodes moves nodes to a group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/MoveNodesToAGroupForAKubernetesCluster
func (h *KubernetesClustersHandler) MoveNodes(ctx context.Context, clusterID string, input KubernetesClusterMoveNodesInput) ([]KubernetesClusterNode, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterNodeMovePath, []interface{}{clusterID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodes []KubernetesClusterNode
	if err := json.Unmarshal(body, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

// UpdateNode updates a node for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/UpdateANodeForAKubernetesCluster
func (h *KubernetesClustersHandler) UpdateNode(ctx context.Context, clusterID string, nodeID string, input KubernetesClusterNodeUpdateInput) (*KubernetesClusterNode, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterNodePathWithID, []interface{}{clusterID, nodeID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var node KubernetesClusterNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// CreateNodeGroup creates a node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/CreateANodeGroupForAKubernetesCluster
func (h *KubernetesClustersHandler) CreateNodeGroup(ctx context.Context, clusterID string, input KubernetesClusterNodeGroupCreateInput) (*KubernetesClusterNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterNodeGroupPath, []interface{}{clusterID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// UpdateNodeGroup updates a node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/UpdateANodeGroupForAKubernetesCluster
func (h *KubernetesClustersHandler) UpdateNodeGroup(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterNodeGroupUpdateInput) (*KubernetesClusterNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// DeleteNodeGroup deletes a node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/DeleteANodeGroupForAKubernetesCluster
func (h *KubernetesClustersHandler) DeleteNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) error {
	url := h.client.buildURL(kubernetesClusterNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
