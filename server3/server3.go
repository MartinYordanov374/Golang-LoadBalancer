package main 

import(
	"net/http"
	"golang-loadbalancer/endpoints"
)

func main(){

	http.HandleFunc("/", endpoints.Server("3"))

	http.ListenAndServe(":8082", nil)

}
