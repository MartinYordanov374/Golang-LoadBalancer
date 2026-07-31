package Structs

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
	"time"
)

type Server struct {
	Host string
	Port int
	IsUp atomic.Bool
	Mutex sync.RWMutex
	DownCount atomic.Uint32
	IsInCooldown atomic.Bool
	CooldownEndTimeStamp time.Time
	CooldownStartTimeStamp time.Time
	PrometheusLabel string
}
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
}
