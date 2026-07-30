package HelperFunctions

import (
	"golang-loadbalancer/Structs"
)


func InitializeServersList() []*Structs.Server {
	// TODO: Automate the initialization process
	Servers := []*Structs.Server{}

	Servers = append(Servers, &Structs.Server{
		Host: "serverone",
		Port: 8080})

	Servers = append(Servers, &Structs.Server{
		Host: "servertwo",
		Port: 8081})

	Servers = append(Servers, &Structs.Server{
		Host: "serverthree",
		Port: 8082})

	return Servers

}
