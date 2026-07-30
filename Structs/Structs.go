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
}
type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

type PrometheusMetrics struct {
	TotalRequests prometheus.Counter
	LoadBalancerResponseLatency prometheus.Histogram
}
