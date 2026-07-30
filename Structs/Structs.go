package Structs

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
)

type Server struct {
	ID   string
	Host string
	Port int
	IsUp bool
	Mutex sync.RWMutex
	DownCount atomic.Uint32
}

type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

type PrometheusMetrics struct {
	TotalRequests prometheus.Counter
	LoadBalancerResponseLatency prometheus.Histogram
}
