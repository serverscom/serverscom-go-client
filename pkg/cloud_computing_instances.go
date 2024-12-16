package serverscom

import (
	"context"
	"encoding/json"
)

const (
	cloudInstanceListPath            = "/cloud_computing/instances"
	cloudInstanceCreatePath          = "/cloud_computing/instances"
	cloudInstancePath                = "/cloud_computing/instances/%s"
	cloudInstanceUpdatePath          = "/cloud_computing/instances/%s"
	cloudInstanceDeletePath          = "/cloud_computing/instances/%s"
	cloudInstanceReinstallPath       = "/cloud_computing/instances/%s/reinstall"
	cloudInstanceRescuePath          = "/cloud_computing/instances/%s/rescue"
	cloudInstanceUnrescuePath        = "/cloud_computing/instances/%s/unrescue"
	cloudInstanceUpgradePath         = "/cloud_computing/instances/%s/upgrade"
	cloudInstanceRevertUpgradePath   = "/cloud_computing/instances/%s/revert_upgrade"
	cloudInstanceApproveUpgradePath  = "/cloud_computing/instances/%s/approve_upgrade"
	cloudInstancePowerOnPath         = "/cloud_computing/instances/%s/switch_power_on"
	cloudInstancePowerOffPath        = "/cloud_computing/instances/%s/switch_power_off"
	cloudInstanceCreatePTRRecordPath = "/cloud_computing/instances/%s/ptr_records"
	cloudInstanceDeletePTRRecordPath = "/cloud_computing/instances/%s/ptr_records/%s"

	cloudInstancePTRsListPath = "/cloud_computing/instances/%s/ptr_records"
)

// CloudComputingInstancesService is an interface to interfacing with the Cloud Instance endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance
type CloudComputingInstancesService interface {
	// Primary collection
	Collection() Collection[CloudComputingInstance]

	// Generic operations
	Get(ctx context.Context, id string) (*CloudComputingInstance, error)
	Create(ctx context.Context, input CloudComputingInstanceCreateInput) (*CloudComputingInstance, error)
	Update(ctx context.Context, id string, input CloudComputingInstanceUpdateInput) (*CloudComputingInstance, error)
	Delete(ctx context.Context, id string) error

	// Additional operations
	Reinstall(ctx context.Context, id string, input CloudComputingInstanceReinstallInput) (*CloudComputingInstance, error)
	Rescue(ctx context.Context, id string) (*CloudComputingInstance, error)
	Unrescue(ctx context.Context, id string) (*CloudComputingInstance, error)
	Upgrade(ctx context.Context, id string, input CloudComputingInstanceUpgradeInput) (*CloudComputingInstance, error)
	RevertUpgrade(ctx context.Context, id string) (*CloudComputingInstance, error)
	ApproveUpgrade(ctx context.Context, id string) (*CloudComputingInstance, error)
	PowerOn(ctx context.Context, id string) (*CloudComputingInstance, error)
	PowerOff(ctx context.Context, id string) (*CloudComputingInstance, error)
	CreatePTRRecord(ctx context.Context, cloudInstanceID string, input PTRRecordCreateInput) (*PTRRecord, error)
	DeletePTRRecord(ctx context.Context, cloudInstanceID string, ptrRecordID string) error

	// Additional collections
	PTRRecords(id string) Collection[PTRRecord]
}

// CloudComputingInstancesHandler handles operations around cloud instances
type CloudComputingInstancesHandler struct {
	client *Client
}

// Collection builds a new Collection[CloudComputingInstance] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/ListCloudInstances
func (h *CloudComputingInstancesHandler) Collection() Collection[CloudComputingInstance] {
	return NewCollection[CloudComputingInstance](h.client, cloudInstanceListPath)
}

// Get cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/GetACloudInstance
func (h *CloudComputingInstancesHandler) Get(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstancePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	cloudInstance := new(CloudComputingInstance)

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Create cloud instace
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/CreateACloudInstance
func (h *CloudComputingInstancesHandler) Create(ctx context.Context, input CloudComputingInstanceCreateInput) (*CloudComputingInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudInstanceCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Update cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/UpdateACloudInstance
func (h *CloudComputingInstancesHandler) Update(ctx context.Context, id string, input CloudComputingInstanceUpdateInput) (*CloudComputingInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudInstanceUpdatePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Delete cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/DeleteACloudInstance
func (h *CloudComputingInstancesHandler) Delete(ctx context.Context, id string) error {
	url := h.client.buildURL(cloudInstanceDeletePath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// Reinstall cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/ReinstallACloudInstance
func (h *CloudComputingInstancesHandler) Reinstall(ctx context.Context, id string, input CloudComputingInstanceReinstallInput) (*CloudComputingInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudInstanceReinstallPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Rescue cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/ActivateRescueModeForACloudInstance
func (h *CloudComputingInstancesHandler) Rescue(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstanceRescuePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Unrescue cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/DeactivateRescueModeForACloudInstance
func (h *CloudComputingInstancesHandler) Unrescue(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstanceUnrescuePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Upgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/UpgradeACloudInstance
func (h *CloudComputingInstancesHandler) Upgrade(ctx context.Context, id string, input CloudComputingInstanceUpgradeInput) (*CloudComputingInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(cloudInstanceUpgradePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// RevertUpgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/RevertUpgradeForACloudInstance
func (h *CloudComputingInstancesHandler) RevertUpgrade(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstanceRevertUpgradePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// ApproveUpgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/ApproveUpgradeForACloudInstance
func (h *CloudComputingInstancesHandler) ApproveUpgrade(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstanceApproveUpgradePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PowerOn cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/PowerOnACloudInstance
func (h *CloudComputingInstancesHandler) PowerOn(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstancePowerOnPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PowerOff cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/PowerOffACloudInstance
func (h *CloudComputingInstancesHandler) PowerOff(ctx context.Context, id string) (*CloudComputingInstance, error) {
	url := h.client.buildURL(cloudInstancePowerOffPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudComputingInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PTRRecords builds a new Collection[PTRRecord] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/ListPtrRecordsForACloudInstance
func (h *CloudComputingInstancesHandler) PTRRecords(id string) Collection[PTRRecord] {
	path := h.client.buildPath(cloudInstancePTRsListPath, []interface{}{id}...)

	return NewCollection[PTRRecord](h.client, path)
}

// CreatePTRRecord creates ptr record for the cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/CreateAPtrRecordForACloudInstance
func (h *CloudComputingInstancesHandler) CreatePTRRecord(ctx context.Context, cloudInstanceID string, input PTRRecordCreateInput) (*PTRRecord, error) {
	url := h.client.buildURL(cloudInstanceCreatePTRRecordPath, []interface{}{cloudInstanceID}...)

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

// DeletePTRRecord deleted ptr record for the cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance/operation/DeleteAPtrRecordForACloudInstance
func (h *CloudComputingInstancesHandler) DeletePTRRecord(ctx context.Context, cloudInstanceID string, ptrRecordID string) error {
	url := h.client.buildURL(cloudInstanceDeletePTRRecordPath, []interface{}{cloudInstanceID, ptrRecordID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
