package serverscom

import (
	"context"
	"encoding/json"
)

const (
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
)

// CloudInstancesService is an interface to interfacing with the Cloud Instance endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Cloud-Instance
type CloudInstancesService interface {
	Collection() CloudInstancesCollection

	Get(ctx context.Context, id string) (*CloudInstance, error)
	Create(ctx context.Context, input CloudInstanceCreateInput) (*CloudInstance, error)
	Update(ctx context.Context, id string, input CloudInstanceUpdateInput) (*CloudInstance, error)
	Delete(ctx context.Context, id string) error

	Reinstall(ctx context.Context, id string, input CloudInstanceReinstallInput) (*CloudInstance, error)

	Rescue(ctx context.Context, id string) (*CloudInstance, error)
	Unrescue(ctx context.Context, id string) (*CloudInstance, error)

	Upgrade(ctx context.Context, id string, input CloudInstanceUpgradeInput) (*CloudInstance, error)
	RevertUpgrade(ctx context.Context, id string) (*CloudInstance, error)
	ApproveUpgrade(ctx context.Context, id string) (*CloudInstance, error)

	PowerOn(ctx context.Context, id string) (*CloudInstance, error)
	PowerOff(ctx context.Context, id string) (*CloudInstance, error)

	PTRRecords(id string) CloudInstancePTRRecordsCollection
	CreatePTRRecord(ctx context.Context, cloudInstanceID string, input PTRRecordCreateInput) (*PTRRecord, error)
	DeletePTRRecord(ctx context.Context, cloudInstanceID string, ptrRecordID string) error
}

// CloudInstancesHandler handles operations around cloud instances
type CloudInstancesHandler struct {
	client *Client
}

// Collection builds a new CloudInstancesCollection interface
func (ci *CloudInstancesHandler) Collection() CloudInstancesCollection {
	return NewCloudInstancesCollection(ci.client)
}

// Get cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ShowCloudInstance
func (ci *CloudInstancesHandler) Get(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstancePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	cloudInstance := new(CloudInstance)

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Create cloud instace
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreateANewCloudInstance
func (ci *CloudInstancesHandler) Create(ctx context.Context, input CloudInstanceCreateInput) (*CloudInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := ci.client.buildURL(cloudInstanceCreatePath)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Update cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateCloudInstance
func (ci *CloudInstancesHandler) Update(ctx context.Context, id string, input CloudInstanceUpdateInput) (*CloudInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := ci.client.buildURL(cloudInstanceUpdatePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Delete cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteInstance
func (ci *CloudInstancesHandler) Delete(ctx context.Context, id string) error {
	url := ci.client.buildURL(cloudInstanceDeletePath, []interface{}{id}...)

	_, err := ci.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// Reinstall cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ReinstallInstanceWithImage
func (ci *CloudInstancesHandler) Reinstall(ctx context.Context, id string, input CloudInstanceReinstallInput) (*CloudInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := ci.client.buildURL(cloudInstanceReinstallPath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Rescue cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/MoveInstanceToRescueState
func (ci *CloudInstancesHandler) Rescue(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstanceRescuePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Unrescue cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ExitFromRescueState
func (ci *CloudInstancesHandler) Unrescue(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstanceUnrescuePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// Upgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpgradeInstance
func (ci *CloudInstancesHandler) Upgrade(ctx context.Context, id string, input CloudInstanceUpgradeInput) (*CloudInstance, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := ci.client.buildURL(cloudInstanceUpgradePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// RevertUpgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/RevertInstanceUpgrade
func (ci *CloudInstancesHandler) RevertUpgrade(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstanceRevertUpgradePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// ApproveUpgrade cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ApproveInstanceUpgrade
func (ci *CloudInstancesHandler) ApproveUpgrade(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstanceApproveUpgradePath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PowerOn cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SwitchPowerOn
func (ci *CloudInstancesHandler) PowerOn(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstancePowerOnPath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PowerOff cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/SwitchPowerOff
func (ci *CloudInstancesHandler) PowerOff(ctx context.Context, id string) (*CloudInstance, error) {
	url := ci.client.buildURL(cloudInstancePowerOffPath, []interface{}{id}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

	if err != nil {
		return nil, err
	}

	var cloudInstance *CloudInstance

	if err := json.Unmarshal(body, &cloudInstance); err != nil {
		return nil, err
	}

	return cloudInstance, nil
}

// PTRRecords builds a new CloudInstancePTRRecordsCollection interface
func (ci *CloudInstancesHandler) PTRRecords(id string) CloudInstancePTRRecordsCollection {
	return NewCloudInstancePTRRecordsCollection(ci.client, id)
}

// CreatePTRRecord creates ptr record for the cloud instance
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/CreatePtrForInstance
func (ci *CloudInstancesHandler) CreatePTRRecord(ctx context.Context, cloudInstanceID string, input PTRRecordCreateInput) (*PTRRecord, error) {
	url := ci.client.buildURL(cloudInstanceCreatePTRRecordPath, []interface{}{cloudInstanceID}...)

	body, err := ci.client.buildAndExecRequest(ctx, "POST", url, nil)

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
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DetetePtrForInstance
func (ci *CloudInstancesHandler) DeletePTRRecord(ctx context.Context, cloudInstanceID string, ptrRecordID string) error {
	url := ci.client.buildURL(cloudInstanceDeletePTRRecordPath, []interface{}{cloudInstanceID, ptrRecordID}...)

	_, err := ci.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
