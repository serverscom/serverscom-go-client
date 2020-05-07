package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestHostDriveSlotsEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostDriveSlotsCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostDriveSlotsCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostDriveSlotsCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostDriveSlotsCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestHostDriveSlotsCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestHostDriveSlotsCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=3&per_page=2>; rel="next",<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="prev",<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="first",<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestHostDriveSlotsCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"position": 2}, {"position": 3}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	var list []HostDriveSlot
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

func TestHostDriveSlotsCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"position": 2}, {"position": 3}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a").SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []HostDriveSlot
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

func TestHostDriveSlotsCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"position": 2}, {"position": 3}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a").SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []HostDriveSlot
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

func TestHostDriveSlotsCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"position": 2}, {"position": 3}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	var list []HostDriveSlot
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

func TestHostDriveSlotsCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"position": 0}, {"position": 1}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/hosts/dedicated_servers/a/drive_slots?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"position": 2}, {"position": 3}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/hosts/dedicated_servers/a/drive_slots").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"position": 4}, {"position": 5}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewHostDriveSlotsCollection(client, "dedicated_servers", "a")

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
