package serverscom

import (
	"context"
	"encoding/json"
)

const (
	locationListPath = "/locations"
	locationPath     = "/locations/%d"

	serverModelOptionListPath = "/locations/%d/order_options/server_models"
	serverModelOptionPath     = "/locations/%d/order_options/server_models/%d"

	ramOptionListPath = "/locations/%d/order_options/server_models/%d/ram"

	operatingSystemOptionListPath = "/locations/%d/order_options/server_models/%d/operating_systems"
	operatingSystemOptionPath     = "/locations/%d/order_options/server_models/%d/operating_systems/%d"

	driveModelListPath = "/locations/%d/order_options/server_models/%d/drive_models"
	driveModelPath     = "/locations/%d/order_options/server_models/%d/drive_models/%d"

	uplinkOptionListPath = "/locations/%d/order_options/server_models/%d/uplink_models"
	uplinkOptionPath     = "/locations/%d/order_options/server_models/%d/uplink_models/%d"

	bandwidthOptionListPath = "/locations/%d/order_options/server_models/%d/uplink_models/%d/bandwidth"
	bandwidthOptionPath     = "/locations/%d/order_options/server_models/%d/uplink_models/%d/bandwidth/%d"

	sbmFlavorOptionListPath          = "/locations/%d/order_options/sbm_flavor_models"
	sbmFlavorOptionPath              = "/locations/%d/order_options/sbm_flavor_models/%d"
	sbmOperatingSystemOptionListPath = "/locations/%d/order_options/sbm_flavor_models/%d/operating_systems"
	sbmOperatingSystemOptionPath     = "/locations/%d/order_options/sbm_flavor_models/%d/operating_systems/%d"

	remoteBlockStorageFlavorListPath = "/locations/%d/order_options/remote_block_storage/flavors"
	remoteBlockStorageFlavorPath     = "/locations/%d/order_options/remote_block_storage/flavors/%d"
)

// LocationsService is an interface to interfacing with the Location and Order options endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Location
// https://developers.servers.com/api-documentation/v1/#tag/Server-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/SBM-Flavor-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Drive-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Ram-Option
// https://developers.servers.com/api-documentation/v1/#tag/Operating-System-Option
// https://developers.servers.com/api-documentation/v1/#tag/Uplink-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Bandwidth-Option
// https://developers.servers.com/api-documentation/v1/#tag/Remote-Block-Storage-Flavor
type LocationsService interface {
	// Primary collection
	Collection() Collection[Location]

	// Generic operations
	GetLocation(ctx context.Context, locationID int64) (*Location, error)
	ServerModelOptions(locationID int64) Collection[ServerModelOption]
	GetServerModelOption(ctx context.Context, locationID, serverModelID int64) (*ServerModelOptionDetail, error)
	RAMOptions(locationID, serverModelID int64) Collection[RAMOption]
	OperatingSystemOptions(locationID, serverModelID int64) Collection[OperatingSystemOption]
	GetOperatingSystemOption(ctx context.Context, locationID, serverModelID, operatingSystemID int64) (*OperatingSystemOption, error)
	DriveModelOptions(locationID, serverModelID int64) Collection[DriveModel]
	GetDriveModelOption(ctx context.Context, locationID, serverModelID, driveModelID int64) (*DriveModel, error)
	UplinkOptions(locationID, serverModelID int64) Collection[UplinkOption]
	GetUplinkOption(ctx context.Context, locationID, serverModelID, uplinkModelID int64) (*UplinkOption, error)
	BandwidthOptions(locationID, serverModelID, uplinkID int64) Collection[BandwidthOption]
	GetBandwidthOption(ctx context.Context, locationID, serverModelID, uplinkModelID, bandwidthID int64) (*BandwidthOption, error)
	SBMFlavorOptions(locationID int64) Collection[SBMFlavor]
	GetSBMFlavorOption(ctx context.Context, locationID, sbmFlavorModelID int64) (*SBMFlavor, error)
	SBMOperatingSystemOptions(locationID, sbmFlavorModelID int64) Collection[OperatingSystemOption]
	GetSBMOperatingSystemOption(ctx context.Context, locationID, sbmFlavorModelID, operatingSystemID int64) (*OperatingSystemOption, error)
	RemoteBlockStorageFlavors(locationID int64) Collection[RemoteBlockStorageFlavor]
	GetRemoteBlockStorageFlavor(ctx context.Context, locationID, flavorID int64) (*RemoteBlockStorageFlavor, error)
}

// LocationsHandler handles operations around cloud instances
type LocationsHandler struct {
	client *Client
}

// Collection builds a new LocationsCollection interface
func (h *LocationsHandler) Collection() Collection[Location] {
	return NewCollection[Location](h.client, locationListPath)
}

// GetLocation returns a location
func (h *LocationsHandler) GetLocation(ctx context.Context, id int64) (*Location, error) {
	url := h.client.buildURL(locationPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	location := new(Location)
	if err := json.Unmarshal(body, location); err != nil {
		return nil, err
	}

	return location, nil
}

// ServerModelOptions builds a new Collection[ServerModelOption interface
func (h *LocationsHandler) ServerModelOptions(LocationID int64) Collection[ServerModelOption] {
	path := h.client.buildPath(serverModelOptionListPath, []interface{}{LocationID}...)

	return NewCollection[ServerModelOption](h.client, path)
}

// GetServerModelOption returns a server model option
func (h *LocationsHandler) GetServerModelOption(ctx context.Context, locationID, serverModelID int64) (*ServerModelOptionDetail, error) {
	url := h.client.buildURL(serverModelOptionPath, locationID, serverModelID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	serverModelOption := new(ServerModelOptionDetail)
	if err := json.Unmarshal(body, serverModelOption); err != nil {
		return nil, err
	}

	return serverModelOption, nil
}

// RAMOptions builds a new Collection[RAMOption] interface
func (h *LocationsHandler) RAMOptions(LocationID, ServerModelID int64) Collection[RAMOption] {
	path := h.client.buildPath(ramOptionListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[RAMOption](h.client, path)
}

// OperatingSystemOptions builds a new Collection[OperatingSystemOption] interface
func (h *LocationsHandler) OperatingSystemOptions(LocationID, ServerModelID int64) Collection[OperatingSystemOption] {
	path := h.client.buildPath(operatingSystemOptionListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[OperatingSystemOption](h.client, path)
}

// GetOperatingSystemOption returns an operating system option
func (h *LocationsHandler) GetOperatingSystemOption(ctx context.Context, locationID, serverModelID, operatingSystemID int64) (*OperatingSystemOption, error) {
	url := h.client.buildURL(operatingSystemOptionPath, locationID, serverModelID, operatingSystemID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	operatingSystemOption := new(OperatingSystemOption)
	if err := json.Unmarshal(body, operatingSystemOption); err != nil {
		return nil, err
	}

	return operatingSystemOption, nil
}

// DriveModelOptions builds a new Collection[DriveModel]  interface
func (h *LocationsHandler) DriveModelOptions(LocationID, ServerModelID int64) Collection[DriveModel] {
	path := h.client.buildPath(driveModelListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[DriveModel](h.client, path)
}

// GetDriveModelOption returns a drive model
func (h *LocationsHandler) GetDriveModelOption(ctx context.Context, locationID, serverModelID, driveModelID int64) (*DriveModel, error) {
	url := h.client.buildURL(driveModelPath, locationID, serverModelID, driveModelID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	driveModel := new(DriveModel)
	if err := json.Unmarshal(body, driveModel); err != nil {
		return nil, err
	}

	return driveModel, nil
}

// UplinkOptions builds a new Collection[UplinkOption] interface
func (h *LocationsHandler) UplinkOptions(LocationID, ServerModelID int64) Collection[UplinkOption] {
	path := h.client.buildPath(uplinkOptionListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[UplinkOption](h.client, path)
}

// GetUplinkOption returns an uplink model
func (h *LocationsHandler) GetUplinkOption(ctx context.Context, locationID, serverModelID, uplinkModelID int64) (*UplinkOption, error) {
	url := h.client.buildURL(uplinkOptionPath, locationID, serverModelID, uplinkModelID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	uplinkOption := new(UplinkOption)
	if err := json.Unmarshal(body, uplinkOption); err != nil {
		return nil, err
	}

	return uplinkOption, nil
}

// BandwidthOptions builds a new Collection[BandwidthOption] interface
func (h *LocationsHandler) BandwidthOptions(LocationID, ServerModelID, uplinkID int64) Collection[BandwidthOption] {
	path := h.client.buildPath(bandwidthOptionListPath, []interface{}{LocationID, ServerModelID, uplinkID}...)

	return NewCollection[BandwidthOption](h.client, path)
}

// GetBandwidthOption returns a bandwidth option
func (h *LocationsHandler) GetBandwidthOption(ctx context.Context, locationID, serverModelID, uplinkModelID, bandwidthID int64) (*BandwidthOption, error) {
	url := h.client.buildURL(bandwidthOptionPath, locationID, serverModelID, uplinkModelID, bandwidthID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	bandwidthOption := new(BandwidthOption)
	if err := json.Unmarshal(body, bandwidthOption); err != nil {
		return nil, err
	}

	return bandwidthOption, nil
}

// SBMFlavorOptions builds a new Collection[SBMFlavor] interface
func (h *LocationsHandler) SBMFlavorOptions(LocationID int64) Collection[SBMFlavor] {
	path := h.client.buildPath(sbmFlavorOptionListPath, []interface{}{LocationID}...)

	return NewCollection[SBMFlavor](h.client, path)
}

// GetSBMFlavorOption returns an SBM flavor model
func (h *LocationsHandler) GetSBMFlavorOption(ctx context.Context, locationID, sbmFlavorModelID int64) (*SBMFlavor, error) {
	url := h.client.buildURL(sbmFlavorOptionPath, locationID, sbmFlavorModelID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	sbmFlavor := new(SBMFlavor)
	if err := json.Unmarshal(body, sbmFlavor); err != nil {
		return nil, err
	}

	return sbmFlavor, nil
}

// SBMOperatingSystemOptions builds a new Collection[OperatingSystemOption] interface
func (h *LocationsHandler) SBMOperatingSystemOptions(LocationID, SBMFlavorModelID int64) Collection[OperatingSystemOption] {
	path := h.client.buildPath(sbmOperatingSystemOptionListPath, []interface{}{LocationID, SBMFlavorModelID}...)

	return NewCollection[OperatingSystemOption](h.client, path)
}

// GetSBMOperatingSystemOption returns an SBM operating system option
func (h *LocationsHandler) GetSBMOperatingSystemOption(ctx context.Context, locationID, sbmFlavorModelID, operatingSystemID int64) (*OperatingSystemOption, error) {
	url := h.client.buildURL(sbmOperatingSystemOptionPath, locationID, sbmFlavorModelID, operatingSystemID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	operatingSystemOption := new(OperatingSystemOption)
	if err := json.Unmarshal(body, operatingSystemOption); err != nil {
		return nil, err
	}

	return operatingSystemOption, nil
}

// RemoteBlockStorageFlavors builds a new Collection[RemoteBlockStorageFlavor] interface
func (h *LocationsHandler) RemoteBlockStorageFlavors(locationID int64) Collection[RemoteBlockStorageFlavor] {
	path := h.client.buildPath(remoteBlockStorageFlavorListPath, []interface{}{locationID}...)
	return NewCollection[RemoteBlockStorageFlavor](h.client, path)
}

// GetRemoteBlockStorageFlavor returns an RBS flavor detail
func (h *LocationsHandler) GetRemoteBlockStorageFlavor(ctx context.Context, locationID, flavorID int64) (*RemoteBlockStorageFlavor, error) {
	url := h.client.buildURL(remoteBlockStorageFlavorPath, locationID, flavorID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	flavor := new(RemoteBlockStorageFlavor)
	if err := json.Unmarshal(body, flavor); err != nil {
		return nil, err
	}

	return flavor, nil
}
