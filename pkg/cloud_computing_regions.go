package serverscom

import (
	"context"
	"encoding/json"
)

const (
	cloudComputingRegionListPath     = "/cloud_computing/regions"
	cloudComputingImageListPath      = "/cloud_computing/regions/%d/images"
	cloudComputingFlavorListPath     = "/cloud_computing/regions/%d/flavors"
	cloudComputingCredentialsPath    = "/cloud_computing/regions/%d/credentials"
	cloudComputingSnapshotListPath   = "/cloud_computing/regions/%d/snapshots"
	cloudComputingSnapshotDeletePath = "/cloud_computing/regions/%d/snapshots/%s"
)

// CloudComputingRegionsService is an interface to interfacing with the cloud computing regions endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Cloud-Region
type CloudComputingRegionsService interface {
	// Primary collection
	Collection() Collection[CloudComputingRegion]

	// Additional collections
	Images(regionID int64) Collection[CloudComputingImage]
	Flavors(regionID int64) Collection[CloudComputingFlavor]
	Snapshots(regionID int64) Collection[CloudSnapshot]

	Credentials(ctx context.Context, regionID int64) (*CloudComputingRegionCredentials, error)
	CreateSnapshot(ctx context.Context, regionID int64, input CloudSnapshotCreateInput) (*CloudSnapshot, error)
	DeleteSnapshot(ctx context.Context, regionID int64, snapshotID string) error
}

// CloudComputingRegionsHandler handles operations around cloud computing regions
type CloudComputingRegionsHandler struct {
	client *Client
}

// Collection builds a new Collection[CloudComputingRegion] interface
func (h *CloudComputingRegionsHandler) Collection() Collection[CloudComputingRegion] {
	return NewCollection[CloudComputingRegion](h.client, cloudComputingRegionListPath)
}

// Images builds a new Collection[CloudComputingImage] interface
func (h *CloudComputingRegionsHandler) Images(regionID int64) Collection[CloudComputingImage] {
	path := h.client.buildPath(cloudComputingImageListPath, regionID)

	return NewCollection[CloudComputingImage](h.client, path)
}

// Flavors builds a new Collection[CloudComputingFlavor] interface
func (h *CloudComputingRegionsHandler) Flavors(regionID int64) Collection[CloudComputingFlavor] {
	path := h.client.buildPath(cloudComputingFlavorListPath, regionID)

	return NewCollection[CloudComputingFlavor](h.client, path)
}

// Snapshots builds a new Collection[CloudSnapshot] interface
func (h *CloudComputingRegionsHandler) Snapshots(regionID int64) Collection[CloudSnapshot] {
	path := h.client.buildPath(cloudComputingSnapshotListPath, regionID)

	return NewCollection[CloudSnapshot](h.client, path)
}

// Credentials returns cloud region OpenStack credentials
func (h *CloudComputingRegionsHandler) Credentials(ctx context.Context, regionID int64) (*CloudComputingRegionCredentials, error) {
	url := h.client.buildURL(cloudComputingCredentialsPath, regionID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	credentials := new(CloudComputingRegionCredentials)

	if err := json.Unmarshal(body, credentials); err != nil {
		return nil, err
	}

	return credentials, nil
}

// CreateSnapshot creates a snapshot for a cloud instance
func (h *CloudComputingRegionsHandler) CreateSnapshot(ctx context.Context, regionID int64, input CloudSnapshotCreateInput) (*CloudSnapshot, error) {
	url := h.client.buildURL(cloudComputingSnapshotListPath, regionID)

	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	snapshot := new(CloudSnapshot)

	if err := json.Unmarshal(body, snapshot); err != nil {
		return nil, err
	}

	return snapshot, nil
}

// DeleteSnapshot deletes a snapshot
func (h *CloudComputingRegionsHandler) DeleteSnapshot(ctx context.Context, regionID int64, snapshotID string) error {
	url := h.client.buildURL(cloudComputingSnapshotDeletePath, regionID, snapshotID)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
