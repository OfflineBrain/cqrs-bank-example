package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TotalDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "kafka",
		Name:       "consumer_total_duration",
		Help:       "Duration of consumed events",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"event"})

	TotalDbAccessDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "db",
		Name:       "total_duration",
		Help:       "Duration of DB access",
		Objectives: nil,
	}, []string{"operation", "function"})

	KafkaErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "kafka",
		Name:      "consumer_errors",
		Help:      "Count of receive errors",
	})
)

func RegisterMetrics() {
	prometheus.MustRegister(
		TotalDuration,
		TotalDbAccessDuration,
		KafkaErrors,
	)
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
