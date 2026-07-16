package serverscom

import (
	"context"
	"encoding/json"
)

const (
	dnsDomainListPath       = "/dns/domains"
	dnsDomainCreatePath     = "/dns/domains"
	dnsDomainDelegationPath = "/dns/domains/delegation_data"
	dnsDomainPath           = "/dns/domains/%s"
	dnsRecordListPath       = "/dns/domains/%s/records"
	dnsRecordCreatePath     = "/dns/domains/%s/records"
	dnsRecordPath           = "/dns/domains/%s/records/%s"
)

// DNSService is an interface to interfacing with the DNS domain and record endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain

type DNSService interface {
	// Primary collection
	Collection() Collection[DNSDomain]

	// Domain operations
	CreateDomain(ctx context.Context, input DNSDomainCreateInput) (*DNSDomain, error)
	GetDomain(ctx context.Context, id string) (*DNSDomain, error)
	UpdateDomain(ctx context.Context, id string, input DNSDomainUpdateInput) (*DNSDomain, error)
	DeleteDomain(ctx context.Context, id string) error
	GetDelegationData(ctx context.Context) (*DNSDomainDelegationData, error)

	// Record sub-collection and operations
	Records(domainID string) Collection[DNSRecord]
	CreateRecord(ctx context.Context, domainID string, input DNSRecordCreateInput) (*DNSRecord, error)
	GetRecord(ctx context.Context, domainID, recordID string) (*DNSRecord, error)
	UpdateRecord(ctx context.Context, domainID, recordID string, input DNSRecordUpdateInput) (*DNSRecord, error)
	DeleteRecord(ctx context.Context, domainID, recordID string) error
}

// DNSHandler handles operations around dns domains and records
type DNSHandler struct {
	client *Client
}

// Collection builds a new Collection[DNSDomain] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/ListDnsDomains
func (h *DNSHandler) Collection() Collection[DNSDomain] {
	return NewCollection[DNSDomain](h.client, dnsDomainListPath)
}

// CreateDomain creates a dns domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/CreateDnsDomain
func (h *DNSHandler) CreateDomain(ctx context.Context, input DNSDomainCreateInput) (*DNSDomain, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dnsDomainCreatePath)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	dnsDomain := new(DNSDomain)

	if err := json.Unmarshal(body, &dnsDomain); err != nil {
		return nil, err
	}

	return dnsDomain, nil
}

// GetDomain returns a dns domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/GetDnsDomain
func (h *DNSHandler) GetDomain(ctx context.Context, id string) (*DNSDomain, error) {
	url := h.client.buildURL(dnsDomainPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	dnsDomain := new(DNSDomain)

	if err := json.Unmarshal(body, &dnsDomain); err != nil {
		return nil, err
	}

	return dnsDomain, nil
}

// UpdateDomain updates a dns domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/UpdateDnsDomain
func (h *DNSHandler) UpdateDomain(ctx context.Context, id string, input DNSDomainUpdateInput) (*DNSDomain, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dnsDomainPath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	dnsDomain := new(DNSDomain)

	if err := json.Unmarshal(body, &dnsDomain); err != nil {
		return nil, err
	}

	return dnsDomain, nil
}

// DeleteDomain deletes a dns domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/DeleteDnsDomain
func (h *DNSHandler) DeleteDomain(ctx context.Context, id string) error {
	url := h.client.buildURL(dnsDomainPath, []interface{}{id}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}

// GetDelegationData returns dns domain delegation data
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/GetDnsDomainsDelegationData
func (h *DNSHandler) GetDelegationData(ctx context.Context) (*DNSDomainDelegationData, error) {
	url := h.client.buildURL(dnsDomainDelegationPath)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	delegationData := new(DNSDomainDelegationData)

	if err := json.Unmarshal(body, &delegationData); err != nil {
		return nil, err
	}

	return delegationData, nil
}

// Records builds a new Collection[DNSRecord] interface
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/ListDnsRecordsForADomain
func (h *DNSHandler) Records(domainID string) Collection[DNSRecord] {
	path := h.client.buildPath(dnsRecordListPath, []interface{}{domainID}...)

	return NewCollection[DNSRecord](h.client, path)
}

// CreateRecord creates a dns record for a domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/CreateADnsRecordForADomain
func (h *DNSHandler) CreateRecord(ctx context.Context, domainID string, input DNSRecordCreateInput) (*DNSRecord, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dnsRecordCreatePath, []interface{}{domainID}...)

	body, err := h.client.buildAndExecRequest(ctx, "POST", url, payload)

	if err != nil {
		return nil, err
	}

	dnsRecord := new(DNSRecord)

	if err := json.Unmarshal(body, &dnsRecord); err != nil {
		return nil, err
	}

	return dnsRecord, nil
}

// GetRecord returns a dns record for a domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/GetDnsRecordForDomain
func (h *DNSHandler) GetRecord(ctx context.Context, domainID, recordID string) (*DNSRecord, error) {
	url := h.client.buildURL(dnsRecordPath, []interface{}{domainID, recordID}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	dnsRecord := new(DNSRecord)

	if err := json.Unmarshal(body, &dnsRecord); err != nil {
		return nil, err
	}

	return dnsRecord, nil
}

// UpdateRecord updates a dns record for a domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/UpdateADnsRecordForADomain
func (h *DNSHandler) UpdateRecord(ctx context.Context, domainID, recordID string, input DNSRecordUpdateInput) (*DNSRecord, error) {
	payload, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	url := h.client.buildURL(dnsRecordPath, []interface{}{domainID, recordID}...)

	body, err := h.client.buildAndExecRequest(ctx, "PUT", url, payload)

	if err != nil {
		return nil, err
	}

	dnsRecord := new(DNSRecord)

	if err := json.Unmarshal(body, &dnsRecord); err != nil {
		return nil, err
	}

	return dnsRecord, nil
}

// DeleteRecord deletes a dns record for a domain
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/DNS-Domain/operation/DeleteADnsRecordForADomain
func (h *DNSHandler) DeleteRecord(ctx context.Context, domainID, recordID string) error {
	url := h.client.buildURL(dnsRecordPath, []interface{}{domainID, recordID}...)

	_, err := h.client.buildAndExecRequest(ctx, "DELETE", url, nil)

	return err
}
