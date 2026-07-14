package main 

import(
	"fmt"
	"net/http"
	"golang-loadbalancer/endpoints"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "The server works")
	})

	http.HandleFunc("/health", endpoints.HealthCheck)

	http.ListenAndServe(":8081", nil)

}
