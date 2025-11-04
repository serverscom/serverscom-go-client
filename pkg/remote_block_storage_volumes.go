package serverscom

import (
	"context"
	"encoding/json"
)

const (
	remoteBlockStorageVolumePath       = "/remote_block_storage/volumes"
	remoteBlockStorageVolumePathWithID = remoteBlockStorageVolumePath + "/%s"
	actionGetCredentials               = "/credentials"
	actionResetCredentials             = "/credentials/reset"
)

// RemoteBlockStorageVolumesService is an interface for interfacing with Remote Block Storage Volume endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume
type RemoteBlockStorageVolumesService interface {
	// Primary collection
	Collection() Collection[RemoteBlockStorageVolume]

	// Generic operations
	Get(ctx context.Context, id string) (*RemoteBlockStorageVolume, error)
	Create(ctx context.Context, input RemoteBlockStorageVolumeCreateInput) (*RemoteBlockStorageVolume, error)
	Update(ctx context.Context, id string, input RemoteBlockStorageVolumeUpdateInput) (*RemoteBlockStorageVolume, error)
	Delete(ctx context.Context, id string) (*RemoteBlockStorageVolume, error)
	GetCredentials(ctx context.Context, id string) (*RemoteBlockStorageVolumeCredentials, error)
	ResetCredentials(ctx context.Context, id string) (*RemoteBlockStorageVolume, error)
}

// RemoteBlockStorageVolumesHandler handles operations around remote block storage volumes
type RemoteBlockStorageVolumesHandler struct {
	client *Client
}

// Collection builds a new Collection[RemoteBlockStorageVolume] interface
func (h *RemoteBlockStorageVolumesHandler) Collection() Collection[RemoteBlockStorageVolume] {
	return NewCollection[RemoteBlockStorageVolume](h.client, remoteBlockStorageVolumePath)
}

// Get a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/GetAnRbsVolumee
func (h *RemoteBlockStorageVolumesHandler) Get(ctx context.Context, id string) (*RemoteBlockStorageVolume, error) {
	url := h.client.buildURL(remoteBlockStorageVolumePathWithID, id)
	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var volume RemoteBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// Create a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/CreateAnRbsVolume
func (h *RemoteBlockStorageVolumesHandler) Create(ctx context.Context, input RemoteBlockStorageVolumeCreateInput) (*RemoteBlockStorageVolume, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	url := h.client.buildURL(remoteBlockStorageVolumePath)
	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}
	var volume RemoteBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// Update a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/UpdateVolume
func (h *RemoteBlockStorageVolumesHandler) Update(ctx context.Context, id string, input RemoteBlockStorageVolumeUpdateInput) (*RemoteBlockStorageVolume, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	url := h.client.buildURL(remoteBlockStorageVolumePathWithID, id)
	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)
	if err != nil {
		return nil, err
	}
	var volume RemoteBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// Delete a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/DeleteAnRbsVolume
func (h *RemoteBlockStorageVolumesHandler) Delete(ctx context.Context, id string) (*RemoteBlockStorageVolume, error) {
	url := h.client.buildURL(remoteBlockStorageVolumePathWithID, id)
	body, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	var volume RemoteBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}
	return &volume, nil
}

// Get credentials for a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/GetCredentialsForAnRbsVolume
func (h *RemoteBlockStorageVolumesHandler) GetCredentials(ctx context.Context, id string) (*RemoteBlockStorageVolumeCredentials, error) {
	url := h.client.buildURL(remoteBlockStorageVolumePathWithID+actionGetCredentials, id)
	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var creds RemoteBlockStorageVolumeCredentials
	if err := json.Unmarshal(body, &creds); err != nil {
		return nil, err
	}
	return &creds, nil
}

// Reset credentials for a remote block storage volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Volume/operation/ResetCredentialsForAnRbsVolume
func (h *RemoteBlockStorageVolumesHandler) ResetCredentials(ctx context.Context, id string) (*RemoteBlockStorageVolume, error) {
	url := h.client.buildURL(remoteBlockStorageVolumePathWithID+actionResetCredentials, id)
	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}
	var volume RemoteBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}
	return &volume, nil
}
