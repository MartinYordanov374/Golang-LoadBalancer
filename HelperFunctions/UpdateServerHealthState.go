package HelperFunctions
import(
	"golang-loadbalancer/Structs"
	"time"
	"log"
)
func UpdateServerHealthState(TargetServer *Structs.Server, StatusCode int, CustomMetrics *Structs.PrometheusMetrics){
	// TODO: Rewrite this to use atomic values instead of mutex as apparently atomics introduce less overhead than mutexes
	//TargetServer.Mutex.Lock()
	if StatusCode == 200{
		TargetServer.IsUp.Store(true)
		if !TargetServer.WentDownTimeStamp.IsZero(){
				BackendDowntimeDuration := time.Since(TargetServer.WentDownTimeStamp).Seconds()
				log.Println(BackendDowntimeDuration)
				CustomMetrics.BackendDowntimeDuration.WithLabelValues(TargetServer.PrometheusLabel).Observe(BackendDowntimeDuration)
				TargetServer.WentDownTimeStamp = time.Time{}
		}

	}else{
		TargetServer.IsUp.Store(false)
		if TargetServer.WentDownTimeStamp.IsZero(){
			TargetServer.WentDownTimeStamp = time.Now()
		}
		if !TargetServer.IsInCooldown.Load(){
			HandleServerDownCounter(TargetServer, CustomMetrics)
		}
	}
	//TargetServer.Mutex.Unlock()
}
