package main 

import(
	"net/http"
	"golang-loadbalancer/endpoints"
	"fmt"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Server 1")
	})

	http.HandleFunc("/health", endpoints.HealthCheck)

	http.ListenAndServe(":8080", nil)

}
