package repo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalEventsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "total_events_processed",
		Help: "Total number of events processed by retranslator",
	})
)

// TotalEventsProcessedAdd ...
func TotalEventsProcessedAdd(v float64) {
	totalEventsProcessed.Add(v)
}
