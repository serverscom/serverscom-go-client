package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDriveModelOptionsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDriveModelOptionsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDriveModelOptionsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDriveModelOptionsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDriveModelOptionsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestDriveModelOptionsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestDriveModelOptionsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=3&per_page=2>; rel="next",<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="prev",<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="first",<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestDriveModelOptionsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	var list []DriveModelOption
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

func TestDriveModelOptionsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1)).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []DriveModelOption
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

func TestDriveModelOptionsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1)).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []DriveModelOption
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

func TestDriveModelOptionsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	var list []DriveModelOption
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

func TestDriveModelOptionsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 1}, {"id": 2}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/locations/1/order_options/server_models/1/drive_models?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": 3}, {"id": 4}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/locations/1/order_options/server_models/1/drive_models").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"id": 5}, {"id": 6}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewDriveModelOptionsCollection(client, int64(1), int64(1))

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
