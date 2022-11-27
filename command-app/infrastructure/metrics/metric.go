package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

var (
	totalRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http",
		Name:      "request_total_count",
		Help:      "Count of HTTP requests",
	}, []string{"path", "method", "code"})

	totalDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "http",
		Name:       "request_total_duration",
		Help:       "Duration of HTTP requests",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"path", "method", "code"})

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

	TotalMongoAccessDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "mongo",
		Name:       "access_total_duration",
		Help:       "Duration of mongo access",
		Objectives: nil,
	}, []string{"operation"})
)

func RegisterMetrics() {
	prometheus.MustRegister(
		totalRequests,
		totalDuration,
		TotalDuration,
		TotalDbAccessDuration,
		KafkaErrors,
		TotalMongoAccessDuration,
	)
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CommonMiddleware(c *gin.Context) {
	path := c.FullPath()
	method := c.Request.Method
	status := strconv.Itoa(http.StatusOK)

	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
		totalDuration.WithLabelValues(path, method, status).Observe(f)
	}))
	c.Next()
	status = strconv.Itoa(c.Writer.Status())

	timer.ObserveDuration()
	totalRequests.WithLabelValues(path, method, status).Inc()
}
