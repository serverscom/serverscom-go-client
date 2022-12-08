package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestNetworkPoolsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.NetworkPools.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestNetworkPoolsGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/network_pools/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	networkPool, err := client.NetworkPools.Get(ctx, "a")

	g.Expect(err).To(BeNil())
	g.Expect(networkPool).ToNot(BeNil())

	g.Expect(networkPool.ID).To(Equal("a"))
	g.Expect(*networkPool.Title).To(Equal("new-network-pool-title"))
	g.Expect(networkPool.CIDR).To(Equal("10.0.0.0/20"))
	g.Expect(networkPool.Type).To(Equal("private"))
	g.Expect(networkPool.Created.String()).To(Equal("2021-03-24 11:46:35 +0000 UTC"))
	g.Expect(networkPool.Updated.String()).To(Equal("2021-03-24 11:46:35 +0000 UTC"))
}

func TestNetworkPoolsUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a").
		WithRequestMethod("PUT").
		WithRequestBody(`{"title":"some"}`).
		WithResponseBodyStubFile("fixtures/network_pools/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newTitle := "some"

	networkPool, err := client.NetworkPools.Update(ctx, "a", NetworkPoolInput{Title: &newTitle})

	g.Expect(err).To(BeNil())
	g.Expect(networkPool).ToNot(BeNil())

	g.Expect(networkPool.ID).To(Equal("a"))
	g.Expect(*networkPool.Title).To(Equal("some"))
	g.Expect(networkPool.CIDR).To(Equal("10.0.0.0/20"))
	g.Expect(networkPool.Type).To(Equal("private"))
	g.Expect(networkPool.Created.String()).To(Equal("2021-03-24 11:46:37 +0000 UTC"))
	g.Expect(networkPool.Updated.String()).To(Equal("2021-03-24 11:46:38 +0000 UTC"))
}

func TestNetworkPoolsCreateSubnetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a/subnetworks").
		WithRequestMethod("POST").
		WithRequestBody(`{"title":"some","cidr":"100.10.32.0/29","mask":29}`).
		WithResponseBodyStubFile("fixtures/network_pools/create_subnetwork_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newTitle := "some"
	newCIDR := "100.10.32.0/29"
	newMask := 29

	subnetwork, err := client.NetworkPools.CreateSubnetwork(ctx, "a", SubnetworkCreateInput{
		Title: &newTitle,
		CIDR:  &newCIDR,
		Mask:  &newMask,
	})

	g.Expect(err).To(BeNil())
	g.Expect(subnetwork).ToNot(BeNil())

	g.Expect(subnetwork.ID).To(Equal("b"))
	g.Expect(subnetwork.NetworkPoolID).To(Equal("a"))
	g.Expect(*subnetwork.Title).To(Equal("some"))
	g.Expect(subnetwork.CIDR).To(Equal("100.10.32.0/29"))
	g.Expect(subnetwork.InterfaceType).To(Equal("public"))
	g.Expect(subnetwork.Attached).To(Equal(false))
	g.Expect(subnetwork.Created.String()).To(Equal("2021-03-24 11:46:55 +0000 UTC"))
	g.Expect(subnetwork.Updated.String()).To(Equal("2021-03-24 11:46:55 +0000 UTC"))
}

func TestNetworkPoolsGetSubnetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a/subnetworks/b").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/network_pools/get_subnetwork_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	subnetwork, err := client.NetworkPools.GetSubnetwork(ctx, "a", "b")

	g.Expect(err).To(BeNil())
	g.Expect(subnetwork).ToNot(BeNil())

	g.Expect(subnetwork.ID).To(Equal("b"))
	g.Expect(subnetwork.NetworkPoolID).To(Equal("a"))
	g.Expect(*subnetwork.Title).To(Equal("Public"))
	g.Expect(subnetwork.CIDR).To(Equal("100.10.32.0/29"))
	g.Expect(subnetwork.InterfaceType).To(Equal("public"))
	g.Expect(subnetwork.Attached).To(Equal(true))
	g.Expect(subnetwork.Created.String()).To(Equal("2021-03-24 11:46:53 +0000 UTC"))
	g.Expect(subnetwork.Updated.String()).To(Equal("2021-03-24 11:46:53 +0000 UTC"))
}

func TestNetworkPoolsUpdateSubnetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a/subnetworks/b").
		WithRequestMethod("PUT").
		WithRequestBody(`{"title":"some"}`).
		WithResponseBodyStubFile("fixtures/network_pools/update_subnetwork_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newTitle := "some"

	subnetwork, err := client.NetworkPools.UpdateSubnetwork(ctx, "a", "b", SubnetworkUpdateInput{Title: &newTitle})

	g.Expect(err).To(BeNil())
	g.Expect(subnetwork).ToNot(BeNil())

	g.Expect(subnetwork.ID).To(Equal("b"))
	g.Expect(subnetwork.NetworkPoolID).To(Equal("a"))
	g.Expect(*subnetwork.Title).To(Equal("some"))
	g.Expect(subnetwork.CIDR).To(Equal("100.10.32.0/29"))
	g.Expect(subnetwork.InterfaceType).To(Equal("public"))
	g.Expect(subnetwork.Attached).To(Equal(true))
	g.Expect(subnetwork.Created.String()).To(Equal("2021-03-24 11:46:56 +0000 UTC"))
	g.Expect(subnetwork.Updated.String()).To(Equal("2021-03-24 11:46:56 +0000 UTC"))
}

func TestNetworkPoolsDeleteSubnetwork(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a/subnetworks/b").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.NetworkPools.DeleteSubnetwork(ctx, "a", "b")

	g.Expect(err).To(BeNil())
}

func TestSubnetworksCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/network_pools/a/subnetworks").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.NetworkPools.Subnetworks("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}
