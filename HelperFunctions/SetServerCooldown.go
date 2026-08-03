package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"time"
)
func SetServerScanningCooldown(TargetServer *Structs.Server){
	CurrentTimeStamp:= time.Now()
	// TODO: Return the cooldown back to 5 minutes once done with testing
	CooldownDurationInMinutes := 15*time.Minute
	CooldownEndTimeStamp := CurrentTimeStamp.Add(CooldownDurationInMinutes)
	TargetServer.Mutex.Lock()
	TargetServer.CooldownEndTimeStamp = CooldownEndTimeStamp
	TargetServer.CooldownStartTimeStamp = CurrentTimeStamp
	TargetServer.Mutex.Unlock()
}
