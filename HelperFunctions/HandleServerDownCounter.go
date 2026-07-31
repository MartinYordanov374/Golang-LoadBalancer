package HelperFunctions

import (
	"golang-loadbalancer/Structs"
	"log"
)
func HandleServerDownCounter(TargetServer *Structs.Server, CustomMetrics *Structs.PrometheusMetrics){
		if TargetServer.DownCount.Load() == 5{
			log.Println("Skipping server ", TargetServer.Port)
			// TODO: Figure out how to count the time until the server
			// will be back to the scanning list
			TargetServer.IsInCooldown.Store(true)
			CustomMetrics.BackendCooldownsCounter.WithLabelValues(TargetServer.PrometheusLabel).Inc()
			CustomMetrics.BackendsOnCooldown.Inc()
			TargetServer.DownCount.Store(0)
			SetServerScanningCooldown(TargetServer)
		}else{
			TargetServer.DownCount.Add(1)

		}
}
