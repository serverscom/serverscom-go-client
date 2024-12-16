package serverscom

import (
	"context"
	"encoding/json"
)

const (
	cloudBlockStorageVolumePath       = "/cloud_block_storage/volumes"
	cloudBlockStorageVolumePathWithID = cloudBlockStorageVolumePath + "/%s"
	actionAttach                      = "/attach"
	actionDetach                      = "/detach"
)

// CloudBlockStorageVolumesService is an interface for interfacing with Cloud Volume endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume
type CloudBlockStorageVolumesService interface {
	// Primary collection
	Collection() Collection[CloudBlockStorageVolume]

	// Generic operations
	Get(ctx context.Context, id string) (*CloudBlockStorageVolume, error)
	Create(ctx context.Context, input CloudBlockStorageVolumeCreateInput) (*CloudBlockStorageVolume, error)
	Update(ctx context.Context, id string, input CloudBlockStorageVolumeUpdateInput) (*CloudBlockStorageVolume, error)
	Delete(ctx context.Context, id string) (*CloudBlockStorageVolume, error)
	Attach(ctx context.Context, id string, input CloudBlockStorageVolumeAttachInput) (*CloudBlockStorageVolume, error)
	Detach(ctx context.Context, id string, input CloudBlockStorageVolumeDetachInput) (*CloudBlockStorageVolume, error)
}

// CloudBlockStorageVolumesHandler handles operations around cloud volumes
type CloudBlockStorageVolumesHandler struct {
	client *Client
}

// Collection builds a new Collection[CloudBlockStorageVolume] interface
func (h *CloudBlockStorageVolumesHandler) Collection() Collection[CloudBlockStorageVolume] {
	return NewCollection[CloudBlockStorageVolume](h.client, cloudBlockStorageVolumePath)
}

// Get a cloud volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/GetACloudVolume
func (h *CloudBlockStorageVolumesHandler) Get(ctx context.Context, id string) (*CloudBlockStorageVolume, error) {
	url := h.client.buildURL(cloudBlockStorageVolumePathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// Create a cloud volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/CreateACloudVolume
func (h *CloudBlockStorageVolumesHandler) Create(ctx context.Context, input CloudBlockStorageVolumeCreateInput) (*CloudBlockStorageVolume, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageVolumePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// Update a cloud volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/UpdateACloudVolume
func (h *CloudBlockStorageVolumesHandler) Update(ctx context.Context, id string, input CloudBlockStorageVolumeUpdateInput) (*CloudBlockStorageVolume, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageVolumePathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// Delete a cloud volume
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/DeleteACloudVolume
func (h *CloudBlockStorageVolumesHandler) Delete(ctx context.Context, id string) (*CloudBlockStorageVolume, error) {
	url := h.client.buildURL(cloudBlockStorageVolumePathWithID, id)

	body, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// Attach a cloud volume to a cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/AttachACloudVolumeToACloudInstance
func (h *CloudBlockStorageVolumesHandler) Attach(ctx context.Context, id string, input CloudBlockStorageVolumeAttachInput) (*CloudBlockStorageVolume, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageVolumePathWithID+actionAttach, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// Detach a cloud volume from a cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Volume/operation/DetachACloudVolumeFromACloudInstance
func (h *CloudBlockStorageVolumesHandler) Detach(ctx context.Context, id string, input CloudBlockStorageVolumeDetachInput) (*CloudBlockStorageVolume, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudBlockStorageVolumePathWithID+actionDetach, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var volume CloudBlockStorageVolume
	if err := json.Unmarshal(body, &volume); err != nil {
		return nil, err
	}

	return &volume, nil
}
