package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"time"
)
func SetServerScanningCooldown(TargetServer *Structs.Server){
	CurrentTimeStamp:= time.Now()
	CooldownDurationInMinutes := 5*time.Minute
	CooldownEndTimeStamp := CurrentTimeStamp.Add(CooldownDurationInMinutes)
	TargetServer.Mutex.Lock()
	TargetServer.CooldownEndTimeStamp = CooldownEndTimeStamp
	TargetServer.Mutex.Unlock()
}
