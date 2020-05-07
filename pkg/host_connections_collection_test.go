package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestHostConnectionsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostConnectionsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostConnectionsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostConnectionsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostConnectionsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostConnectionsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestHostConnectionsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=3&per_page=2>; rel="next",<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="prev",<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="first",<https://dummy.api.com/hosts/dedicated_server/a/connections?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestHostConnectionsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"port": "NIC 3"}, {"port": "NIC 4"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	var list []HostConnection
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

func TestHostConnectionsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 3"}, {"port": "NIC 4"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a").SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []HostConnection
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

func TestHostConnectionsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 3"}, {"port": "NIC 4"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a").SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []HostConnection
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

func TestHostConnectionsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"port": "NIC 3"}, {"port": "NIC 4"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	var list []HostConnection
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

func TestHostConnectionsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 1"}, {"port": "NIC 2"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_server/a/connections?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"port": "NIC 3"}, {"port": "NIC 4"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_server/a/connections").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"port": "NIC 5"}, {"port": "NIC 6"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostConnectionsCollection(client, "dedicated_server", "a")

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
