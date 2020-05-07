package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestL2LocationGroupsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestL2LocationGroupsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=3&per_page=2>; rel="next",<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="prev",<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="first",<https://dummy.api.com/l2_segments/location_groups?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestL2LocationGroupsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	var list []L2LocationGroup
	var err error

	list, err = collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))

	list, err = collection.NextPage(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []L2LocationGroup
	var err error

	list, err = collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))

	list, err = collection.PreviousPage(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []L2LocationGroup
	var err error

	list, err = collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasFirstPage()).To(Equal(true))

	list, err = collection.FirstPage(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	var list []L2LocationGroup
	var err error

	list, err = collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasLastPage()).To(Equal(true))

	list, err = collection.LastPage(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestL2LocationGroupsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/l2_segments/location_groups?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/l2_segments/location_groups").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 5}, {"id": 6}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewL2LocationGroupsCollection(client)

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
