package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	ID   string
	Host string
	Port int
	IsUp bool
}

func main() {
	go func() {
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

func PerformHealthCheck(TargetServer *Server, WaitGroup *sync.WaitGroup) {

	defer WaitGroup.Done()
	uri := fmt.Sprintf("http://%s:%d/health", TargetServer.Host, TargetServer.Port)
	resp, err := http.Get(uri)

	if err != nil {
		log.Println("Health check failed for ", TargetServer.ID)
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	jsonbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(jsonbody))
}
