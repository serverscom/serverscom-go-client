package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestRAMOptionsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRAMOptionsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRAMOptionsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRAMOptionsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRAMOptionsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRAMOptionsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestRAMOptionsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=3&per_page=2>; rel="next",<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="prev",<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="first",<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestRAMOptionsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"ram": 128}, {"ram": 256}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	var list []RAMOption
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

func TestRAMOptionsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 128}, {"ram": 256}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1)).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []RAMOption
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

func TestRAMOptionsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 128}, {"ram": 256}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1)).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []RAMOption
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

func TestRAMOptionsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"ram": 128}, {"ram": 256}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	var list []RAMOption
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

func TestRAMOptionsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 32}, {"ram": 64}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/ram?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"ram": 128}, {"ram": 256}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/ram").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"ram": 512}, {"ram": 1024}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewRAMOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
