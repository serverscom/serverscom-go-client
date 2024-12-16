package serverscom

import (
	"context"
	"encoding/json"
)

const (
	cloudBlockStorageBackupPath       = "/cloud_block_storage/backups"
	cloudBlockStorageBackupPathWithID = cloudBlockStorageBackupPath + "/%s"
	actionRestore                     = "/restore"
)

// CloudBlockStorageBackupsService is an interface for interfacing with Cloud Backup endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup
type CloudBlockStorageBackupsService interface {
	// Primary collection
	Collection() Collection[CloudBlockStorageBackup]

	// Generic operations
	Get(ctx context.Context, id string) (*CloudBlockStorageBackup, error)
	Create(ctx context.Context, input CloudBlockStorageBackupCreateInput) (*CloudBlockStorageBackup, error)
	Update(ctx context.Context, id string, input CloudBlockStorageBackupUpdateInput) (*CloudBlockStorageBackup, error)
	Delete(ctx context.Context, id string) (*CloudBlockStorageBackup, error)
	Restore(ctx context.Context, id string, input CloudBlockStorageBackupRestoreInput) (*CloudBlockStorageBackup, error)
}

// CloudBlockStorageBackupsHandler handles operations around cloud backups
type CloudBlockStorageBackupsHandler struct {
	client *Client
}

// Collection builds a new Collection[CloudBlockStorageBackup] interface
func (h *CloudBlockStorageBackupsHandler) Collection() Collection[CloudBlockStorageBackup] {
	return NewCollection[CloudBlockStorageBackup](h.client, cloudBlockStorageBackupPath)
}

// Get a volume backup
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup/operation/GetAVolumeBackup
func (h *CloudBlockStorageBackupsHandler) Get(ctx context.Context, id string) (*CloudBlockStorageBackup, error) {
	url := h.client.buildURL(cloudBlockStorageBackupPathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var backup CloudBlockStorageBackup
	if err := json.Unmarshal(body, &backup); err != nil {
		return nil, err
	}

	return &backup, nil
}

// Create a backup from a cloud volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup/operation/CreateABackupFromACloudVolume
func (h *CloudBlockStorageBackupsHandler) Create(ctx context.Context, input CloudBlockStorageBackupCreateInput) (*CloudBlockStorageBackup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageBackupPath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var backup CloudBlockStorageBackup
	if err := json.Unmarshal(body, &backup); err != nil {
		return nil, err
	}

	return &backup, nil
}

// Update a volume backup
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup/operation/UpdateAVolumeBackup
func (h *CloudBlockStorageBackupsHandler) Update(ctx context.Context, id string, input CloudBlockStorageBackupUpdateInput) (*CloudBlockStorageBackup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageBackupPathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var backup CloudBlockStorageBackup
	if err := json.Unmarshal(body, &backup); err != nil {
		return nil, err
	}

	return &backup, nil
}

// Delete a volume backup
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup/operation/DeleteAVolumeBackup
func (h *CloudBlockStorageBackupsHandler) Delete(ctx context.Context, id string) (*CloudBlockStorageBackup, error) {
	url := h.client.buildURL(cloudBlockStorageBackupPathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	var backup CloudBlockStorageBackup
	if err := json.Unmarshal(body, &backup); err != nil {
		return nil, err
	}

	return &backup, nil
}

// Restore a volume backup
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Backup/operation/RestoreAVolumeBackup
func (h *CloudBlockStorageBackupsHandler) Restore(ctx context.Context, id string, input CloudBlockStorageBackupRestoreInput) (*CloudBlockStorageBackup, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageBackupPathWithID+actionRestore, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var backup CloudBlockStorageBackup
	if err := json.Unmarshal(body, &backup); err != nil {
		return nil, err
	}

	return &backup, nil
}
