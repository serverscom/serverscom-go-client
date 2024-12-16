package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	volumeID = "Jrb2KdWL"
)

func TestCloudBlockStorageVolumesCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudBlockStorageVolumes.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudBlockStorageVolumesCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_volumes/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := CloudBlockStorageVolumeCreateInput{
		Name:   "test-volume-1",
		Labels: map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	cloudVolume, err := client.CloudBlockStorageVolumes.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("pending"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(cloudVolume.OpenstackUUID).To(BeNil())
	g.Expect(cloudVolume.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudVolume.Created).To(BeNil())
}

func TestCloudBlockStorageVolumesGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes/" + volumeID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_volumes/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudVolume, err := client.CloudBlockStorageVolumes.Get(ctx, volumeID)

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("backing-up"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(*cloudVolume.OpenstackUUID).To(Equal("7318f292-8f6a-48f4-8924-6cb087710abe"))
	g.Expect(cloudVolume.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudVolume.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageVolumesUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes/" + volumeID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/cloud_volumes/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	cloudVolume, err := client.CloudBlockStorageVolumes.Update(ctx, volumeID, CloudBlockStorageVolumeUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("creating"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(*cloudVolume.OpenstackUUID).To(Equal("7318f292-8f6a-48f4-8924-6cb087710abe"))
	g.Expect(cloudVolume.Labels).To(Equal(newLabels))
	g.Expect(cloudVolume.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageVolumesDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes/" + volumeID).
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/cloud_volumes/delete_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudVolume, err := client.CloudBlockStorageVolumes.Delete(ctx, volumeID)

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("deleting"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(*cloudVolume.OpenstackUUID).To(Equal("7318f292-8f6a-48f4-8924-6cb087710abe"))
	g.Expect(cloudVolume.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudVolume.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageVolumesDetach(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes/" + volumeID + "/detach").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_volumes/detach_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudVolume, err := client.CloudBlockStorageVolumes.Detach(ctx, volumeID, CloudBlockStorageVolumeDetachInput{InstanceID: "some-id"})

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("detaching"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(*cloudVolume.OpenstackUUID).To(Equal("7318f292-8f6a-48f4-8924-6cb087710abe"))
	g.Expect(cloudVolume.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudVolume.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageVolumesAttach(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/volumes/" + volumeID + "/attach").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_volumes/attach_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudVolume, err := client.CloudBlockStorageVolumes.Attach(ctx, volumeID, CloudBlockStorageVolumeAttachInput{InstanceID: "some-id"})

	g.Expect(err).To(BeNil())
	g.Expect(cloudVolume).ToNot(BeNil())

	g.Expect(cloudVolume.ID).To(Equal(volumeID))
	g.Expect(cloudVolume.Status).To(Equal("attaching"))
	g.Expect(cloudVolume.Name).To(Equal("test-volume-1"))
	g.Expect(cloudVolume.Size).To(Equal(101))
	g.Expect(cloudVolume.RegionID).To(Equal(int64(1)))
	g.Expect(cloudVolume.Bootable).To(BeFalse())
	g.Expect(cloudVolume.Attachments).To(HaveLen(0))
	g.Expect(*cloudVolume.Description).To(Equal("A fine volume"))
	g.Expect(*cloudVolume.OpenstackUUID).To(Equal("7318f292-8f6a-48f4-8924-6cb087710abe"))
	g.Expect(cloudVolume.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudVolume.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}
