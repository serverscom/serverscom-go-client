package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestL2SegmentsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.L2Segments.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2SegmentsCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/l2_segments/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	l2SegmentName := "name87"

	input := L2SegmentCreateInput{
		Name:            &l2SegmentName,
		Type:            "public",
		LocationGroupID: int64(1),
		Members: []L2SegmentMemberInput{
			{ID: "a", Mode: "native"},
			{ID: "b", Mode: "trunk"},
		},
	}

	ctx := context.TODO()

	L2Segments, err := client.L2Segments.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(L2Segments).ToNot(BeNil())

	g.Expect(L2Segments.ID).To(Equal("J0dN6dLO"))
	g.Expect(L2Segments.Status).To(Equal("pending"))
	g.Expect(L2Segments.Name).To(Equal("name87"))
	g.Expect(L2Segments.LocationGroupID).To(Equal(int64(1)))
	g.Expect(L2Segments.Created.String()).To(Equal("2020-04-22 06:22:51 +0000 UTC"))
	g.Expect(L2Segments.Updated.String()).To(Equal("2020-04-22 06:22:51 +0000 UTC"))
}

func TestL2SegmentsGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/y1aKReQG").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/l2_segments/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	L2Segments, err := client.L2Segments.Get(ctx, "y1aKReQG")

	g.Expect(err).To(BeNil())
	g.Expect(L2Segments).ToNot(BeNil())

	g.Expect(L2Segments.ID).To(Equal("y1aKReQG"))
	g.Expect(L2Segments.Status).To(Equal("active"))
	g.Expect(L2Segments.Name).To(Equal("name84"))
	g.Expect(L2Segments.Type).To(Equal("public"))
	g.Expect(L2Segments.LocationGroupID).To(Equal(int64(15)))
	g.Expect(L2Segments.Created.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
	g.Expect(L2Segments.Updated.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
}

func TestL2SegmentsUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/y1aKReQG").
		WithRequestMethod("PUT").
		WithRequestBody(`{"name":"some"}`).
		WithResponseBodyStubFile("fixtures/l2_segments/update_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newName := "some"

	L2Segments, err := client.L2Segments.Update(ctx, "y1aKReQG", L2SegmentUpdateInput{Name: &newName})

	g.Expect(err).To(BeNil())
	g.Expect(L2Segments).ToNot(BeNil())

	g.Expect(L2Segments.ID).To(Equal("y1aKReQG"))
	g.Expect(L2Segments.Status).To(Equal("pending"))
	g.Expect(L2Segments.Name).To(Equal("some"))
	g.Expect(L2Segments.Type).To(Equal("public"))
	g.Expect(L2Segments.LocationGroupID).To(Equal(int64(15)))
	g.Expect(L2Segments.Created.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
	g.Expect(L2Segments.Updated.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
}

func TestL2SegmentsDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/BDbDxbl2").
		WithRequestMethod("DELETE").
		WithResponseCode(204).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	err := client.L2Segments.Delete(ctx, "BDbDxbl2")

	g.Expect(err).To(BeNil())
}

func TestL2SegmentsChangeNetworksDeleteOnly(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/y1aKReQG/networks").
		WithRequestMethod("PUT").
		WithRequestBody(`{"delete":["a","b"]}`).
		WithResponseBodyStubFile("fixtures/l2_segments/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	input := L2SegmentChangeNetworksInput{
		Delete: []string{"a", "b"},
	}

	L2Segments, err := client.L2Segments.ChangeNetworks(ctx, "y1aKReQG", input)

	g.Expect(err).To(BeNil())
	g.Expect(L2Segments).ToNot(BeNil())

	g.Expect(L2Segments.ID).To(Equal("y1aKReQG"))
	g.Expect(L2Segments.Status).To(Equal("pending"))
	g.Expect(L2Segments.Name).To(Equal("some"))
	g.Expect(L2Segments.Type).To(Equal("public"))
	g.Expect(L2Segments.LocationGroupID).To(Equal(int64(15)))
	g.Expect(L2Segments.Created.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
	g.Expect(L2Segments.Updated.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
}

func TestL2SegmentsChangeNetworksCreateOnly(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/y1aKReQG/networks").
		WithRequestMethod("PUT").
		WithRequestBody(`{"create":[{"mask":32,"distribution_method":"route"}]}`).
		WithResponseBodyStubFile("fixtures/l2_segments/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	input := L2SegmentChangeNetworksInput{
		Create: []L2SegmentCreateNetworksInput{
			{Mask: 32, DistributionMethod: "route"},
		},
	}

	L2Segments, err := client.L2Segments.ChangeNetworks(ctx, "y1aKReQG", input)

	g.Expect(err).To(BeNil())
	g.Expect(L2Segments).ToNot(BeNil())

	g.Expect(L2Segments.ID).To(Equal("y1aKReQG"))
	g.Expect(L2Segments.Status).To(Equal("pending"))
	g.Expect(L2Segments.Name).To(Equal("some"))
	g.Expect(L2Segments.Type).To(Equal("public"))
	g.Expect(L2Segments.LocationGroupID).To(Equal(int64(15)))
	g.Expect(L2Segments.Created.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
	g.Expect(L2Segments.Updated.String()).To(Equal("2020-04-22 06:22:50 +0000 UTC"))
}

func TestL2SegmentsLocationGroupsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.L2Segments.LocationGroups()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2SegmentsMembersCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/a/members").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.L2Segments.Members("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2SegmentsNetworksCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/a/networks").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.L2Segments.Networks("a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}
