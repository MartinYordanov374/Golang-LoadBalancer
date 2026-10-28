package main 

import(
	"net/http"
	"golang-loadbalancer/endpoints"
)

func main(){

	http.HandleFunc("/", endpoints.Server("2"))

	http.ListenAndServe(":8081", nil)

}
