package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLoadBalancersCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.LoadBalancers.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCreateL4LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l4").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/load_balancers/l4/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	loadBalancerName := "name87"

	input := L4LoadBalancerCreateInput{
		Name:       loadBalancerName,
		LocationID: int64(1),
		Labels:     map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	loadBalancer, err := client.LoadBalancers.CreateL4LoadBalancer(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())

	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Status).To(Equal("in_process"))
	g.Expect(loadBalancer.Name).To(Equal("name87"))
	g.Expect(loadBalancer.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(loadBalancer.Created.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
}

func TestGetL4LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l4/y1aKReQG").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/load_balancers/l4/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	loadBalancer, err := client.LoadBalancers.GetL4LoadBalancer(ctx, "y1aKReQG")

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())

	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Status).To(Equal("active"))
	g.Expect(loadBalancer.Name).To(Equal("name87"))
	g.Expect(loadBalancer.Type).To(Equal("l4"))
	g.Expect(loadBalancer.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(loadBalancer.Created.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
}

func TestUpdateL4LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l4/y1aKReQG").
		WithRequestMethod("PUT").
		WithRequestBody(`{"name":"some","labels":{"env":"new-test"}}`).
		WithResponseBodyStubFile("fixtures/load_balancers/l4/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newName := "some"
	newLabels := map[string]string{"env": "new-test"}

	loadBalancer, err := client.LoadBalancers.UpdateL4LoadBalancer(ctx, "y1aKReQG", L4LoadBalancerUpdateInput{Name: &newName, Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())

	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Status).To(Equal("pending"))
	g.Expect(loadBalancer.Name).To(Equal("some"))
	g.Expect(loadBalancer.Type).To(Equal("l4"))
	g.Expect(loadBalancer.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancer.Labels).To(Equal(newLabels))
	g.Expect(loadBalancer.Created.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2022-09-13 12:00:11 +0000 UTC"))
}

func TestDeleteL4LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l4/y1aKReQG").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.LoadBalancers.DeleteL4LoadBalancer(ctx, "y1aKReQG")

	g.Expect(err).To(BeNil())
}

func TestCreateL7LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l7").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/load_balancers/l7/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := L7LoadBalancerCreateInput{
		Name:       "test-l7",
		LocationID: int64(1),
		Labels:     map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	loadBalancer, err := client.LoadBalancers.CreateL7LoadBalancer(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())
	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Name).To(Equal("name87"))
	g.Expect(loadBalancer.Type).To(Equal("l7"))
	g.Expect(loadBalancer.Status).To(Equal("active"))
	g.Expect(loadBalancer.Domains).To(ConsistOf("example.com", "www.example.com"))
	g.Expect(loadBalancer.ExternalAddresses).To(ConsistOf("10.0.0.1"))
	g.Expect(loadBalancer.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancer.Geoip).To(BeTrue())
	g.Expect(loadBalancer.StoreLogs).To(BeTrue())
	g.Expect(loadBalancer.StoreLogsRegionID).To(Equal(int64(2)))
	g.Expect(loadBalancer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(loadBalancer.Created.String()).To(Equal("2024-01-01 12:00:00 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2024-01-01 12:10:00 +0000 UTC"))
}

func TestGetL7LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l7/y1aKReQG").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/load_balancers/l7/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	loadBalancer, err := client.LoadBalancers.GetL7LoadBalancer(ctx, "y1aKReQG")

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())
	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Name).To(Equal("name87"))
	g.Expect(loadBalancer.Status).To(Equal("active"))
	g.Expect(loadBalancer.StoreLogs).To(BeTrue())
	g.Expect(loadBalancer.Geoip).To(BeFalse())
	g.Expect(loadBalancer.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(loadBalancer.Created.String()).To(Equal("2024-01-01 12:00:00 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2024-01-02 12:10:00 +0000 UTC"))
}

func TestUpdateL7LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l7/y1aKReQG").
		WithRequestMethod("PUT").
		WithRequestBody(`{"name":"some","labels":{"env":"new-test"}}`).
		WithResponseBodyStubFile("fixtures/load_balancers/l7/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newName := "some"
	newLabels := map[string]string{"env": "new-test"}

	loadBalancer, err := client.LoadBalancers.UpdateL7LoadBalancer(ctx, "y1aKReQG", L7LoadBalancerUpdateInput{Name: newName, Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(loadBalancer).ToNot(BeNil())
	g.Expect(loadBalancer.ID).To(Equal("y1aKReQG"))
	g.Expect(loadBalancer.Name).To(Equal(newName))
	g.Expect(loadBalancer.Type).To(Equal("l7"))
	g.Expect(loadBalancer.Status).To(Equal("pending"))
	g.Expect(loadBalancer.Domains).To(ConsistOf("updated.example.com"))
	g.Expect(loadBalancer.StoreLogs).To(BeTrue())
	g.Expect(loadBalancer.StoreLogsRegionID).To(Equal(int64(3)))
	g.Expect(loadBalancer.ExternalAddresses).To(ConsistOf("10.0.0.1"))
	g.Expect(loadBalancer.LocationID).To(Equal(int64(1)))
	g.Expect(loadBalancer.Geoip).To(BeFalse())
	g.Expect(loadBalancer.Labels).To(Equal(newLabels))
	g.Expect(loadBalancer.Created.String()).To(Equal("2024-01-01 12:00:00 +0000 UTC"))
	g.Expect(loadBalancer.Updated.String()).To(Equal("2024-01-02 12:10:00 +0000 UTC"))
}

func TestDeleteL7LoadBalancer(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/load_balancers/l7/y1aKReQG").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.LoadBalancers.DeleteL7LoadBalancer(ctx, "y1aKReQG")

	g.Expect(err).To(BeNil())
}
