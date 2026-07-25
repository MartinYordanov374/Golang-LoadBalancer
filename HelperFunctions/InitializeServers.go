package HelperFunctions

import (
	"github.com/google/uuid"
	"golang-loadbalancer/Structs"
)


func InitializeServersList() []*Structs.Server {
	// TODO: Automate the initialization process
	Servers := []*Structs.Server{}

	Servers = append(Servers, &Structs.Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8080,
		IsUp: false})

	Servers = append(Servers, &Structs.Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8081,
		IsUp: false})

	Servers = append(Servers, &Structs.Server{
		ID:   uuid.New().String(),
		Host: "localhost",
		Port: 8082,
		IsUp: false})

	return Servers

}
