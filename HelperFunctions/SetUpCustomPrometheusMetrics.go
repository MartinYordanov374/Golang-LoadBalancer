package HelperFunctions
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang-loadbalancer/Structs"
)

func SetUpCustomMetrics(PrometheusRegistry prometheus.Registerer) *Structs.PrometheusMetrics{
	CustomMetrics := &Structs.PrometheusMetrics{
		TotalRequests: promauto.With(PrometheusRegistry).NewCounter(prometheus.CounterOpts{
			Namespace: "golang-loadbalancer",
			Name: "total_requests",
			Help: "The total requests sent toward the load balancer",
		}),
	}
	return CustomMetrics
}
