package serverscom

import (
	"context"
)

const (
	hostsMetricsPath = "/metrics/hosts"
	racksMetricsPath = "/metrics/racks"
)

// MetricsService is an interface for interfacing with Metrics endpoints
// API documentation: https://developers.servers.com/api-documentation/v1/#tag/Metrics
type MetricsService interface {
	// Generic operations
	ListHostsMetrics(ctx context.Context) (string, error)
	ListRacksMetrics(ctx context.Context) (string, error)
}

// MetricsHandler handles operations around metrics
type MetricsHandler struct {
	client *Client
}

// ListHostsMetrics returns hosts metrics in text exposition format
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Metrics/operation/ListHostsMetrics
func (h *MetricsHandler) ListHostsMetrics(ctx context.Context) (string, error) {
	url := h.client.buildURL(hostsMetricsPath)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// ListRacksMetrics returns racks metrics in text exposition format
// Endpoint: https://developers.servers.com/api-documentation/v1/#tag/Metrics/operation/ListRacksMetrics
func (h *MetricsHandler) ListRacksMetrics(ctx context.Context) (string, error) {
	url := h.client.buildURL(racksMetricsPath)

	body, err := h.client.buildAndExecRequest(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
