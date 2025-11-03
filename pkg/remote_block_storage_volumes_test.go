package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	rbsVolumeID = "Rbs2KdWL"
)

func TestRemoteBlockStorageVolumesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.RemoteBlockStorageVolumes.Collection()

	ctx := context.TODO()
	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestRemoteBlockStorageVolumesCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/rbs_volumes/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := RemoteBlockStorageVolumeCreateInput{
		Name:   "rbs-volume-1",
		Labels: map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	volume, err := client.RemoteBlockStorageVolumes.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(volume).ToNot(BeNil())
	g.Expect(volume.ID).To(Equal(rbsVolumeID))
	g.Expect(volume.Status).To(Equal("pending"))
	g.Expect(volume.Name).To(Equal("rbs-volume-1"))
}

func TestRemoteBlockStorageVolumesGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes/" + rbsVolumeID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/rbs_volumes/get_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	volume, err := client.RemoteBlockStorageVolumes.Get(ctx, rbsVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(volume).ToNot(BeNil())
	g.Expect(volume.ID).To(Equal(rbsVolumeID))
	g.Expect(volume.Status).To(Equal("available"))
	g.Expect(volume.Name).To(Equal("rbs-volume-1"))
}

func TestRemoteBlockStorageVolumesUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes/" + rbsVolumeID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/rbs_volumes/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	volume, err := client.RemoteBlockStorageVolumes.Update(ctx, rbsVolumeID, RemoteBlockStorageVolumeUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(volume).ToNot(BeNil())
	g.Expect(volume.ID).To(Equal(rbsVolumeID))
	g.Expect(volume.Labels).To(Equal(newLabels))
}

func TestRemoteBlockStorageVolumesDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes/" + rbsVolumeID).
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/rbs_volumes/delete_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	volume, err := client.RemoteBlockStorageVolumes.Delete(ctx, rbsVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(volume).ToNot(BeNil())
	g.Expect(volume.ID).To(Equal(rbsVolumeID))
	g.Expect(volume.Status).To(Equal("deleting"))
}

func TestRemoteBlockStorageVolumesGetCredentials(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes/" + rbsVolumeID + "/credentials").
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/rbs_volumes/get_credentials_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	creds, err := client.RemoteBlockStorageVolumes.GetCredentials(ctx, rbsVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(creds).ToNot(BeNil())
	g.Expect(creds.Username).To(Equal("rbs-user"))
	g.Expect(creds.Password).To(Equal("passwd123"))
}

func TestRemoteBlockStorageVolumesResetCredentials(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/remote_block_storage/volumes/" + rbsVolumeID + "/credentials/reset").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/rbs_volumes/reset_credentials_response.json").
		WithResponseCode(200).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	volume, err := client.RemoteBlockStorageVolumes.ResetCredentials(ctx, rbsVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(volume).ToNot(BeNil())
	g.Expect(volume.ID).To(Equal(rbsVolumeID))
	g.Expect(volume.Status).To(Equal("pending"))
	g.Expect(volume.Name).To(Equal("rbs-volume-1"))
}
