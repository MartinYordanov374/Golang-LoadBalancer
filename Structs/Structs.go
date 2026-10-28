package Structs

import (
	"github.com/prometheus/client_golang/prometheus"
)

type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

type PrometheusMetrics struct {
	TotalRequests prometheus.Counter
	LoadBalancerResponseLatency prometheus.Histogram
	BackendsCount prometheus.Gauge
	HealthyBackendsCount prometheus.Gauge
	BackendDowntimeDuration *prometheus.HistogramVec
	TotalHealthCheckFailures prometheus.Counter
	BackendHealthCheckFailures *prometheus.CounterVec
	BackendsOnCooldown prometheus.Gauge
	TotalSuccessfulHealthChecks prometheus.Counter
	BackendCooldownsCounter *prometheus.GaugeVec
}
