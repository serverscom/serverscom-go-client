package serverscom

import (
	"context"
	"encoding/json"
)

const (
	invoicesListPath = "/billing/invoices"

	InvoicePath = "/billing/invoices/%s"
)

// BilligService is an interface for interfacing with Billing Invoices endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Invoice
type InvoiceService interface {
	// Primary collection
	Collection() Collection[InvoiceList]

	// Generic operations
	GetBillingInvoice(ctx context.Context, id string) (*Invoice, error)
}

// InvoiceHandler handles operations around hosts
type InvoiceHandler struct {
	client *Client
}

// Collection builds a new Collection[InvoiceList] interface
func (h *InvoiceHandler) Collection() Collection[InvoiceList] {
	return NewCollection[InvoiceList](h.client, invoicesListPath)
}

// GetBillingInvoice returns an invoice
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Invoice/operation/GetAnInvoice
func (h *InvoiceHandler) GetBillingInvoice(ctx context.Context, id string) (*Invoice, error) {
	url := h.client.buildURL(InvoicePath, []interface{}{id}...)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	invoice := new(Invoice)

	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, err
	}

	return invoice, nil
}
