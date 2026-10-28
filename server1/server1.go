package main 

import(
	"net/http"
	"golang-loadbalancer/endpoints"
)

func main(){

	http.HandleFunc("/", endpoints.Server("1"))

	http.ListenAndServe(":8080", nil)

}
