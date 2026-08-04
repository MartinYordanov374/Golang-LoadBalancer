package HelperFunctions

import (
	"golang-loadbalancer/Structs"
)
func InitializeServersList() []*Structs.Server {
	Servers := []*Structs.Server{}

	Servers = append(Servers, &Structs.Server{
		PrometheusLabel:"ServerOne",
		Host: "serverone",
		Port: 8080})

	Servers = append(Servers, &Structs.Server{
		PrometheusLabel: "ServerTwo",
		Host: "servertwo",
		Port: 8081})

	Servers = append(Servers, &Structs.Server{
		PrometheusLabel:"ServerThree",
		Host: "serverthree",
		Port: 8082})
	return Servers
}
