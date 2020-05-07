package serverscom

import (
	"context"
	"encoding/json"
)

const (
	dedicatedServerCreatePath          = "/hosts/dedicated_servers"
	dedicatedServerPath                = "/hosts/dedicated_servers/%s"
	dedicatedServerScheduleReleasePath = "/hosts/dedicated_servers/%s/schedule_release"
	dedicatedServerAbortReleasePath    = "/hosts/dedicated_servers/%s/abort_release"
	dedicatedServerPowerOnPath         = "/hosts/dedicated_servers/%s/power_on"
	dedicatedServerPowerOffPath        = "/hosts/dedicated_servers/%s/power_off"
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

// DedicatedServerPowerOn sends power on command to the dedicated server
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

// DedicatedServerPowerOff sends power on command to the dedicated server
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
