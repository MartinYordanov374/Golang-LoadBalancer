package HelperFunctions
import(
	"golang-loadbalancer/Structs"
	"time"
)
func UpdateServerHealthState(TargetServer *Structs.Server, StatusCode int, CustomMetrics *Structs.PrometheusMetrics){
	if StatusCode == 200{
		TargetServer.IsUp.Store(true)
		if !TargetServer.WentDownTimeStamp.IsZero(){
				BackendDowntimeDuration := time.Since(TargetServer.WentDownTimeStamp).Seconds()
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
}
