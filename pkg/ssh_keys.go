package serverscom

import (
	"context"
	"encoding/json"
)

const (
	sshKeyListPath   = "/ssh_keys"
	sshKetCreatePath = "/ssh_keys"
	sshKeyPath       = "/ssh_keys/%s"
)

// SSHKeysService is an interface to interfacing with the SSH Key endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/SSH-Key
type SSHKeysService interface {
	// Primary collection
	Collection() Collection[SSHKey]

	// Generic operations
	Get(ctx context.Context, fingerprint string) (*SSHKey, error)
	Create(ctx context.Context, input SSHKeyCreateInput) (*SSHKey, error)
	Update(ctx context.Context, fingerprint string, input SSHKeyUpdateInput) (*SSHKey, error)
	Delete(ctx context.Context, fingerprint string) error
}

// SSHKeysHandler handles operations around ssh keys
type SSHKeysHandler struct {
	client *Client
}

// Collection builds a new Collection[SSHKey] interface
func (h *SSHKeysHandler) Collection() Collection[SSHKey] {
	return NewCollection[SSHKey](h.client, sshKeyListPath)
}

// Get ssh key
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/ShowSshKey
func (h *SSHKeysHandler) Get(ctx context.Context, fingerprint string) (*SSHKey, error) {
	url := h.client.buildURL(sshKeyPath, []interface{}{fingerprint}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	SSHKey := new(SSHKey)

	if err := json.Unmarshal(body, &SSHKey); err != nil {
		return nil, err
	}

	return SSHKey, nil
}

// Create ssh key
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/AddNewSshKey
func (h *SSHKeysHandler) Create(ctx context.Context, input SSHKeyCreateInput) (*SSHKey, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sshKetCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	var SSHKey *SSHKey

	if err := json.Unmarshal(body, &SSHKey); err != nil {
		return nil, err
	}

	return SSHKey, nil
}

// Update ssh key
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/UpdateTheNameOfSshKey
func (h *SSHKeysHandler) Update(ctx context.Context, fingerprint string, input SSHKeyUpdateInput) (*SSHKey, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sshKeyPath, []interface{}{fingerprint}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var SSHKey *SSHKey

	if err := json.Unmarshal(body, &SSHKey); err != nil {
		return nil, err
	}

	return SSHKey, nil
}

// Delete ssh key
// Endpoint: https://developers.servers.com/api-documentation/v1/#operation/DeleteSshKey
func (h *SSHKeysHandler) Delete(ctx context.Context, fingerprint string) error {
	url := h.client.buildURL(sshKeyPath, []interface{}{fingerprint}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
