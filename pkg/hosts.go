package serverscom

import (
	"context"
	"encoding/json"
)

const (
	dedicatedServerCreatePath          = "/hosts/dedicated_servers"
	dedicatedServerPath                = "/hosts/dedicated_servers/%s"
	dedicatedServerScheduleReleasePath = "/hosts/dedicated_servers/%s/schedule_release"
)

// HostsService is an interface for interfacing with Host, Dedicated Server endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Hosts
// https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server
type HostsService interface {
	Collection() HostsCollection

	GetDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	CreateDedicatedServer(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error)
	ScheduleReleaseDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
}

// HostsHandler handles operations around hosts
type HostsHandler struct {
	client *Client
}

// Collection builds a new HostsCollection interface
func (h *HostsHandler) Collection() HostsCollection {
	return NewHostsCollection(h.client)
}

// GetDedicatedServer returns a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RetrieveAnExistingDedicatedServer
func (h *HostsHandler) GetDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// CreateDedicatedServer creates a dedicated servers
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateANewDedicatedServer
func (h *HostsHandler) CreateDedicatedServer(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error) {
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

// ScheduleReleaseDedicatedServer schedules dedicated server release
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ScheduleReleaseForAnExistingDedicatedServer
func (h *HostsHandler) ScheduleReleaseDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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
