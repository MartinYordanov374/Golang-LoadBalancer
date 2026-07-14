package endpoints

import (
	"net/http"
	"encoding/json"
)

type HealthCheckResponse struct {
	StatusCode int	`json:"StatusCode"`
	Message	   string	`json:"Message"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthCheckResponse{StatusCode: 200, Message: "The server is up!"})
}
