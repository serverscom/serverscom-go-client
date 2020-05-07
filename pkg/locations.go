package serverscom

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
	Collection() LocationsCollection

	ServerModelOptionsCollection(LocationID int64) ServerModelOptionsCollection
	RAMOptionsCollection(LocationID, ServerModelID int64) RAMOptionsCollection
	OperatingSystemOptionsCollection(LocationID, ServerModelID int64) OperatingSystemOptionsCollection
	DriveModelOptionsCollection(LocationID, ServerModelID int64) DriveModelOptionsCollection
	UplinkOptionsCollection(LocationID, ServerModelID int64) UplinkOptionsCollection
	BandwidthOptionsCollection(LocationID, ServerModelID, uplinkID int64) BandwidthOptionsCollection
}

// LocationsHandler handles operations around cloud instances
type LocationsHandler struct {
	client *Client
}

// Collection builds a new LocationsCollection interface
func (resource *LocationsHandler) Collection() LocationsCollection {
	return NewLocationsCollection(resource.client)
}

// ServerModelOptionsCollection builds a new ServerModelOptionsCollection interface
func (resource *LocationsHandler) ServerModelOptionsCollection(LocationID int64) ServerModelOptionsCollection {
	return NewServerModelOptionsCollection(resource.client, LocationID)
}

// RAMOptionsCollection builds a new RAMOptionsCollection interface
func (resource *LocationsHandler) RAMOptionsCollection(LocationID, ServerModelID int64) RAMOptionsCollection {
	return NewRAMOptionsCollection(resource.client, LocationID, ServerModelID)
}

// OperatingSystemOptionsCollection builds a new OperatingSystemOptionsCollection interface
func (resource *LocationsHandler) OperatingSystemOptionsCollection(LocationID, ServerModelID int64) OperatingSystemOptionsCollection {
	return NewOperatingSystemOptionsCollection(resource.client, LocationID, ServerModelID)
}

// DriveModelOptionsCollection builds a new DriveModelOptionsCollection interface
func (resource *LocationsHandler) DriveModelOptionsCollection(LocationID, ServerModelID int64) DriveModelOptionsCollection {
	return NewDriveModelOptionsCollection(resource.client, LocationID, ServerModelID)
}

// UplinkOptionsCollection builds a new UplinkOptionsCollection interface
func (resource *LocationsHandler) UplinkOptionsCollection(LocationID, ServerModelID int64) UplinkOptionsCollection {
	return NewUplinkOptionsCollection(resource.client, LocationID, ServerModelID)
}

// BandwidthOptionsCollection builds a new BandwidthOptionsCollection interface
func (resource *LocationsHandler) BandwidthOptionsCollection(LocationID, ServerModelID, uplinkID int64) BandwidthOptionsCollection {
	return NewBandwidthOptionsCollection(resource.client, LocationID, ServerModelID, uplinkID)
}
