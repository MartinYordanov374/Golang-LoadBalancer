package main

import (
	"net/http"
	"io"
	"log"
	"github.com/google/uuid"
)

type Server struct {
	ID string
	Host string
	Port int
	IsUp bool
}

func main(){
	go func(){
		http.ListenAndServe(":8000", nil)
	}()

	Servers := make(map[string]Server)

	Servers["ServerOne"] = Server{ID: uuid.New().String(), Host: "localhost", Port: 8080, IsUp: false}
	Servers["ServerTwo"] = Server{ID: uuid.New().String(), Host: "localhost", Port: 8081, IsUp: false}
	Servers["ServerThree"] = Server{ID: uuid.New().String(), Host: "localhost", Port: 8082, IsUp: false}


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
