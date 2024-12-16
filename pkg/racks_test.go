package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	testRackID = "MYer7LaO"
)

func TestRacksCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/racks").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.Racks.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRacksGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/racks/" + testRackID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/racks/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	rack, err := client.Racks.Get(ctx, testRackID)

	g.Expect(err).To(BeNil())
	g.Expect(rack).ToNot(BeNil())

	g.Expect(rack.ID).To(Equal(testRackID))
	g.Expect(rack.Name).To(Equal("test-rack-1"))
	g.Expect(rack.LocationID).To(Equal(int64(1)))
	g.Expect(rack.LocationCode).To(Equal("location-1"))
	g.Expect(rack.Labels).To(Equal(map[string]string{"env": "test"}))
}

func TestRacksUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/racks/" + testRackID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/racks/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	rack, err := client.Racks.Update(ctx, testRackID, RackUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(rack).ToNot(BeNil())

	g.Expect(rack.ID).To(Equal(testRackID))
	g.Expect(rack.Name).To(Equal("test-rack-1"))
	g.Expect(rack.LocationID).To(Equal(int64(1)))
	g.Expect(rack.LocationCode).To(Equal("location-1"))
	g.Expect(rack.Labels).To(Equal(newLabels))
}
