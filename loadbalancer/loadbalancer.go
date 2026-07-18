package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"encoding/json"
)

type Server struct {
	ID   string
	Host string
	Port int
	IsUp bool
	Mutex sync.RWMutex
}

type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

func main() {
	go func() {
		// TODO: Consider adding a handler here instead of nil
		http.ListenAndServe(":8000", nil)
	}()

	ServersList := InitializeServersList()

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		wg := new(sync.WaitGroup)
		wg.Add(len(ServersList))
		for _, Server := range ServersList {
			go PerformHealthCheck(Server, wg)
		}
		wg.Wait()
	}
}

func InitializeServersList() map[string]*Server {

	Servers := make(map[string]*Server)

	Servers["ServerOne"] = &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8080,
		IsUp: false}

	Servers["ServerTwo"] = &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8081,
		IsUp: false}

	Servers["ServerThree"] = &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8082,
		IsUp: false}

	return Servers

}

func PerformHealthCheck(TargetServer *Server, WaitGroup *sync.WaitGroup){

	defer WaitGroup.Done()
	uri := fmt.Sprintf("http://%s:%d/health", TargetServer.Host, TargetServer.Port)
	resp, err := http.Get(uri)

	if err != nil {
		UpdateServerHealthState(TargetServer, 503)
		return
	}
	defer resp.Body.Close()

	jsonbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}else{
		ParsedJSON := ParseJSONResponse(jsonbody)
		UpdateServerHealthState(TargetServer, ParsedJSON.StatusCode)
	}
}

func ParseJSONResponse(JSONResponse []byte) HealthCheckEndpointResponse {

	var ParsedJSON HealthCheckEndpointResponse
	json.Unmarshal([]byte(JSONResponse), &ParsedJSON)
	return ParsedJSON
}

func UpdateServerHealthState(TargetServer *Server, StatusCode int){

	TargetServer.Mutex.RLock()
	if StatusCode == 200{
		TargetServer.IsUp = true
	}else{
		TargetServer.IsUp = false
	}
	TargetServer.Mutex.RUnlock()
}
