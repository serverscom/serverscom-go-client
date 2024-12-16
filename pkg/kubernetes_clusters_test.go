package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	clusterID = "YQdJqobO"
	nodeID    = "MYer06bO"
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
