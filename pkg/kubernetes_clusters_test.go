package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	clusterID            = "YQdJqobO"
	nodeID               = "MYer06bO"
	nodeGroupID          = "node-group-id"
	autoscaleNodeGroupID = "autoscale-node-group-id"
)

func TestKubernetesClusterCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.KubernetesClusters.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestKubernetesClusterGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cluster, err := client.KubernetesClusters.Get(ctx, clusterID)

	g.Expect(err).To(BeNil())
	g.Expect(cluster).ToNot(BeNil())

	g.Expect(cluster.ID).To(Equal(clusterID))
	g.Expect(cluster.Status).To(Equal("pending"))
	g.Expect(cluster.Name).To(Equal("k8s-cluster"))
	g.Expect(cluster.LocationID).To(Equal(int64(1)))
	g.Expect(cluster.LocationCode).To(Equal("location2155"))
	g.Expect(cluster.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cluster.Created.String()).To(Equal("2024-11-11 09:59:01 +0000 UTC"))
	g.Expect(cluster.Updated.String()).To(Equal("2024-11-11 09:59:01 +0000 UTC"))
}

func TestKubernetesClusterUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	cluster, err := client.KubernetesClusters.Update(ctx, clusterID, KubernetesClusterUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(cluster).ToNot(BeNil())

	g.Expect(cluster.ID).To(Equal(clusterID))
	g.Expect(cluster.Status).To(Equal("pending"))
	g.Expect(cluster.Name).To(Equal("k8s-cluster"))
	g.Expect(cluster.LocationID).To(Equal(int64(1)))
	g.Expect(cluster.LocationCode).To(Equal("location2155"))
	g.Expect(cluster.Labels).To(Equal(newLabels))
	g.Expect(cluster.Created.String()).To(Equal("2024-11-11 09:59:01 +0000 UTC"))
	g.Expect(cluster.Updated.String()).To(Equal("2024-11-11 09:59:01 +0000 UTC"))
}

func TestKubernetesClusterNodesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/nodes").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.KubernetesClusters.Nodes(clusterID)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestKubernetesClusterNodeGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/nodes/" + nodeID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/get_node_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cluster, err := client.KubernetesClusters.GetNode(ctx, clusterID, nodeID)

	g.Expect(err).To(BeNil())
	g.Expect(cluster).ToNot(BeNil())

	g.Expect(cluster.ID).To(Equal(nodeID))
	g.Expect(cluster.Number).To(Equal(int64(49)))
	g.Expect(cluster.Hostname).To(Equal("name585"))
	g.Expect(cluster.Configuration).To(Equal("SSD.50"))
	g.Expect(cluster.Type).To(Equal("cloud"))
	g.Expect(cluster.Role).To(Equal("node"))
	g.Expect(cluster.Status).To(Equal("pending"))
	g.Expect(cluster.PrivateIPv4Address).To(Equal("127.0.3.1"))
	g.Expect(cluster.PublicIPv4Address).To(Equal("127.0.5.2"))
	g.Expect(cluster.RefID).To(Equal("y5eVMdEP"))
	g.Expect(cluster.ClusterID).To(Equal(clusterID))
	g.Expect(cluster.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cluster.Created.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
	g.Expect(cluster.Updated.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
}

func TestKubernetesClusterNodeGroupsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/node_groups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.KubernetesClusters.NodeGroups(clusterID)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestKubernetesClusterCreateNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/node_groups").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/node_group_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.CreateNodeGroup(ctx, clusterID, KubernetesClusterNodeGroupCreateInput{
		Name:        "workers",
		Description: "worker node group",
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(nodeGroupID))
	g.Expect(nodeGroup.Name).To(Equal("workers"))
	g.Expect(*nodeGroup.Description).To(Equal("worker node group"))
	g.Expect(nodeGroup.NodeCount).To(Equal(int64(3)))
	g.Expect(nodeGroup.Created.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
	g.Expect(nodeGroup.Updated.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
}

func TestKubernetesClusterUpdateNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/node_groups/" + nodeGroupID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/node_group_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.UpdateNodeGroup(ctx, clusterID, nodeGroupID, KubernetesClusterNodeGroupUpdateInput{
		Name:        "workers",
		Description: "worker node group",
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(nodeGroupID))
	g.Expect(nodeGroup.Name).To(Equal("workers"))
	g.Expect(*nodeGroup.Description).To(Equal("worker node group"))
	g.Expect(nodeGroup.NodeCount).To(Equal(int64(3)))
}

func TestKubernetesClusterDeleteNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/node_groups/" + nodeGroupID).
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.KubernetesClusters.DeleteNodeGroup(ctx, clusterID, nodeGroupID)

	g.Expect(err).To(BeNil())
}

func TestKubernetesClusterMoveNodes(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/nodes/move").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/move_nodes_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodes, err := client.KubernetesClusters.MoveNodes(ctx, clusterID, KubernetesClusterMoveNodesInput{
		NodeGroupID: nodeGroupID,
		NodeIDs:     []string{nodeID},
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodes).To(HaveLen(1))
	g.Expect(nodes[0].ID).To(Equal(nodeID))
	g.Expect(nodes[0].NodeGroup.ID).To(Equal(nodeGroupID))
	g.Expect(nodes[0].NodeGroup.Name).To(Equal("workers"))
	g.Expect(nodes[0].NodeGroup.Type).To(Equal("static"))
}

func TestKubernetesClusterUpdateNode(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/nodes/" + nodeID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/update_node_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	node, err := client.KubernetesClusters.UpdateNode(ctx, clusterID, nodeID, KubernetesClusterNodeUpdateInput{
		NodeGroupID: nodeGroupID,
	})

	g.Expect(err).To(BeNil())
	g.Expect(node).ToNot(BeNil())
	g.Expect(node.ID).To(Equal(nodeID))
	g.Expect(node.NodeGroup.ID).To(Equal(nodeGroupID))
	g.Expect(node.NodeGroup.Name).To(Equal("workers"))
	g.Expect(node.NodeGroup.Type).To(Equal("static"))
}

func TestKubernetesClusterAutoscaleNodeGroupsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.KubernetesClusters.AutoscaleNodeGroups(clusterID)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

// The autoscale_enabled filter is a plain query param on the listing endpoint, so it has to
// survive the way a collection assembles its URL.
func TestKubernetesClusterAutoscaleNodeGroupsCollectionFilteredByAutoscaleEnabled(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups").
		WithRequestMethod("GET").
		WithRequestParams("autoscale_enabled=true").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	list, err := client.KubernetesClusters.AutoscaleNodeGroups(clusterID).
		SetParam("autoscale_enabled", "true").
		List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
}

func TestKubernetesClusterCreateAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.CreateAutoscaleNodeGroup(ctx, clusterID, KubernetesClusterAutoscaleNodeGroupCreateInput{
		Name:     "autoscale-workers",
		NodeType: "sbm",
		MinNodes: 1,
		MaxNodes: 5,
		NodeSpec: map[string]any{"flavor_id": "flavor-id"},
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
	g.Expect(nodeGroup.Name).To(Equal("autoscale-workers"))
	g.Expect(*nodeGroup.Description).To(Equal("autoscale worker node group"))
	g.Expect(nodeGroup.Type).To(Equal("autoscale"))
	g.Expect(nodeGroup.NodeType).To(Equal("sbm"))
	g.Expect(nodeGroup.MinNodes).To(Equal(int64(1)))
	g.Expect(nodeGroup.MaxNodes).To(Equal(int64(5)))
	g.Expect(nodeGroup.TargetNodes).To(Equal(int64(2)))
	g.Expect(nodeGroup.CurrentNodes).To(Equal(int64(2)))
	g.Expect(nodeGroup.Created.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
	g.Expect(nodeGroup.Updated.String()).To(Equal("2024-11-11 09:57:56 +0000 UTC"))
}

func TestKubernetesClusterGetAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.GetAutoscaleNodeGroup(ctx, clusterID, autoscaleNodeGroupID)

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
	g.Expect(nodeGroup.Name).To(Equal("autoscale-workers"))
	g.Expect(nodeGroup.AutoscaleEnabled).To(Equal(true))
}

func TestKubernetesClusterUpdateAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.UpdateAutoscaleNodeGroup(ctx, clusterID, autoscaleNodeGroupID, KubernetesClusterAutoscaleNodeGroupUpdateInput{
		Name:     "autoscale-workers",
		MaxNodes: 5,
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
	g.Expect(nodeGroup.Name).To(Equal("autoscale-workers"))
}

func TestKubernetesClusterDeleteAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID).
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.KubernetesClusters.DeleteAutoscaleNodeGroup(ctx, clusterID, autoscaleNodeGroupID)

	g.Expect(err).To(BeNil())
}

func TestKubernetesClusterEnableAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/enable").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.EnableAutoscaleNodeGroup(ctx, clusterID, autoscaleNodeGroupID)

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
	g.Expect(nodeGroup.AutoscaleEnabled).To(Equal(true))
}

func TestKubernetesClusterDisableAutoscaleNodeGroup(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/disable").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_disabled_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.DisableAutoscaleNodeGroup(ctx, clusterID, autoscaleNodeGroupID)

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
	g.Expect(nodeGroup.AutoscaleEnabled).To(Equal(false))
}

func TestKubernetesClusterGetAutoscaleNodeGroupTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/autoscale_template").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_template_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	template, err := client.KubernetesClusters.GetAutoscaleNodeGroupTemplate(ctx, clusterID, autoscaleNodeGroupID)

	g.Expect(err).To(BeNil())
	g.Expect(template).ToNot(BeNil())
	g.Expect(template.FlavorName).To(Equal("SBM-Autoscale-Flavor"))
	g.Expect(*template.LogicalCPUCount).To(Equal(int64(32)))
	g.Expect(*template.RAMSize).To(Equal(int64(128)))
}

func TestKubernetesClusterDecreaseAutoscaleNodeGroupTargetSize(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/decrease_target_size").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.DecreaseAutoscaleNodeGroupTargetSize(ctx, clusterID, autoscaleNodeGroupID, KubernetesClusterAutoscaleNodeGroupDecreaseTargetSizeInput{
		Delta: 1,
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
}

func TestKubernetesClusterIncreaseAutoscaleNodeGroupSize(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/increase_size").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.IncreaseAutoscaleNodeGroupSize(ctx, clusterID, autoscaleNodeGroupID, KubernetesClusterAutoscaleNodeGroupIncreaseSizeInput{
		Delta: 1,
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
}

func TestKubernetesClusterDeleteAutoscaleNodeGroupNodes(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/delete_nodes").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/kubernetes_clusters/autoscale_node_group_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	nodeGroup, err := client.KubernetesClusters.DeleteAutoscaleNodeGroupNodes(ctx, clusterID, autoscaleNodeGroupID, KubernetesClusterAutoscaleNodeGroupDeleteNodesInput{
		NodeIDs: []string{nodeID},
	})

	g.Expect(err).To(BeNil())
	g.Expect(nodeGroup).ToNot(BeNil())
	g.Expect(nodeGroup.ID).To(Equal(autoscaleNodeGroupID))
}

func TestKubernetesClusterAutoscaleNodeGroupNodesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/kubernetes_clusters/" + clusterID + "/autoscale_node_groups/" + autoscaleNodeGroupID + "/nodes").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.KubernetesClusters.AutoscaleNodeGroupNodes(clusterID, autoscaleNodeGroupID)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}
