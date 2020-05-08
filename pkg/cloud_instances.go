package serverscom

import (
	"context"
	"encoding/json"
)

const (
	cloudInstanceCreatePath    = "/cloud_computing/instances"
	cloudInstancePath          = "/cloud_computing/instances/%s"
	cloudInstanceUpdatePath    = "/cloud_computing/instances/%s"
	cloudInstanceDeletePath    = "/cloud_computing/instances/%s"
	cloudInstanceReinstallPath = "/cloud_computing/instances/%s/reinstall"
	cloudInstanceRescuePath    = "/cloud_computing/instances/%s/rescue"
	cloudInstanceUnrescuePath  = "/cloud_computing/instances/%s/unrescue"
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
