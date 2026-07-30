package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"net/http/httputil"
	"net/url"
	"golang-loadbalancer/HelperFunctions"
  	"golang-loadbalancer/Structs"
    "golang-loadbalancer/GlobalVariables"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main(){

	go func() {
		OriginURL, Error := url.Parse("http://127.0.0.1:0")
		if Error != nil{
			log.Println(Error)
		}

		PrometheusRegistry := prometheus.NewRegistry()
		CustomMetrics := HelperFunctions.SetUpCustomMetrics(PrometheusRegistry)
		http.Handle("/metrics", promhttp.HandlerFor(PrometheusRegistry, promhttp.HandlerOpts{}))

		ReverseProxy := httputil.NewSingleHostReverseProxy(OriginURL)
		OriginalDirector := ReverseProxy.Director
		ReverseProxy.Director = func(req *http.Request){
			OriginalDirector(req)
			CustomMetrics.TotalRequests.Inc()
			HealthyServers := HelperFunctions.LoadHealthyServersList()
			if len(HealthyServers) > 0{
				RequestStartTime := time.Now()
				NextServer := HealthyServers[GlobalVariables.CurrentServerIdx.Load()]
				req.URL.Scheme = "http"
				req.URL.Host = fmt.Sprintf("%s:%d", NextServer.Host, NextServer.Port)
				req.Host = req.URL.Host
				RequestEndTime := time.Since(RequestStartTime).Seconds()
				CustomMetrics.LoadBalancerResponseLatency.Observe(RequestEndTime)
			}else{
				log.Println("No servers up")
			}
		HelperFunctions.FindNextServerIdx()
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
			go PerformHealthCheck(Server, wg)
		}
		wg.Wait()
		UpServersList := make([]*Structs.Server, 0)
		for _, CurrentServer := range ServersList{
			CurrentServer.Mutex.Lock()
			if CurrentServer.IsUp.Load() == true{
				UpServersList = append(UpServersList, CurrentServer)
			}
			CurrentServer.Mutex.Unlock()
		}
		GlobalVariables.HealthyServersList.Store(UpServersList)
	}
}

func PerformHealthCheck(TargetServer *Structs.Server, WaitGroup *sync.WaitGroup){
	defer WaitGroup.Done()
	uri := fmt.Sprintf("http://%s:%d/health", TargetServer.Host, TargetServer.Port)
	resp, err := GlobalVariables.HttpClient.Get(uri)
	if err != nil {
		HelperFunctions.UpdateServerHealthState(TargetServer, 503)
		return
	}
	defer resp.Body.Close()
	jsonbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}else{
		ParsedJSON := HelperFunctions.ParseJSONResponse(jsonbody)
		HelperFunctions.UpdateServerHealthState(TargetServer, ParsedJSON.StatusCode)
	}
}
