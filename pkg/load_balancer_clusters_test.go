package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLoadBalancerClusterCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancer_clusters").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.LoadBalancerClusters.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestGetLoadBalancerCluster(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancer_clusters/Jrb2XMeW").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/load_balancer_clusters/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	loadBalancerCluster, err := client.LoadBalancerClusters.GetLoadBalancerCluster(ctx, "Jrb2XMeW")

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancerCluster).ToNot(BeNil())

	g.Expect(loadBalancerCluster.ID).To(Equal("Jrb2XMeW"))
	g.Expect(loadBalancerCluster.Status).To(Equal("active"))
	g.Expect(loadBalancerCluster.Name).To(Equal("dedic-test-ams1"))
	g.Expect(loadBalancerCluster.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancerCluster.Created.String()).To(Equal("2024-10-17 08:38:07 +0000 UTC"))
	g.Expect(loadBalancerCluster.Updated.String()).To(Equal("2024-10-17 08:39:40 +0000 UTC"))
}
