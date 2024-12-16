package serverscom

import (
	"context"
	"encoding/json"
)

const (
	sslCertificateListPath     = "/ssl_certificates"
	sslCreatificatedCreatePath = "/ssl_certificates/custom"
	sslCertificatePath         = "/ssl_certificates/custom/%s"
	sslCertificateLEPath       = "/ssl_certificates/letsencrypt/%s"
)

// SSLCertificatesService is an interface to interfacing with the SSL Certificate endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/SSL-Certificate
type SSLCertificatesService interface {
	// Primary collection
	Collection() Collection[SSLCertificate]

	// Generic operations
	CreateCustom(ctx context.Context, input SSLCertificateCreateCustomInput) (*SSLCertificateCustom, error)
	UpdateCustom(ctx context.Context, id string, input SSLCertificateUpdateCustomInput) (*SSLCertificateCustom, error)
	GetCustom(ctx context.Context, id string) (*SSLCertificateCustom, error)
	DeleteCustom(ctx context.Context, id string) error
	GetLE(ctx context.Context, id string) (*SSLCertificateLE, error)
	UpdateLE(ctx context.Context, id string, input SSLCertificateUpdateLEInput) (*SSLCertificateLE, error)
	DeleteLE(ctx context.Context, id string) error
}

// SSLCertificatesHandler handles operations around ssl certificates
type SSLCertificatesHandler struct {
	client *Client
}

// Collection builds a new Collection[SSLCertificate] interface
func (h *SSLCertificatesHandler) Collection() Collection[SSLCertificate] {
	return NewCollection[SSLCertificate](h.client, sslCertificateListPath)
}

// CreateCustom creates a custom ssl certificate
func (h *SSLCertificatesHandler) CreateCustom(ctx context.Context, input SSLCertificateCreateCustomInput) (*SSLCertificateCustom, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sslCreatificatedCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	SSLCertificateCustom := new(SSLCertificateCustom)

	if err := json.Unmarshal(body, &SSLCertificateCustom); err != nil {
		return nil, err
	}

	return SSLCertificateCustom, nil
}

// GetCustom returns a custom ssl certificate
func (h *SSLCertificatesHandler) GetCustom(ctx context.Context, id string) (*SSLCertificateCustom, error) {
	url := h.client.buildURL(sslCertificatePath, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	SSLCertificate := new(SSLCertificateCustom)

	if err := json.Unmarshal(body, &SSLCertificate); err != nil {
		return nil, err
	}

	return SSLCertificate, nil
}

// UpdateCustom updates a custom SSL certificate
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/SSL-Certificate/operation/UpdateACustomSslCertificate
func (h *SSLCertificatesHandler) UpdateCustom(ctx context.Context, id string, input SSLCertificateUpdateCustomInput) (*SSLCertificateCustom, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sslCertificatePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var SSLCertificate SSLCertificateCustom
	if err := json.Unmarshal(body, &SSLCertificate); err != nil {
		return nil, err
	}

	return &SSLCertificate, nil
}

// DeleteCustom deletes a custom SSL certificate
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/SSL-Certificate/operation/DeleteACustomSslCertificate
func (h *SSLCertificatesHandler) DeleteCustom(ctx context.Context, id string) error {
	url := h.client.buildURL(sslCertificatePath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// GetLE returns a Let's Encrypt SSL certificate
func (h *SSLCertificatesHandler) GetLE(ctx context.Context, id string) (*SSLCertificateLE, error) {
	url := h.client.buildURL(sslCertificateLEPath, id)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	var SSLCertificate SSLCertificateLE

	if err := json.Unmarshal(body, &SSLCertificate); err != nil {
		return nil, err
	}

	return &SSLCertificate, nil
}

// UpdateLE updates a Let's Encrypt SSL certificate
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/SSL-Certificate/operation/UpdateALetsEncryptSslCertificate
func (h *SSLCertificatesHandler) UpdateLE(ctx context.Context, id string, input SSLCertificateUpdateLEInput) (*SSLCertificateLE, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(sslCertificateLEPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	var SSLCertificate SSLCertificateLE
	if err := json.Unmarshal(body, &SSLCertificate); err != nil {
		return nil, err
	}

	return &SSLCertificate, nil
}

// DeleteLE deletes a Let's Encrypt SSL certificate
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/SSL-Certificate/operation/DeleteALetsEncryptSslCertificate
func (h *SSLCertificatesHandler) DeleteLE(ctx context.Context, id string) error {
	url := h.client.buildURL(sslCertificateLEPath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
