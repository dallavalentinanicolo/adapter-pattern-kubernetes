package metricsProm

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// podPendingGauge tracks the current number of pending pods
	podPendingGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pod_pending",
		Help: "Current number of pending pods",
	})
)

// UpdatePodPendingMetric updates the "pod_pending" metric with the given count
func UpdatePodPendingMetric(pendingPodCount int) {
	podPendingGauge.Set(float64(pendingPodCount))
}

// ExposeMetrics starts an HTTP server to expose the Prometheus metrics at /metrics endpoint.
func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
