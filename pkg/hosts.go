package serverscom

import (
	"context"
	"encoding/json"
)

const (
	hostListPath           = "/hosts"
	hostConnectionListPath = "/hosts/%s/%s/connections"
	hostNetworksListPath   = "/hosts/%s/%s/networks"
	hostDriveSlotListPath  = "/hosts/%s/%s/drive_slots"
	hostPTRsListPath       = "/hosts/%s/%s/ptr_records"

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

	kubernetesBaremetalNodePath = "/hosts/kubernetes_baremetal_nodes/%s"

	sbmServerCreatePath = "/hosts/sbm_servers"
	sbmServerPath       = "/hosts/sbm_servers/%s"
)

// HostsService is an interface for interfacing with Host, Dedicated Server endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Hosts
// https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server
type HostsService interface {
	// Primary collection
	Collection() Collection[Host]

	// Generic operations
	GetDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	GetKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error)
	GetSBMServer(ctx context.Context, id string) (*SBMServer, error)
	CreateDedicatedServers(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error)
	CreateSBMServers(ctx context.Context, input SBMServerCreateInput) ([]SBMServer, error)
	ReleaseSBMServer(ctx context.Context, id string) (*SBMServer, error)

	// Additional operations
	ScheduleReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	AbortReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerOnDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerOffDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerCycleDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	CreatePTRRecordForDedicatedServer(ctx context.Context, id string, input PTRRecordCreateInput) (*PTRRecord, error)
	DeletePTRRecordForDedicatedServer(ctx context.Context, hostID string, ptrRecordID string) error
	ReinstallOperatingSystemForDedicatedServer(ctx context.Context, id string, input OperatingSystemReinstallInput) (*DedicatedServer, error)

	// Additional collections
	DedicatedServerPowerFeeds(ctx context.Context, id string) ([]HostPowerFeed, error)
	DedicatedServerConnections(id string) Collection[HostConnection]
	DedicatedServerNetworks(id string) Collection[Network]
	DedicatedServerDriveSlots(id string) Collection[HostDriveSlot]
	DedicatedServerPTRRecords(id string) Collection[PTRRecord]
}

// HostsHandler handles operations around hosts
type HostsHandler struct {
	client *Client
}

// Collection builds a new Collection[Host] interface
func (h *HostsHandler) Collection() Collection[Host] {
	return NewCollection[Host](h.client, hostListPath)
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

// GetKubernetesBaremetalNode returns a kubernetes baremetal node
func (h *HostsHandler) GetKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error) {
	url := h.client.buildURL(kubernetesBaremetalNodePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	kubernetesBaremetalNode := new(KubernetesBaremetalNode)

	if err := json.Unmarshal(body, &kubernetesBaremetalNode); err != nil {
		return nil, err
	}

	return kubernetesBaremetalNode, nil
}

// CreateDedicatedServers creates a dedicated servers
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateANewDedicatedServer
func (h *HostsHandler) CreateDedicatedServers(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error) {
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

// ScheduleReleaseForDedicatedServer schedules release for for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ScheduleReleaseForAnExistingDedicatedServer
func (h *HostsHandler) ScheduleReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// AbortReleaseForDedicatedServer aborts scheduled release for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/AbortReleaseForAnExistingDedicatedServer
func (h *HostsHandler) AbortReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// PowerOnDedicatedServer sends power-on command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerOnCommandToAnExistingDedicatedServer
func (h *HostsHandler) PowerOnDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// PowerOffDedicatedServer sends power-off command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerOffCommandToAnExistingDedicatedServer
func (h *HostsHandler) PowerOffDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// PowerCycleDedicatedServer sends power-cycle command to the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SendPowerCycleCommandToAnExistingDedicatedServer
func (h *HostsHandler) PowerCycleDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error) {
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

// CreatePTRRecordForDedicatedServer creates ptr record for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreatePtrRecordForServerNetworks
func (h *HostsHandler) CreatePTRRecordForDedicatedServer(ctx context.Context, id string, input PTRRecordCreateInput) (*PTRRecord, error) {
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

// DeletePTRRecordForDedicatedServer deleted ptr record for the dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteAnExistingPtrRecord
func (h *HostsHandler) DeletePTRRecordForDedicatedServer(ctx context.Context, hostID string, ptrRecordID string) error {
	url := h.client.buildURL(dedicatedServerPTRRecordDeletePath, []interface{}{hostID, ptrRecordID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// ReinstallOperatingSystemForDedicatedServer performs operating system reinstallation
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/StartOperatingSystemReinstallProcess
func (h *HostsHandler) ReinstallOperatingSystemForDedicatedServer(ctx context.Context, id string, input OperatingSystemReinstallInput) (*DedicatedServer, error) {
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

// DedicatedServerConnections builds a new Collection[HostConnection] interface
func (h *HostsHandler) DedicatedServerConnections(id string) Collection[HostConnection] {
	path := h.client.buildPath(hostConnectionListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[HostConnection](h.client, path)
}

// DedicatedServerNetworks builds a new Collection[Network] interface
func (h *HostsHandler) DedicatedServerNetworks(id string) Collection[Network] {
	path := h.client.buildPath(hostNetworksListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[Network](h.client, path)
}

// DedicatedServerDriveSlots builds a new Collection[HostDriveSlot] interface
func (h *HostsHandler) DedicatedServerDriveSlots(id string) Collection[HostDriveSlot] {
	path := h.client.buildPath(hostDriveSlotListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[HostDriveSlot](h.client, path)
}

// DedicatedServerPTRRecords builds a new Collection[PTRRecord] interface
func (h *HostsHandler) DedicatedServerPTRRecords(id string) Collection[PTRRecord] {
	path := h.client.buildPath(hostPTRsListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[PTRRecord](h.client, path)
}

// GetSBMServer returns an sbm server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/GetAnSbmServer
func (h *HostsHandler) GetSBMServer(ctx context.Context, id string) (*SBMServer, error) {
	url := h.client.buildURL(sbmServerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	sbmServer := new(SBMServer)

	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return sbmServer, nil
}

// CreateSBMServers creates an SBM servers
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/CreateAnSbmServer
func (h *HostsHandler) CreateSBMServers(ctx context.Context, input SBMServerCreateInput) ([]SBMServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sbmServerCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var sbmServers []SBMServer

	if err := json.Unmarshal(body, &sbmServers); err != nil {
		return nil, err
	}

	return sbmServers, nil
}

// ReleaseSBMServer removes an SBM server from account.
// This action is irreversible and the removal process will be initiated immediately!!!
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/ReleaseAnSbmServer
func (h *HostsHandler) ReleaseSBMServer(ctx context.Context, id string) (*SBMServer, error) {
	url := h.client.buildURL(sbmServerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	sbmServer := new(SBMServer)

	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return sbmServer, nil
}
