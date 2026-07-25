package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"golang-loadbalancer/GlobalVariables"
)
func LoadHealthyServersList() []*Structs.Server{
	TempHealthyServersList := GlobalVariables.HealthyServersList.Load()
	if TempHealthyServersList != nil{
		LoadedHealthyServersList, _ := TempHealthyServersList.([]*Structs.Server)
		return LoadedHealthyServersList
	}else{
		return nil
	}
}
