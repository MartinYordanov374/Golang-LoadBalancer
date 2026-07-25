package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"sync/atomic"
	"net/http/httputil"
	"net/url"
	"golang-loadbalancer/HelperFunctions"
	"golang-loadbalancer/Structs"
)


var HealthyServersList atomic.Value
var CurrentServerIdx uint32 = 0

var HttpClient = &http.Client{
	Timeout: 5*time.Second,
}

func main() {
	go func() {
		OriginURL, Error := url.Parse("http://127.0.0.1:0")
		if Error != nil{
			log.Println(Error)
		}
		ReverseProxy := httputil.NewSingleHostReverseProxy(OriginURL)

		OriginalDirector := ReverseProxy.Director
		ReverseProxy.Director = func(req *http.Request){
			OriginalDirector(req)
			HealthyServers := LoadHealthyServersList()
			if len(HealthyServers) > 0{
				NextServer := HealthyServers[CurrentServerIdx]
				req.URL.Scheme = "http"
				req.URL.Host = fmt.Sprintf("%s:%d", NextServer.Host, NextServer.Port)
				req.Host = req.URL.Host
			}else{
				log.Println("No servers up")
			}
		FindNextServerIdx()
		}
		http.Handle("/", ReverseProxy)
		http.ListenAndServe(":8000", nil)
	}()

	ServersList := HelperFunctions.InitializeServersList()

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		wg := new(sync.WaitGroup)
		wg.Add(len(ServersList))
		for _, Server := range ServersList {
			log.Println(Server)
			go PerformHealthCheck(Server, wg)
		}
		wg.Wait()

		UpServersList := make([]*Structs.Server, 0)
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



func PerformHealthCheck(TargetServer *Structs.Server, WaitGroup *sync.WaitGroup){

	defer WaitGroup.Done()
	uri := fmt.Sprintf("http://%s:%d/health", TargetServer.Host, TargetServer.Port)
	resp, err := HttpClient.Get(uri)
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
		ParsedJSON := HelperFunctions.ParseJSONResponse(jsonbody)
		UpdateServerHealthState(TargetServer, ParsedJSON.StatusCode)
	}
}



func UpdateServerHealthState(TargetServer *Structs.Server, StatusCode int){
	// TODO: Rewrite this to use atomic values instead of mutex as apparently atomics introduce less overhead than mutexes
	TargetServer.Mutex.Lock()
	if StatusCode == 200{
		TargetServer.IsUp = true
	}else{
		TargetServer.IsUp = false
	}
	TargetServer.Mutex.Unlock()
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

func LoadHealthyServersList() []*Structs.Server{
	TempHealthyServersList := HealthyServersList.Load()
	if TempHealthyServersList != nil{
		LoadedHealthyServersList, _ := TempHealthyServersList.([]*Structs.Server)
		return LoadedHealthyServersList
	}else{
		return nil
	}
}
