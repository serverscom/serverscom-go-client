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

	// ds networks
	dedicatedServerNetworkUsagePath          = "/hosts/dedicated_servers/%s/network_utilization"
	dedicatedServerNetworkPath               = "/hosts/dedicated_servers/%s/networks/%s"
	dedicatedServerAddPublicIPv4NetworkPath  = "/hosts/dedicated_servers/%s/networks/public_ipv4"
	dedicatedServerAddPrivateIPv4NetworkPath = "/hosts/dedicated_servers/%s/networks/private_ipv4"
	dedicatedServerActivatePublicIPv6Path    = "/hosts/dedicated_servers/%s/networks/public_ipv6"
	dedicatedServerDeleteNetworkPath         = "/hosts/dedicated_servers/%s/networks/%s"

	kubernetesBaremetalNodePath           = "/hosts/kubernetes_baremetal_nodes/%s"
	kubernetesBaremetalNodePowerOnPath    = "/hosts/kubernetes_baremetal_nodes/%s/power_on"
	kubernetesBaremetalNodePowerOffPath   = "/hosts/kubernetes_baremetal_nodes/%s/power_off"
	kubernetesBaremetalNodePowerCyclePath = "/hosts/kubernetes_baremetal_nodes/%s/power_cycle"

	sbmServerCreatePath     = "/hosts/sbm_servers"
	sbmServerPath           = "/hosts/sbm_servers/%s"
	sbmServerPowerOnPath    = "/hosts/sbm_servers/%s/power_on"
	sbmServerPowerOffPath   = "/hosts/sbm_servers/%s/power_off"
	sbmServerPowerCyclePath = "/hosts/sbm_servers/%s/power_cycle"
	sbmServerReinstallPath  = "/hosts/sbm_servers/%s/reinstall"
)

// HostsService is an interface for interfacing with Host, Dedicated Server endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Host
// https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server
// https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node
// https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server
type HostsService interface {
	// Primary collection
	Collection() Collection[Host]

	// Generic operations
	// dedicated
	CreateDedicatedServers(ctx context.Context, input DedicatedServerCreateInput) ([]DedicatedServer, error)
	GetDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	UpdateDedicatedServer(ctx context.Context, id string, input DedicatedServerUpdateInput) (*DedicatedServer, error)

	// kubernetes
	GetKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error)
	UpdateKubernetesBaremetalNode(ctx context.Context, id string, input KubernetesBaremetalNodeUpdateInput) (*KubernetesBaremetalNode, error)

	// sbm
	CreateSBMServers(ctx context.Context, input SBMServerCreateInput) ([]SBMServer, error)
	GetSBMServer(ctx context.Context, id string) (*SBMServer, error)
	UpdateSBMServer(ctx context.Context, id string, input SBMServerUpdateInput) (*SBMServer, error)
	ReleaseSBMServer(ctx context.Context, id string) (*SBMServer, error)

	// Additional operations
	// dedicated
	ScheduleReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	AbortReleaseForDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerOnDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerOffDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	PowerCycleDedicatedServer(ctx context.Context, id string) (*DedicatedServer, error)
	CreatePTRRecordForDedicatedServer(ctx context.Context, id string, input PTRRecordCreateInput) (*PTRRecord, error)
	DeletePTRRecordForDedicatedServer(ctx context.Context, serverID string, ptrRecordID string) error
	ReinstallOperatingSystemForDedicatedServer(ctx context.Context, id string, input OperatingSystemReinstallInput) (*DedicatedServer, error)

	// ds network methods
	GetDedicatedServerNetworkUsage(ctx context.Context, id string) (*NetworkUsage, error)
	GetDedicatedServerNetwork(ctx context.Context, serverID string, networkID string) (*Network, error)
	AddDedicatedServerPublicIPv4Network(ctx context.Context, id string, input NetworkInput) (*Network, error)
	AddDedicatedServerPrivateIPv4Network(ctx context.Context, id string, input NetworkInput) (*Network, error)
	ActivateDedicatedServerPubliIPv6Network(ctx context.Context, id string) (*Network, error)
	DeleteDedicatedServerNetwork(ctx context.Context, serverID string, networkID string) (*Network, error)

	// sbm
	PowerOnSBMServer(ctx context.Context, id string) (*SBMServer, error)
	PowerOffSBMServer(ctx context.Context, id string) (*SBMServer, error)
	PowerCycleSBMServer(ctx context.Context, id string) (*SBMServer, error)
	ReinstallOperatingSystemForSBMServer(ctx context.Context, id string, input SBMOperatingSystemReinstallInput) (*SBMServer, error)

	// kubernetes
	PowerOnKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error)
	PowerOffKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error)
	PowerCycleKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error)

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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node/operation/GetAKubernetesBareMetalNode
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/CreateADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ScheduleReleaseForADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/AbortReleaseForADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/PowerOnADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/PowerOffADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/PowercycleADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListPowerFeedsForADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListPtrRecordsForADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/DeleteAPtrRecordForADedicatedServer
func (h *HostsHandler) DeletePTRRecordForDedicatedServer(ctx context.Context, hostID string, ptrRecordID string) error {
	url := h.client.buildURL(dedicatedServerPTRRecordDeletePath, []interface{}{hostID, ptrRecordID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// ReinstallOperatingSystemForDedicatedServer performs operating system reinstallation
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ReinstallOsForADedicatedServer
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
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListConnectionsForADedicatedServer
func (h *HostsHandler) DedicatedServerConnections(id string) Collection[HostConnection] {
	path := h.client.buildPath(hostConnectionListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[HostConnection](h.client, path)
}

// DedicatedServerNetworks builds a new Collection[Network] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListNetworksForADedicatedServer
func (h *HostsHandler) DedicatedServerNetworks(id string) Collection[Network] {
	path := h.client.buildPath(hostNetworksListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[Network](h.client, path)
}

// DedicatedServerDriveSlots builds a new Collection[HostDriveSlot] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListDriveSlotsForADedicatedServer
func (h *HostsHandler) DedicatedServerDriveSlots(id string) Collection[HostDriveSlot] {
	path := h.client.buildPath(hostDriveSlotListPath, []interface{}{dedicatedServerTypePrefix, id}...)

	return NewCollection[HostDriveSlot](h.client, path)
}

// DedicatedServerPTRRecords builds a new Collection[PTRRecord] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ListPtrRecordsForADedicatedServer
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

// Update a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/UpdateADedicatedServer
func (h *HostsHandler) UpdateDedicatedServer(ctx context.Context, id string, input DedicatedServerUpdateInput) (*DedicatedServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dedicatedServerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var dedicatedServer DedicatedServer
	if err := json.Unmarshal(body, &dedicatedServer); err != nil {
		return nil, err
	}

	return &dedicatedServer, nil
}

// Update a Kubernetes bare metal node
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node/operation/UpdateAKubernetesBareMetalNode
func (h *HostsHandler) UpdateKubernetesBaremetalNode(ctx context.Context, id string, input KubernetesBaremetalNodeUpdateInput) (*KubernetesBaremetalNode, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(kubernetesBaremetalNodePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var node KubernetesBaremetalNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// UpdateSBMServer updates an SBM server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/UpdateAnSbmServer
func (h *HostsHandler) UpdateSBMServer(ctx context.Context, id string, input SBMServerUpdateInput) (*SBMServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sbmServerPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var sbmServer SBMServer
	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return &sbmServer, nil
}

// Send a power on command for an SBM server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/SendAPowerOnCommandForAnSbmServer
func (h *HostsHandler) PowerOnSBMServer(ctx context.Context, id string) (*SBMServer, error) {
	url := h.client.buildURL(sbmServerPowerOnPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var sbmServer SBMServer
	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return &sbmServer, nil
}

// Send a power off command for an SBM server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/SendAPowerOffCommandForAnSbmServer
func (h *HostsHandler) PowerOffSBMServer(ctx context.Context, id string) (*SBMServer, error) {
	url := h.client.buildURL(sbmServerPowerOffPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var sbmServer SBMServer
	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return &sbmServer, nil
}

// Send a power cycle command for an SBM server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/SendAPowerCycleCommandForAnSbmServer
func (h *HostsHandler) PowerCycleSBMServer(ctx context.Context, id string) (*SBMServer, error) {
	url := h.client.buildURL(sbmServerPowerCyclePath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var sbmServer SBMServer
	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return &sbmServer, nil
}

// Reinstall an OS for an SBM server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Scalable-Baremetal-Server/operation/ReinstallAnOsForAnSbmServer
func (h *HostsHandler) ReinstallOperatingSystemForSBMServer(ctx context.Context, id string, input SBMOperatingSystemReinstallInput) (*SBMServer, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sbmServerReinstallPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var sbmServer SBMServer
	if err := json.Unmarshal(body, &sbmServer); err != nil {
		return nil, err
	}

	return &sbmServer, nil
}

// Power on a Kubernetes bare metal node
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node/operation/PowerOnAKubernetesBareMetalNode
func (h *HostsHandler) PowerOnKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error) {
	url := h.client.buildURL(kubernetesBaremetalNodePowerOnPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var node KubernetesBaremetalNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// Power off a Kubernetes bare metal node
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node/operation/PowerOffAKubernetesBareMetalNode
func (h *HostsHandler) PowerOffKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error) {
	url := h.client.buildURL(kubernetesBaremetalNodePowerOffPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var node KubernetesBaremetalNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// Powercycle a Kubernetes bare metal node
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Kubernetes-Baremetal-Node/operation/PowercycleAKubernetesBareMetalNode
func (h *HostsHandler) PowerCycleKubernetesBaremetalNode(ctx context.Context, id string) (*KubernetesBaremetalNode, error) {
	url := h.client.buildURL(kubernetesBaremetalNodePowerCyclePath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var node KubernetesBaremetalNode
	if err := json.Unmarshal(body, &node); err != nil {
		return nil, err
	}

	return &node, nil
}

// Get network utilization for a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/GetNetworkUsageForADedicatedServer
func (h *HostsHandler) GetDedicatedServerNetworkUsage(ctx context.Context, id string) (*NetworkUsage, error) {
	url := h.client.buildURL(dedicatedServerNetworkUsagePath, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var networkUsage NetworkUsage
	if err := json.Unmarshal(body, &networkUsage); err != nil {
		return nil, err
	}

	return &networkUsage, nil
}

// Get network details for a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/GetANetworkForADedicatedServer
func (h *HostsHandler) GetDedicatedServerNetwork(ctx context.Context, serverID, networkID string) (*Network, error) {
	url := h.client.buildURL(dedicatedServerNetworkPath, serverID, networkID)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Add a public IPv4 network to a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/CreateAPublicIpv4NetworkForADedicatedServer
func (h *HostsHandler) AddDedicatedServerPublicIPv4Network(ctx context.Context, id string, input NetworkInput) (*Network, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dedicatedServerAddPublicIPv4NetworkPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Add a private IPv4 network to a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/CreateAPrivateIpv4NetworkForADedicatedServer
func (h *HostsHandler) AddDedicatedServerPrivateIPv4Network(ctx context.Context, id string, input NetworkInput) (*Network, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dedicatedServerAddPrivateIPv4NetworkPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Activate a public IPv6 network for a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/ActivateAPublicIpv6NetworkForADedicatedServer
func (h *HostsHandler) ActivateDedicatedServerPubliIPv6Network(ctx context.Context, id string) (*Network, error) {
	url := h.client.buildURL(dedicatedServerActivatePublicIPv6Path, id)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Delete a network from a dedicated server
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Dedicated-Server/operation/DeleteANetworkForADedicatedServer
func (h *HostsHandler) DeleteDedicatedServerNetwork(ctx context.Context, serverID, networkID string) (*Network, error) {
	url := h.client.buildURL(dedicatedServerDeleteNetworkPath, serverID, networkID)

	body, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(body, &network); err != nil {
		return nil, err
	}

	return &network, nil
}
