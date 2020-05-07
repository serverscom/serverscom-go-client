package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSSHKeysEmptyCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSHKeysCollectionList(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSHKeysCollectionHasNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSHKeysCollectionHasPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSHKeysCollectionHasFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestSSHKeysCollectionHasLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestSSHKeysCollectionHasRelations(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=3&per_page=2>; rel="next",<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="prev",<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="first",<https://dummy.api.com/ssh_keys?page=3&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(2))
	g.Expect(collection.HasNextPage()).To(Equal(true))
	g.Expect(collection.HasPreviousPage()).To(Equal(true))
	g.Expect(collection.HasFirstPage()).To(Equal(true))
	g.Expect(collection.HasLastPage()).To(Equal(true))
}

func TestSSHKeysCollectionNext(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"fingerprint": "c"}, {"fingerprint": "d"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	var list []SSHKey
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

func TestSSHKeysCollectionPrevious(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="prev"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "c"}, {"fingerprint": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []SSHKey
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

func TestSSHKeysCollectionFirst(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=1&per_page=2>; rel="first"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "c"}, {"fingerprint": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=1&per_page=2`).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client).SetPage(2).SetPerPage(2)

	ctx := context.TODO()

	var list []SSHKey
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

func TestSSHKeysCollectionLast(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=2&per_page=2>; rel="last"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseBodyStubInline(`[{"fingerprint": "c"}, {"fingerprint": "d"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	var list []SSHKey
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

func TestSSHKeysCollectionCollect(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=2&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "a"}, {"fingerprint": "b"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=2&per_page=2`).
		WithResponseHeaders(map[string]string{
			"Link": `<https://dummy.api.com/ssh_keys?page=3&per_page=2>; rel="next"`,
		}).
		WithResponseBodyStubInline(`[{"fingerprint": "c"}, {"fingerprint": "d"}]`).
		WithResponseCode(200).
		Next().
		WithRequestPath("/ssh_keys").
		WithRequestMethod("GET").
		WithRequestParams(`page=3&per_page=2`).
		WithResponseBodyStubInline(`[{"fingerprint": "e"}, {"fingerprint": "f"}]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := NewSSHKeysCollection(client)

	ctx := context.TODO()

	list, err := collection.Collect(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(len(list)).To(Equal(6))
	g.Expect(collection.HasNextPage()).To(Equal(false))
}
