package HelperFunctions

import (
	"golang-loadbalancer/Structs"
)
func HandleServerDownCounter(TargetServer *Structs.Server, CustomMetrics *Structs.PrometheusMetrics){
		if TargetServer.DownCount.Load() == 5{
			TargetServer.IsInCooldown.Store(true)
			CustomMetrics.BackendCooldownsCounter.WithLabelValues(TargetServer.PrometheusLabel).Inc()
			CustomMetrics.BackendsOnCooldown.Inc()
			TargetServer.DownCount.Store(0)
			SetServerScanningCooldown(TargetServer)
		}else{
			TargetServer.DownCount.Add(1)

		}
}
