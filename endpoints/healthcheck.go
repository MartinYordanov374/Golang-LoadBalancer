package endpoints

import (
	"net/http"
	"encoding/json"
	"golang-loadbalancer/Structs"
)


func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Structs.HealthCheckEndpointResponse{StatusCode: 200, Message: "The server is up!"})
}
