package gohttpbreaker

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
)

var defaultHistogramBuckets = []float64{
	0.001, // 1ms
	0.005, // 5ms
	0.01,  // 10ms
	0.025, // 25ms
	0.05,  // 50ms
	0.075, // 75ms
	0.1,   // 100ms
	0.15,  // 150ms
	0.2,   // 200ms
	0.5,   // 500ms
	1,     // 1s
	1.5,   // 1.5s
	2,     // 2s
	4,     // 4s
	6,     // 6s
	8,     // 8s
	10,    // 10s
}

type customLabelValuesKeyType string

const (
	customLabelValuesKey = customLabelValuesKeyType("customLabelValues")
	reqsName             = "http_client_requests"
	latencyName          = "http_client_requests_seconds"
	cbName               = "circuit_breaker_requests"
)

var defaultTransformUrlRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

type ClientPromMiddleware struct {
	reqs             *prometheus.CounterVec
	latency          *prometheus.HistogramVec
	transformUrlFn   func(string) string
	name             string
	labels           []string
	histogramBuckets []float64
}

func ApplyDefaultUrlTransformation(rawUrl string) string {
	reqURL := defaultTransformUrlRegex.ReplaceAllString(rawUrl, ":id")
	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		parsedURL = &url.URL{}
	}

	return parsedURL.Path
}

func (mw *ClientPromMiddleware) MiddlewareFunc() func(*resty.Client, *resty.Response) error {

	if mw.histogramBuckets == nil {
		mw.histogramBuckets = defaultHistogramBuckets
	}

	if mw.transformUrlFn == nil {
		mw.transformUrlFn = ApplyDefaultUrlTransformation
	}

	mw.reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        reqsName,
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"clientName": mw.name},
		},
		append([]string{"status", "method", "path"}, mw.labels...),
	)
	err := prometheus.Register(mw.reqs)
	if err != nil {
		alreadyRegistered := &prometheus.AlreadyRegisteredError{}
		if errors.As(err, alreadyRegistered) {
			mw.reqs = alreadyRegistered.ExistingCollector.(*prometheus.CounterVec)
		}
	}

	mw.latency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        latencyName,
			Help:        "How long it took to process the request, partitioned by status code, method and HTTP path.",
			ConstLabels: prometheus.Labels{"clientName": mw.name},
			Buckets:     mw.histogramBuckets,
		},
		append([]string{"status", "method", "path"}, mw.labels...),
	)
	err = prometheus.Register(mw.latency)
	if err != nil {
		alreadyRegistered := &prometheus.AlreadyRegisteredError{}
		if errors.As(err, alreadyRegistered) {
			mw.latency = alreadyRegistered.ExistingCollector.(*prometheus.HistogramVec)
		}
	}

	return func(c *resty.Client, res *resty.Response) error {
		statusCode := strconv.Itoa(res.StatusCode())
		duration := res.Time()
		method := res.Request.Method

		var path string
		url := res.Request.Context().Value(requestUrlContextKey)
		if url, ok := url.(string); ok {
			path = mw.transformUrlFn(url)
		} else {
			path = mw.transformUrlFn(res.Request.URL)
		}

		labelValues := []string{statusCode, method, path}
		if len(mw.labels) > 0 {
			customLabelValues := []string{}
			ctx := res.Request.Context()

			customLabelValuesFromCtx := ctx.Value(customLabelValuesKey)
			if customLabelValuesFromCtx != nil {
				customLabelValuesFromCtx, ok := customLabelValuesFromCtx.([]string)
				if ok {
					customLabelValues = customLabelValuesFromCtx
				}
			}

			lenProvided := len(customLabelValues)
			lenExpected := len(mw.labels)

			if lenProvided > lenExpected {
				customLabelValues = customLabelValues[:lenExpected]
			} else if lenProvided < lenExpected {
				filler := make([]string, lenExpected-lenProvided)
				customLabelValues = append(customLabelValues, filler...)
			}

			labelValues = append(labelValues, customLabelValues...)
		}

		mw.reqs.WithLabelValues(labelValues...).Inc()
		mw.latency.WithLabelValues(labelValues...).Observe(duration.Seconds())

		return nil
	}
}

type CircuitBreakerPrometheus struct {
	cbCounterVec *prometheus.CounterVec
}

func newCircuitBreakerPrometheus(clientName string) *CircuitBreakerPrometheus {
	cbp := &CircuitBreakerPrometheus{
		cbCounterVec: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name:        cbName,
				Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
				ConstLabels: prometheus.Labels{"clientName": clientName},
			}, []string{"state"},
		),
	}
	_ = prometheus.Register(cbp.cbCounterVec)

	return cbp
}

func (cbp *CircuitBreakerPrometheus) RegisterEvent(err error) {
	event := errorToEvent(err)
	cbp.cbCounterVec.WithLabelValues(event).Inc()
}

func errorToEvent(err error) string {
	event := "failure"
	switch err {
	case nil:
		event = "success"
	case context.DeadlineExceeded:
		event = "context-deadline-exceeded"
	case context.Canceled:
		event = "context-canceled"
	case gobreaker.ErrTooManyRequests:
		event = "too-many-requests"
	case gobreaker.ErrOpenState:
		event = "circuit-open"
	}

	return event
}
