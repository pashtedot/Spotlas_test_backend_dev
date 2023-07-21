package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	IncCounter()
	SetsRepositoryGauge(operation string, duration time.Duration)
}

// Prometheus - struct for metrics
type Prometheus struct {
	SomeCounter prometheus.Counter
	// measures duration of a db query
	RepositoryGauge *prometheus.GaugeVec
}

// NewPrometheus creates new Prometheus instance
func NewPrometheus() *Prometheus {
	prefix := "project"
	metrics := &Prometheus{
		SomeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name:      "some_counter",
			Namespace: prefix,
		}),
		RepositoryGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:      "db_duration",
				Namespace: prefix,
			},
			[]string{"operation"},
		),
	}

	prometheus.MustRegister(metrics.SomeCounter)
	prometheus.MustRegister(metrics.RepositoryGauge)

	return metrics
}

// IncCounter add 1 to counter
func (m *Prometheus) IncCounter() {
	m.SomeCounter.Inc()
}

// SetsRepositoryGauge sets duration gauges for chosen processes
func (m *Prometheus) SetsRepositoryGauge(operation string, duration time.Duration) {
	labels := prometheus.Labels{
		"operation": operation,
	}
	m.RepositoryGauge.With(labels).Set(float64(duration) / float64(time.Second))
}
