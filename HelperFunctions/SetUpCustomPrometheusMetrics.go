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
		LoadBalancerResponseLatency: promauto.With(PrometheusRegistry).NewHistogram(prometheus.HistogramOpts{
			Namespace: "golang-loadbalancer",
			Name: "loadbalancer_response_latency",
			Help: "The latency of the responses returned from the loadbalancer",
			Buckets: []float64{
				0.005,
				0.01,
				0.025,
				0.05,
				0.1,
				0.25,
				0.5,
				1.0,
				2.5,
				5.0,
				10.0,
			},
		}),
	}
	return CustomMetrics
}
