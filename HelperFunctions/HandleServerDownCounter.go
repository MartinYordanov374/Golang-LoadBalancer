package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"log"
)
func HandleServerDownCounter(TargetServer *Structs.Server){
		if TargetServer.DownCount.Load() > 5{
			log.Println("Skipping server ", TargetServer.Port)
			// TODO: Figure out how to count the time until the server
			// will be back to the scanning list
			TargetServer.IsInCooldown.Store(true)
			SetServerScanningCooldown(TargetServer)
			TargetServer.DownCount.Store(0)
			return;
		}else{
			TargetServer.DownCount.Add(1)
		}
}
