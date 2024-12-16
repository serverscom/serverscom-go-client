package serverscom

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	backupVolumeID = "X7ax9byv"
)

func TestCloudBlockStorageBackupsCollection(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups").
		WithRequestMethod("GET").
		WithResponseBodyStubInline(`[]`).
		WithResponseCode(200).
		Build()

	defer ts.Close()

	collection := client.CloudBlockStorageBackups.Collection()

	ctx := context.TODO()

	list, err := collection.List(ctx)

	g.Expect(err).To(BeNil())
	g.Expect(list).To(BeEmpty())
	g.Expect(collection.HasNextPage()).To(Equal(false))
	g.Expect(collection.HasPreviousPage()).To(Equal(false))
	g.Expect(collection.HasFirstPage()).To(Equal(false))
	g.Expect(collection.HasLastPage()).To(Equal(false))
}

func TestCloudBlockStorageBackupsCreate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_backups/create_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	input := CloudBlockStorageBackupCreateInput{
		VolumeID:    backupVolumeID,
		Name:        "test-backup-1",
		Incremental: true,
		Force:       false,
		Labels:      map[string]string{"env": "test"},
	}

	ctx := context.TODO()

	cloudBackup, err := client.CloudBlockStorageBackups.Create(ctx, input)

	g.Expect(err).To(BeNil())
	g.Expect(cloudBackup).ToNot(BeNil())

	g.Expect(cloudBackup.ID).To(Equal(backupVolumeID))
	g.Expect(cloudBackup.Status).To(Equal("pending"))
	g.Expect(cloudBackup.Name).To(Equal("test-backup-1"))
	g.Expect(cloudBackup.OpenstackVolumeUUID).To(Equal("1ac953dc-2414-4831-8124-c2dcbd405926"))
	g.Expect(cloudBackup.OpenstackUUID).To(BeNil())
	g.Expect(cloudBackup.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudBackup.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageBackupsGet(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups/" + backupVolumeID).
		WithRequestMethod("GET").
		WithResponseBodyStubFile("fixtures/cloud_backups/get_response.json").
		WithResponseCode(201).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudBackup, err := client.CloudBlockStorageBackups.Get(ctx, backupVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(cloudBackup).ToNot(BeNil())

	g.Expect(cloudBackup.ID).To(Equal(backupVolumeID))
	g.Expect(cloudBackup.Status).To(Equal("available"))
	g.Expect(cloudBackup.Name).To(Equal("test-backup-1"))
	g.Expect(cloudBackup.Size).To(Equal(1))
	g.Expect(cloudBackup.RegionID).To(Equal(1))
	g.Expect(*cloudBackup.OpenstackUUID).To(Equal("d1473317-bb2c-4d68-992c-0c117a68c01a"))
	g.Expect(cloudBackup.OpenstackVolumeUUID).To(Equal("1ac953dc-2414-4831-8124-c2dcbd405926"))
	g.Expect(cloudBackup.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudBackup.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageBackupsUpdate(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups/" + backupVolumeID).
		WithRequestMethod("PUT").
		WithResponseBodyStubFile("fixtures/cloud_backups/update_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newLabels := map[string]string{"env": "new-test"}

	cloudBackup, err := client.CloudBlockStorageBackups.Update(ctx, backupVolumeID, CloudBlockStorageBackupUpdateInput{Labels: newLabels})

	g.Expect(err).To(BeNil())
	g.Expect(cloudBackup).ToNot(BeNil())

	g.Expect(cloudBackup.ID).To(Equal(backupVolumeID))
	g.Expect(cloudBackup.Status).To(Equal("available"))
	g.Expect(cloudBackup.Name).To(Equal("test-backup-1"))
	g.Expect(cloudBackup.Size).To(Equal(1))
	g.Expect(cloudBackup.RegionID).To(Equal(1))
	g.Expect(*cloudBackup.OpenstackUUID).To(Equal("d1473317-bb2c-4d68-992c-0c117a68c01a"))
	g.Expect(cloudBackup.OpenstackVolumeUUID).To(Equal("1ac953dc-2414-4831-8124-c2dcbd405926"))
	g.Expect(cloudBackup.Labels).To(Equal(newLabels))
	g.Expect(cloudBackup.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageBackupsDelete(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups/" + backupVolumeID).
		WithRequestMethod("DELETE").
		WithResponseBodyStubFile("fixtures/cloud_backups/delete_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	cloudBackup, err := client.CloudBlockStorageBackups.Delete(ctx, backupVolumeID)

	g.Expect(err).To(BeNil())
	g.Expect(cloudBackup).ToNot(BeNil())

	g.Expect(cloudBackup.ID).To(Equal(backupVolumeID))
	g.Expect(cloudBackup.Status).To(Equal("deleting"))
	g.Expect(cloudBackup.Name).To(Equal("test-backup-1"))
	g.Expect(cloudBackup.Size).To(Equal(1))
	g.Expect(cloudBackup.RegionID).To(Equal(1))
	g.Expect(*cloudBackup.OpenstackUUID).To(Equal("d1473317-bb2c-4d68-992c-0c117a68c01a"))
	g.Expect(cloudBackup.OpenstackVolumeUUID).To(Equal("1ac953dc-2414-4831-8124-c2dcbd405926"))
	g.Expect(cloudBackup.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudBackup.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}

func TestCloudBlockStorageBackupsRestore(t *testing.T) {
	g := NewGomegaWithT(t)

	ts, client := newFakeServer().
		WithRequestPath("/cloud_block_storage/backups/" + backupVolumeID + "/restore").
		WithRequestMethod("POST").
		WithResponseBodyStubFile("fixtures/cloud_backups/restore_response.json").
		WithResponseCode(202).
		Build()

	defer ts.Close()

	ctx := context.TODO()

	newVolumeID := backupVolumeID

	cloudBackup, err := client.CloudBlockStorageBackups.Restore(ctx, backupVolumeID, CloudBlockStorageBackupRestoreInput{VolumeID: newVolumeID})

	g.Expect(err).To(BeNil())
	g.Expect(cloudBackup).ToNot(BeNil())

	g.Expect(cloudBackup.ID).To(Equal(backupVolumeID))
	g.Expect(cloudBackup.Status).To(Equal("restoring"))
	g.Expect(cloudBackup.Name).To(Equal("test-backup-1"))
	g.Expect(cloudBackup.Size).To(Equal(1))
	g.Expect(cloudBackup.RegionID).To(Equal(1))
	g.Expect(*cloudBackup.OpenstackUUID).To(Equal("d1473317-bb2c-4d68-992c-0c117a68c01a"))
	g.Expect(cloudBackup.OpenstackVolumeUUID).To(Equal("1ac953dc-2414-4831-8124-c2dcbd405926"))
	g.Expect(cloudBackup.Labels).To(Equal(map[string]string{"env": "test"}))
	g.Expect(cloudBackup.Created.String()).To(Equal("2024-11-11 09:57:28 +0000 UTC"))
}
