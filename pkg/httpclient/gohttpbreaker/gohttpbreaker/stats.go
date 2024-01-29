package gohttpbreaker

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
)

const (
	namespace = "http"
	subsystem = "circuit_breaker"
)

// StatsCollector implements the prometheus.Collector interface.
type StatsCollector struct {
	cb *gobreaker.CircuitBreaker

	// descriptions of exported metrics
	state *prometheus.Desc
}

// NewStatsCollector creates a new StatsCollector.
func newPrometheusStatsCollector(clientName string, cb *gobreaker.CircuitBreaker) *StatsCollector {
	labels := prometheus.Labels{"clientName": clientName}

	return &StatsCollector{
		cb: cb,
		state: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "state"),
			"Return states from circuit breaker.",
			[]string{"state"},
			labels,
		),
	}
}

// Describe implements the prometheus.Collector interface.
func (c StatsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.state
}

// Collect implements the prometheus.Collector interface.
func (c StatsCollector) Collect(ch chan<- prometheus.Metric) {
	state := c.cb.State()
	stateClosed, stateHalfOpen, stateOpen := c.stateToCode(state)

	ch <- prometheus.MustNewConstMetric(
		c.state,
		prometheus.GaugeValue,
		stateClosed,
		"closed",
	)
	ch <- prometheus.MustNewConstMetric(
		c.state,
		prometheus.GaugeValue,
		stateHalfOpen,
		"half_open",
	)
	ch <- prometheus.MustNewConstMetric(
		c.state,
		prometheus.GaugeValue,
		stateOpen,
		"open",
	)
}

func (c StatsCollector) stateToCode(state gobreaker.State) (float64, float64, float64) {

	stateClosed, stateHalfOpen, stateOpen := 0.0, 0.0, 0.0

	switch state {
	case 0:
		stateClosed = 1.0
	case 1:
		stateHalfOpen = 1.0
	case 2:
		stateOpen = 1.0
	}

	return stateClosed, stateHalfOpen, stateOpen
}
