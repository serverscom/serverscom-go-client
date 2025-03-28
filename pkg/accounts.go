package serverscom

import (
	"context"
	"encoding/json"
)

const (
	accountBalancePath = "/account/balance"
)

// AccountService is an interface for interfacing with accoount endpoints
// API documentation:
// https://developers.servers.com/api-documentation/v1/#tag/Account
type AccountService interface {

	// Generic operations
	GetBalance(ctx context.Context) (*AccountBalance, error)
}

// AccountHandler handles operations around account
type AccountHandler struct {
	client *Client
}

// GetBalance returns account balance information
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Account/operation/GetCurrentAccountBalance
func (h *AccountHandler) GetBalance(ctx context.Context) (*AccountBalance, error) {
	url := h.client.baseURL + accountBalancePath

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	balance := new(AccountBalance)

	if err := json.Unmarshal(body, &balance); err != nil {
		return nil, err
	}

	return balance, nil
}
