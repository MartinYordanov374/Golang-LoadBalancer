package Structs

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	ID   string
	Host string
	Port int
	IsUp bool
	Mutex sync.RWMutex
}

type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

type PrometheusMetrics struct {
	TotalRequests prometheus.Counter
}
