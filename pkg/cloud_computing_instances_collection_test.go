package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudComputingInstancesEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudComputingInstancesCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestCloudComputingInstancesCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=3&per_page=2>; rel="next",<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="prev",<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="first",<https://dummy.api.com/cloud_computing/instances?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestCloudComputingInstancesCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": "c"}, {"id": "d"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	var list []CloudComputingInstance
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

func TestCloudComputingInstancesCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"id": "c"}, {"id": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []CloudComputingInstance
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

func TestCloudComputingInstancesCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"id": "c"}, {"id": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []CloudComputingInstance
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

func TestCloudComputingInstancesCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"id": "c"}, {"id": "d"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	var list []CloudComputingInstance
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

func TestCloudComputingInstancesCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": "a"}, {"id": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/cloud_computing/instances?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"id": "c"}, {"id": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/cloud_computing/instances").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"id": "e"}, {"id": "f"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewCloudComputingInstancesCollection(client)

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
