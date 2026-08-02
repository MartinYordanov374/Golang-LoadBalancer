package HelperFunctions
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang-loadbalancer/Structs"
)

func SetUpCustomMetrics(PrometheusRegistry prometheus.Registerer) *Structs.PrometheusMetrics{
	CustomMetrics := &Structs.PrometheusMetrics{
		TotalRequests: promauto.With(PrometheusRegistry).NewCounter(prometheus.CounterOpts{
			Name: "total_requests",
			Help: "The total requests sent toward the load balancer",
		}),
		LoadBalancerResponseLatency: promauto.With(PrometheusRegistry).NewHistogram(prometheus.HistogramOpts{
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
		BackendsCount: promauto.With(PrometheusRegistry).NewGauge(prometheus.GaugeOpts{
			Name: "total_backends_count",
			Help: "The total amount of backends that the load balancer is routing to",
		}),
		HealthyBackendsCount: promauto.With(PrometheusRegistry).NewGauge(prometheus.GaugeOpts{
			Name: "healthy_backends_count",
			Help: "The amount of backends that the health check succeeded for",
		}),
		BackendDowntimeDuration: promauto.With(PrometheusRegistry).NewHistogramVec(prometheus.HistogramOpts{
			Name: "backend_downtime_duration",
			Help: "The amount of time which a backend has been down for",
			Buckets: []float64{
				5,
				10,
				15,
				20,
				25,
				30,
				60,
				80,
				120,
				240,
				360,
			},
		},
		[]string{"backend"}),
		TotalHealthCheckFailures: promauto.With(PrometheusRegistry).NewCounter(prometheus.CounterOpts{
			Name: "total_healthcheck_failures",
			Help: "The total amount of healtchecks that have failed, i.e. server was not detected as running",
		}),
		BackendHealthCheckFailures: promauto.With(PrometheusRegistry).NewCounterVec(prometheus.CounterOpts{
			Name: "backend_healthcheck_failures",
			Help: "The total amount of times a specific backend has failed health checks",
		},
			[]string{"backend"}),
		BackendsOnCooldown: promauto.With(PrometheusRegistry).NewGauge(prometheus.GaugeOpts{
			Name: "backends_on_cooldown",
			Help: "The backends currently on cooldown",
		}),
		TotalSuccessfulHealthChecks: promauto.With(PrometheusRegistry).NewCounter(prometheus.CounterOpts{
			Name: "total_healthcheck_successful",
			Help: "The total amount of successful health checks",
		}),
		BackendCooldownsCounter: promauto.With(PrometheusRegistry).NewGaugeVec(prometheus.GaugeOpts{
			Name: "backend_cooldowns_count",
			Help: "The total amount of cooldowns that a backend was put on after repeatedly failing health checks",
		},
			[]string{"backend"}),
	}
	return CustomMetrics
}
