package HelperFunctions
import(
	"golang-loadbalancer/Structs"
)
func UpdateServerHealthState(TargetServer *Structs.Server, StatusCode int){
	// TODO: Rewrite this to use atomic values instead of mutex as apparently atomics introduce less overhead than mutexes
	TargetServer.Mutex.Lock()
	if StatusCode == 200{
		TargetServer.IsUp = true
	}else{
		TargetServer.IsUp = false
	}
	TargetServer.Mutex.Unlock()
}
