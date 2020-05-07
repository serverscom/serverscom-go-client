package serverscom

import (
	"context"
	"encoding/json"
)

const (
	dedicatedServerTypePrefix = "dedicated_servers"

	dedicatedServerCreatePath          = "/hosts/dedicated_servers"
	dedicatedServerPath                = "/hosts/dedicated_servers/%s"
	dedicatedServerScheduleReleasePath = "/hosts/dedicated_servers/%s/schedule_release"
	dedicatedServerAbortReleasePath    = "/hosts/dedicated_servers/%s/abort_release"
	dedicatedServerPowerOnPath         = "/hosts/dedicated_servers/%s/power_on"
	dedicatedServerPowerOffPath        = "/hosts/dedicated_servers/%s/power_off"
	dedicatedServerPowerCyclePath      = "/hosts/dedicated_servers/%s/power_cycle"
	dedicatedServerPowerFeedsPath      = "/hosts/dedicated_servers/%s/power_feeds"
	dedicatedServerPTRRecordCreatePath = "/hosts/dedicated_servers/%s/ptr_records"
	dedicatedServerPTRRecordDeletePath = "/hosts/dedicated_servers/%s/ptr_records/%s"
	dedicatedServerReinstallPath       = "/hosts/dedicated_servers/%s/reinstall"
)

// HostsService is an interface for interfacing with Host, Dedicated Server endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Hosts
// https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server
type HostsService interface {
	Collection() HostsCollection

	DedicatedServerGet(ctx context.Context, id string) (*DedicatedServer, error)
	DedicatedServersCreate(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error)

	DedicatedServerScheduleRelease(ctx context.Context, id string) (*DedicatedServer, error)
	DedicatedServerAbortRelease(ctx context.Context, id string) (*DedicatedServer, error)

	DedicatedServerPowerOn(ctx context.Context, id string) (*DedicatedServer, error)
	DedicatedServerPowerOff(ctx context.Context, id string) (*DedicatedServer, error)
	DedicatedServerPowerCycle(ctx context.Context, id string) (*DedicatedServer, error)
	DedicatedServerPowerFeeds(ctx context.Context, id string) ([]HostPowerFeed, error)

	DedicatedServerConnections(ctx context.Context, id string) HostConnectionsCollection

	DedicatedServerNetworks(ctx context.Context, id string) HostNetworksCollection

	DedicatedServerPTRRecords(ctx context.Context, id string) HostPTRRecordsCollection
	DedicatedServerPTRRecordCreate(ctx context.Context, id string, input PTRRecordCreateInput) (*PTRRecord, error)
	DedicatedServerPTRRecordDelete(ctx context.Context, hostID string, ptrRecordID string) error

	DedicatedServerOperatingSystemReinstall(ctx context.Context, id string, input OperatingSystemReinstallInput) (*DedicatedServer, error)
}

// HostsHandler handles operations around hosts
type HostsHandler struct {
	client *Client
}

// Collection builds a new HostsCollection interface
func (h *HostsHandler) Collection() HostsCollection {
	return NewHostsCollection(h.client)
}

// DedicatedServerGet returns a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerGet(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServersCreate creates a dedicated servers
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateANewDedicatedServer
func (h *HostsHandler) DedicatedServersCreate(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dedicatedServerCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var dedicatedServers []DedicatedServer

	if err := json.Unmarshal(body, &dedicatedServers); err != nil {
		return nil, err
	}

	return dedicatedServers, nil
}

// DedicatedServerScheduleRelease schedules release for for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ScheduleReleaseForAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerScheduleRelease(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerScheduleReleasePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServerAbortRelease aborts scheduled release for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/AbortReleaseForAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerAbortRelease(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerAbortReleasePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServerPowerOn sends power-on command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerOnCommandToAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerPowerOn(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerPowerOnPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServerPowerOff sends power-off command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerOffCommandToAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerPowerOff(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerPowerOffPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServerPowerCycle sends power-cycle command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerCycleCommandToAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerPowerCycle(ctx context.Context, id string) (*DedicatedServer, error) {
	url := h.client.buildURL(dedicatedServerPowerCyclePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}

// DedicatedServerPowerFeeds returns list of dedicated server power feeds with status
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ListAllPowerFeedsForAnExistingDedicatedServer
func (h *HostsHandler) DedicatedServerPowerFeeds(ctx context.Context, id string) ([]HostPowerFeed, error) {
	url := h.client.buildURL(dedicatedServerPowerFeedsPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var powerFeeds []HostPowerFeed

	if err := json.Unmarshal(body, &powerFeeds); err != nil {
		return nil, err
	}

	return powerFeeds, nil
}

// DedicatedServerConnections builds a new HostConnectionsCollection interface
func (h *HostsHandler) DedicatedServerConnections(ctx context.Context, id string) HostConnectionsCollection {
	return NewHostConnectionsCollection(h.client, dedicatedServerTypePrefix, id)
}

// DedicatedServerNetworks builds a new HostNetworksCollection interface
func (h *HostsHandler) DedicatedServerNetworks(ctx context.Context, id string) HostNetworksCollection {
	return NewHostNetworksCollection(h.client, dedicatedServerTypePrefix, id)
}

// DedicatedServerPTRRecords builds a new HostPTRRecordsCollection interface
func (h *HostsHandler) DedicatedServerPTRRecords(ctx context.Context, id string) HostPTRRecordsCollection {
	return NewHostPTRRecordsCollection(h.client, dedicatedServerTypePrefix, id)
}

// DedicatedServerPTRRecordCreate creates ptr record for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreatePtrRecordForServerNetworks
func (h *HostsHandler) DedicatedServerPTRRecordCreate(ctx context.Context, id string, input PTRRecordCreateInput) (*PTRRecord, error) {
	url := h.client.buildURL(dedicatedServerPTRRecordCreatePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	ptrRecord := new(PTRRecord)

	if err := json.Unmarshal(body, &ptrRecord); err != nil {
		return nil, err
	}

	return ptrRecord, nil
}

// DedicatedServerPTRRecordDelete deleted ptr record for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteAnExistingPtrRecord
func (h *HostsHandler) DedicatedServerPTRRecordDelete(ctx context.Context, hostID string, ptrRecordID string) error {
	url := h.client.buildURL(dedicatedServerPTRRecordDeletePath, []interface{}{hostID, ptrRecordID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// DedicatedServerOperatingSystemReinstall performs operating system reinstallation
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/StartOperatingSystemReinstallProcess
func (h *HostsHandler) DedicatedServerOperatingSystemReinstall(ctx context.Context, id string, input OperatingSystemReinstallInput) (*DedicatedServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dedicatedServerReinstallPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	dedicatedServer := new(DedicatedServer)

	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return dedicatedServer, nil
}
