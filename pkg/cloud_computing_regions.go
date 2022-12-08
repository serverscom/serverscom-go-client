package serverscom

const (
	cloudComputingRegionListPath = "/cloud_computing/regions"
	cloudComputingImageListPath  = "/cloud_computing/regions/%d/images"
	cloudComputingFlavorListPath = "/cloud_computing/regions/%d/flavors"
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
	path := h.client.buildPath(cloudComputingImageListPath, []interface{}{regionID}...)

	return NewCollection[CloudComputingImage](h.client, path)
}

// Flavors builds a new Collection[CloudComputingFlavor] interface
func (h *CloudComputingRegionsHandler) Flavors(regionID int64) Collection[CloudComputingFlavor] {
	path := h.client.buildPath(cloudComputingFlavorListPath, []interface{}{regionID}...)

	return NewCollection[CloudComputingFlavor](h.client, path)
}
