package Structs

import (
	"sync"
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
