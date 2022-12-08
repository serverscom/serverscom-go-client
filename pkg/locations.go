package serverscom

const (
	locationListPath = "/locations"

	serverModelOptionListPath = "/locations/%d/order_options/server_models"

	ramOptionListPath = "/locations/%d/order_options/server_models/%d/ram"

	operatingSystemOptionListPath = "/locations/%d/order_options/server_models/%d/operating_systems"

	driveModelListPath = "/locations/%d/order_options/server_models/%d/drive_models"

	uplinkOptionListPath = "/locations/%d/order_options/server_models/%d/uplink_models"

	bandwidthOptionListPath = "/locations/%d/order_options/server_models/%d/uplink_models/%d/bandwidth"
)

// LocationsService is an interface to interfacing with the Location and Order options endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Location
// https://developers.servers.com/api-documentation/v1/#tag/Server-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Drive-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Ram-Option
// https://developers.servers.com/api-documentation/v1/#tag/Operating-System-Option
// https://developers.servers.com/api-documentation/v1/#tag/Uplink-Model-Option
// https://developers.servers.com/api-documentation/v1/#tag/Bandwidth-Option
type LocationsService interface {
	// Primary collection
	Collection() Collection[Location]

	// Generic operations
	ServerModelOptions(LocationID int64) Collection[ServerModelOption]
	RAMOptions(LocationID, ServerModelID int64) Collection[RAMOption]
	OperatingSystemOptions(LocationID, ServerModelID int64) Collection[OperatingSystemOption]
	DriveModelOptions(LocationID, ServerModelID int64) Collection[DriveModel]
	UplinkOptions(LocationID, ServerModelID int64) Collection[UplinkOption]
	BandwidthOptions(LocationID, ServerModelID, uplinkID int64) Collection[BandwidthOption]
}

// LocationsHandler handles operations around cloud instances
type LocationsHandler struct {
	client *Client
}

// Collection builds a new LocationsCollection interface
func (h *LocationsHandler) Collection() Collection[Location] {
	return NewCollection[Location](h.client, locationListPath)
}

// ServerModelOptions builds a new Collection[ServerModelOption interface
func (h *LocationsHandler) ServerModelOptions(LocationID int64) Collection[ServerModelOption] {
	path := h.client.buildPath(serverModelOptionListPath, []interface{}{LocationID}...)

	return NewCollection[ServerModelOption](h.client, path)
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

// DriveModelOptions builds a new Collection[DriveModel]  interface
func (h *LocationsHandler) DriveModelOptions(LocationID, ServerModelID int64) Collection[DriveModel] {
	path := h.client.buildPath(driveModelListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[DriveModel](h.client, path)
}

// UplinkOptions builds a new Collection[UplinkOption] interface
func (h *LocationsHandler) UplinkOptions(LocationID, ServerModelID int64) Collection[UplinkOption] {
	path := h.client.buildPath(uplinkOptionListPath, []interface{}{LocationID, ServerModelID}...)

	return NewCollection[UplinkOption](h.client, path)
}

// BandwidthOptions builds a new Collection[BandwidthOption] interface
func (h *LocationsHandler) BandwidthOptions(LocationID, ServerModelID, uplinkID int64) Collection[BandwidthOption] {
	path := h.client.buildPath(bandwidthOptionListPath, []interface{}{LocationID, ServerModelID, uplinkID}...)

	return NewCollection[BandwidthOption](h.client, path)
}
