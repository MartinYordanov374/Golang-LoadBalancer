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
	"sync/atomic"
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

var HealthyServersList atomic.Value

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

		UpServersList := make([]*Server, 0)
		for _, CurrentServer := range ServersList{
			CurrentServer.Mutex.Lock()
			if CurrentServer.IsUp{
				UpServersList = append(UpServersList, CurrentServer)
			}
			CurrentServer.Mutex.Unlock()
		}
		HealthyServersList.Store(UpServersList)
		log.Println(HealthyServersList)

	}

}

func InitializeServersList() []*Server {
	//TODO: Automate the initialization process
	Servers := []*Server{}

	Servers = append(Servers, &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8080,
		IsUp: false})

	Servers = append(Servers, &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8081,
		IsUp: false})

	Servers = append(Servers, &Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8082,
		IsUp: false})

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

	TargetServer.Mutex.Lock()
	if StatusCode == 200{
		TargetServer.IsUp = true
	}else{
		TargetServer.IsUp = false
	}
	TargetServer.Mutex.Unlock()
}
//TODO: Implement round robin function here. It will redirect all incoming requests(to the load balancer) to the backend servers in a sequential manner.
func RoundRobin(ServersList []*Server, CurrentServerIdx int){
	// 1. Keep track of current server
	// 2. Route to current server and identify the next server
	// 2.1 If the current server was the last in line, return the server counter to the 0-th index
	// 2.2 Identify only up servers and skip the down servers when routing, i.e. if servers 1 and 3 are up, skip 2.

	var NextServerIdx int = FindNextServerIdx(ServersList, CurrentServerIdx)
	log.Println(NextServerIdx)
}

func FindNextServerIdx(ServersList []*Server, CurrentServerIdx int) int {
	//TODO: Add locks in case multiple requests come in at the same time.
	var NextServerIdx = CurrentServerIdx + 1
	var EndServerIdx = len(ServersList) - 1

	if NextServerIdx > EndServerIdx{
		CurrentServerIdx = 0
		return 0
	}else{
		return NextServerIdx
	}
}
