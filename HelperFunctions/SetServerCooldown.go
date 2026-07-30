package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"time"
	"log"
)
func SetServerScanningCooldown(TargetServer *Structs.Server){
	CurrentTimeStamp:= time.Now()
	CooldownDurationInMinutes := 1*time.Minute
	CooldownEndTimeStamp := CurrentTimeStamp.Add(CooldownDurationInMinutes)
	TargetServer.Mutex.Lock()
	TargetServer.CooldownEndTimeStamp = CooldownEndTimeStamp
	TargetServer.Mutex.Unlock()
	log.Println("Cooldown set for server on port, ", TargetServer.Port)
}
