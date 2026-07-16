package main

import (
	"net/http"
	"io"
	"log"
)

func main(){
	go func(){
		http.ListenAndServe(":8000", nil)
	}()

	PerformHealthCheck()


}

func PerformHealthCheck(){
	resp, err := http.Get("http://localhost:8080/health")

	if err != nil {
		log.Println("Health check failed for server 1")
		log.Println(err)
		select{}
	}
	defer resp.Body.Close()

	jsonbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		select{}
	}
	log.Println(string(jsonbody))

	select{}
}
