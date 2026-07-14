package endpoints

import (
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request){
	log.Printf("server is up")
}
