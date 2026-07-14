package main

import (
	"net/http"
	"log"
)

func main(){
	resp, err := http.Get("http://localhost:8080")
	log.Printf(string(resp))
	http.ListenAndServe(":8000", nil)
}
