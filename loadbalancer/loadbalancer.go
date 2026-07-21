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
var CurrentServerIdx uint32 = 0

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			RoundRobin()
		})
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
	}

}

func InitializeServersList() []*Server {
	// TODO: Automate the initialization process
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
	// TODO: Rewrite this to use atomic values instead of mutex as apparently atomics introduce less overhead than mutexes
	TargetServer.Mutex.Lock()
	if StatusCode == 200{
		TargetServer.IsUp = true
	}else{
		TargetServer.IsUp = false
	}
	TargetServer.Mutex.Unlock()
}

func RoundRobin(){
	log.Println("The current server index is ", atomic.LoadUint32(&CurrentServerIdx))
	LoadedHealthyServersList := LoadHealthyServersList()
	if LoadedHealthyServersList != nil && len(LoadedHealthyServersList) > 0{
		FindNextServerIdx()
		NextServer := LoadedHealthyServersList[CurrentServerIdx]
		// TODO: Build the URL for the selected server and route to it
		NextServerURL := fmt.Sprintf("http://%s:%d/", NextServer.Host, NextServer.Port)
		log.Println(uri)
	}

}

func FindNextServerIdx() {
	LoadedHealthyServersList := LoadHealthyServersList()
	if LoadedHealthyServersList != nil{
		if len(LoadedHealthyServersList) > 0{
			var AtomicCurrentServerIdx = atomic.LoadUint32(&CurrentServerIdx)
			var AtomicHealthyServersListEndIdx = uint32(len(LoadedHealthyServersList)-1)
			if AtomicCurrentServerIdx+1 > AtomicHealthyServersListEndIdx{
				atomic.StoreUint32(&CurrentServerIdx, 0)
			}else{
				atomic.AddUint32(&CurrentServerIdx, 1)
			}
		}
	}
}

func LoadHealthyServersList() []*Server{
	TempHealthyServersList := HealthyServersList.Load()
	if TempHealthyServersList != nil{
		LoadedHealthyServersList, _ := TempHealthyServersList.([]*Server)
		return LoadedHealthyServersList
	}else{
		return nil
	}
}
