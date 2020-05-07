package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLocationsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestLocationsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestLocationsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=3&per_page=2>; rel="next",<https://dummy.api.com/locations?page=1&per_page=2>; rel="prev",<https://dummy.api.com/locations?page=1&per_page=2>; rel="first",<https://dummy.api.com/locations?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestLocationsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	var list []Location
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

func TestLocationsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []Location
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

func TestLocationsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []Location
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

func TestLocationsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	var list []Location
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

func TestLocationsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 5}, {"id": 6}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewLocationsCollection(client)

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
