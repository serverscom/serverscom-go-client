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

	kubernetesClusterAutoscaleNodeGroupPath             = kubernetesClusterPathWithID + "/autoscale_node_groups"
	kubernetesClusterAutoscaleNodeGroupPathWithID       = kubernetesClusterAutoscaleNodeGroupPath + "/%s"
	kubernetesClusterAutoscaleNodeGroupTemplatePath     = kubernetesClusterAutoscaleNodeGroupPathWithID + "/autoscale_template"
	kubernetesClusterAutoscaleNodeGroupDecreaseSizePath = kubernetesClusterAutoscaleNodeGroupPathWithID + "/decrease_target_size"
	kubernetesClusterAutoscaleNodeGroupDeleteNodesPath  = kubernetesClusterAutoscaleNodeGroupPathWithID + "/delete_nodes"
	kubernetesClusterAutoscaleNodeGroupDisablePath      = kubernetesClusterAutoscaleNodeGroupPathWithID + "/disable"
	kubernetesClusterAutoscaleNodeGroupEnablePath       = kubernetesClusterAutoscaleNodeGroupPathWithID + "/enable"
	kubernetesClusterAutoscaleNodeGroupIncreaseSizePath = kubernetesClusterAutoscaleNodeGroupPathWithID + "/increase_size"
	kubernetesClusterAutoscaleNodeGroupNodesPath        = kubernetesClusterAutoscaleNodeGroupPathWithID + "/nodes"
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

	// autoscale node group operations
	GetAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error)
	CreateAutoscaleNodeGroup(ctx context.Context, clusterID string, input KubernetesClusterAutoscaleNodeGroupCreateInput) (*KubernetesClusterAutoscaleNodeGroup, error)
	UpdateAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupUpdateInput) (*KubernetesClusterAutoscaleNodeGroup, error)
	DeleteAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) error
	EnableAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error)
	DisableAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error)
	GetAutoscaleNodeGroupTemplate(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroupTemplate, error)
	DecreaseAutoscaleNodeGroupTargetSize(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupDecreaseTargetSizeInput) (*KubernetesClusterAutoscaleNodeGroup, error)
	IncreaseAutoscaleNodeGroupSize(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupIncreaseSizeInput) (*KubernetesClusterAutoscaleNodeGroup, error)
	DeleteAutoscaleNodeGroupNodes(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupDeleteNodesInput) (*KubernetesClusterAutoscaleNodeGroup, error)

	// Additional collections
	Nodes(id string) Collection[KubernetesClusterNode]
	NodeGroups(id string) Collection[KubernetesClusterNodeGroup]
	AutoscaleNodeGroups(clusterID string) Collection[KubernetesClusterAutoscaleNodeGroup]
	AutoscaleNodeGroupNodes(clusterID string, nodeGroupID string) Collection[KubernetesClusterNode]
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

// AutoscaleNodeGroups builds a new Collection[KubernetesClusterAutoscaleNodeGroup] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/KubernetesClusterAutoscaleNodeGroups
func (h *KubernetesClustersHandler) AutoscaleNodeGroups(clusterID string) Collection[KubernetesClusterAutoscaleNodeGroup] {
	path := h.client.buildPath(kubernetesClusterAutoscaleNodeGroupPath, []interface{}{clusterID}...)

	return NewCollection[KubernetesClusterAutoscaleNodeGroup](h.client, path)
}

// AutoscaleNodeGroupNodes builds a new Collection[KubernetesClusterNode] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ListTheNodesOfAnAutoscaleNodeGroup
func (h *KubernetesClustersHandler) AutoscaleNodeGroupNodes(clusterID string, nodeGroupID string) Collection[KubernetesClusterNode] {
	path := h.client.buildPath(kubernetesClusterAutoscaleNodeGroupNodesPath, []interface{}{clusterID, nodeGroupID}...)

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
// DeleteNodeGroup deletes a node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/DeleteANodeGroupForAKubernetesCluster
func (h *KubernetesClustersHandler) DeleteNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) error {
	url := h.client.buildURL(kubernetesClusterNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// GetAutoscaleNodeGroup gets an autoscale node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ShowAnAutoscaleNodeGroup
func (h *KubernetesClustersHandler) GetAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error) {
	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// CreateAutoscaleNodeGroup creates an autoscale node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/CreateANewAutoscaleNodeGroup
func (h *KubernetesClustersHandler) CreateAutoscaleNodeGroup(ctx context.Context, clusterID string, input KubernetesClusterAutoscaleNodeGroupCreateInput) (*KubernetesClusterAutoscaleNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupPath, []interface{}{clusterID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// UpdateAutoscaleNodeGroup updates an autoscale node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/UpdateAnAutoscaleNodeGroup
func (h *KubernetesClustersHandler) UpdateAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupUpdateInput) (*KubernetesClusterAutoscaleNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// DeleteAutoscaleNodeGroup deletes an autoscale node group for a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/DeleteAnAutoscaleNodeGroup
func (h *KubernetesClustersHandler) DeleteAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) error {
	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupPathWithID, []interface{}{clusterID, nodeGroupID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// EnableAutoscaleNodeGroup switches autoscaling on for an autoscale node group of a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/EnableAutoscalingForANodeGroup
func (h *KubernetesClustersHandler) EnableAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error) {
	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupEnablePath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// DisableAutoscaleNodeGroup switches autoscaling off for an autoscale node group of a Kubernetes cluster
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/DisableAutoscalingForANodeGroup
func (h *KubernetesClustersHandler) DisableAutoscaleNodeGroup(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroup, error) {
	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupDisablePath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// GetAutoscaleNodeGroupTemplate shows what a node of an autoscale node group looks like
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ShowWhatANodeOfTheGroupLooksLike
func (h *KubernetesClustersHandler) GetAutoscaleNodeGroupTemplate(ctx context.Context, clusterID string, nodeGroupID string) (*KubernetesClusterAutoscaleNodeGroupTemplate, error) {
	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupTemplatePath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var template KubernetesClusterAutoscaleNodeGroupTemplate
	if err := json.Unmarshal(body, &template); err != nil {
		return nil, err
	}

	return &template, nil
}

// DecreaseAutoscaleNodeGroupTargetSize lowers the target size of an autoscale node group without touching its nodes
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/LowerTheTargetSizeOfTheGroupWithoutTouchingItsNodes
func (h *KubernetesClustersHandler) DecreaseAutoscaleNodeGroupTargetSize(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupDecreaseTargetSizeInput) (*KubernetesClusterAutoscaleNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupDecreaseSizePath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// IncreaseAutoscaleNodeGroupSize raises the target size of an autoscale node group and orders the nodes
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/RaiseTheTargetSizeOfTheGroupAndOrderTheNodes
func (h *KubernetesClustersHandler) IncreaseAutoscaleNodeGroupSize(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupIncreaseSizeInput) (*KubernetesClusterAutoscaleNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupIncreaseSizePath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}

// DeleteAutoscaleNodeGroupNodes releases nodes of an autoscale node group and lowers the target size to match
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Cluster/operation/ReleaseNodesOfTheGroupAndLowerTheTargetSizeToMatch
func (h *KubernetesClustersHandler) DeleteAutoscaleNodeGroupNodes(ctx context.Context, clusterID string, nodeGroupID string, input KubernetesClusterAutoscaleNodeGroupDeleteNodesInput) (*KubernetesClusterAutoscaleNodeGroup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesClusterAutoscaleNodeGroupDeleteNodesPath, []interface{}{clusterID, nodeGroupID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var nodeGroup KubernetesClusterAutoscaleNodeGroup
	if err := json.Unmarshal(body, &nodeGroup); err != nil {
		return nil, err
	}

	return &nodeGroup, nil
}
