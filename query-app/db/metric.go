package db

import "github.com/prometheus/client_golang/prometheus"

var (
	totalDbAccessDuration = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  "db",
		Name:       "total_duration",
		Help:       "Duration of DB access",
		Objectives: nil,
	}, []string{"operation", "function"})
)

func RegisterMetrics() {
	prometheus.MustRegister(totalDbAccessDuration)
}
