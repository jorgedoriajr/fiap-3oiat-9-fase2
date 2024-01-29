package gohttpbreaker

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"hamburgueria/pkg/logger"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	httpdatadogtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

var (
	ErrConfigNotSet                                           = errors.New("config not set")
	ErrNoBaseURL                                              = errors.New("BaseURL must be provided")
	ErrMaxIdleConnectionsPerHostGreaterThanMaxIdleConnections = errors.New("MaxIdleConnectionsPerHost should be <= MaxIdleConnections")
)

type (
	GoHttpBreaker struct {
		config                     *Config
		circuitBreaker             *gobreaker.CircuitBreaker
		httpClient                 *resty.Client
		circuitBreakerPrometheus   *CircuitBreakerPrometheus
		tracer                     trace.Tracer
		transport                  http.RoundTripper
		prometheusCustomLabels     []string
		prometheusHistogramBuckets []float64
		prometheusTransformUrlFn   func(string) string
		onBeforeRequest            []resty.RequestMiddleware
		retryCondition             resty.RetryConditionFunc
		onErrorFunc                []resty.ErrorHook
		onAfterResponseFunc        []resty.ResponseMiddleware
		clientLevelHeaders         map[string]string
		withoutDatadogIntegration  bool
		datadogTransportOptions    []httpdatadogtrace.RoundTripperOption
		httpCodesToIgnore          []int
		statusCheckFn              func(statusCode int) bool
		tlsConfig                  *tls.Config
		circuitIsSuccessfulFn      func(err error) bool
		logger                     zerolog.Logger
	}
	PositiveValueInvalidError struct {
		name  string
		value int
	}
	Option func(*GoHttpBreaker) error
)

func (e PositiveValueInvalidError) Error() string {
	return fmt.Sprintf("supplied value %s = %d must be greater than 0", e.name, e.value)
}

func (ghb *GoHttpBreaker) addClientLevelHeader(headerName string, headerValue string) {
	if ghb.clientLevelHeaders == nil {
		ghb.clientLevelHeaders = make(map[string]string)
	}
	ghb.clientLevelHeaders[headerName] = headerValue
}

func WithConfig(config *Config) Option {
	return func(ghb *GoHttpBreaker) error {
		if config.Client.TimeOutMillis <= 0 {
			return PositiveValueInvalidError{"TimeOutMillis", config.Client.TimeOutMillis}
		}
		if config.Client.MaxWaitTimeMillis <= 0 {
			return PositiveValueInvalidError{"MaxWaitTimeMillis", config.Client.MaxWaitTimeMillis}
		}
		if config.Client.WaitTimeMillis <= 0 {
			return PositiveValueInvalidError{"WaitTimeMillis", config.Client.WaitTimeMillis}
		}
		if config.Client.MaxIdleConnsPerHost > config.Client.MaxIdleConns {
			return ErrMaxIdleConnectionsPerHostGreaterThanMaxIdleConnections
		}
		if config.Client.BaseURL == "" {
			return ErrNoBaseURL
		}
		if config.Name == "" {
			config.Name = config.Client.BaseURL
		}

		ghb.config = config

		ghb.logger = logger.Get()

		return nil
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.tracer = tracer

		return nil
	}
}

func WithTransport(transport http.RoundTripper) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.transport = transport

		return nil
	}
}

func WithCustomHistogramBuckets(buckets []float64) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.prometheusHistogramBuckets = buckets

		return nil
	}
}

func WithCustomPrometheusLabels(labels ...string) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.prometheusCustomLabels = labels

		return nil
	}
}

func ContextWithCustomPrometheusLabelValues(parent context.Context, customLabelValues ...string) context.Context {
	return context.WithValue(parent, customLabelValuesKey, customLabelValues)
}

func WithClientLevelHeader(headerName string, headerValue string) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.addClientLevelHeader(headerName, headerValue)

		return nil
	}
}

func WithRequesterTokenFromEnv() Option {
	return func(ghb *GoHttpBreaker) error {
		token := os.Getenv("REQUESTER_TOKEN")
		if token != "" {
			ghb.addClientLevelHeader("x-requester-token", token)
		}

		return nil
	}
}

func WithPrometheusUrlTransformFn(prometheusUrlTransformFn func(string) string) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.prometheusTransformUrlFn = prometheusUrlTransformFn

		return nil
	}
}

func WithRetryCondition(retryConditionFn resty.RetryConditionFunc) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.retryCondition = retryConditionFn

		return nil
	}
}

// WithOverrideHttpCodesToIgnore Ignores (does not count to open the circuit) any of the given HTTP codes.
// NOTE: Default is {404}
// NOTE: overridden by WithCustomStatusCheck
func WithOverrideHttpCodesToIgnore(codes ...int) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.httpCodesToIgnore = codes

		return nil
	}
}

// WithCustomStatusCheck allows specifying a function for error check based on response's status code
// returning true means "accounted for in open circuit threshold"
// NOTE: overrides WithOverrideHttpCodesToIgnore
func WithCustomStatusCheck(fn func(statusCode int) bool) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.statusCheckFn = fn

		return nil
	}
}

// WithCustomCircuitIsSuccessful allows specifying a custom function for determining if circuit will trigger on an error
// returning true means error won't be accounted for
// NOTE: default behavior is ```return err == nil```
// NOTE: this is useful if i.e.: a cancelled context shouldn't be accounted for
func WithCustomCircuitIsSuccessful(fn func(err error) bool) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.circuitIsSuccessfulFn = fn

		return nil
	}
}

func WithOnAfterResponseFuncMiddleware(onAfterResponseFunc ...resty.ResponseMiddleware) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.onAfterResponseFunc = onAfterResponseFunc

		return nil
	}
}

func WithOnErrorFuncMiddleware(onErrorFunc ...resty.ErrorHook) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.onErrorFunc = onErrorFunc

		return nil
	}
}

func Options(opts ...Option) []Option {
	options := []Option{}
	options = append(options, opts...)

	return options
}

func NewGoHttpBreaker(opts ...Option) (*GoHttpBreaker, error) {
	ghb := &GoHttpBreaker{}
	ghb.httpCodesToIgnore = []int{404}
	ghb.statusCheckFn = ghb.defaultStatusCheck

	// call Option functions on instance to set options on it
	for _, opt := range opts {
		err := opt(ghb)
		if err != nil {
			return nil, err
		}
	}

	if ghb.config == nil {
		return nil, ErrConfigNotSet
	}

	withCircuitBreaker(ghb)
	withHttpClient(ghb)

	return ghb, nil
}

func withCircuitBreaker(g *GoHttpBreaker) {
	config := g.config

	readyToTripFunc := func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)

		return counts.Requests >= config.CircuitBreaker.MinRequestsToOpen && failureRatio >= config.CircuitBreaker.FailureAllowedRatio
	}

	onStateChangeFunc := func(name string, from gobreaker.State, to gobreaker.State) {
		if to == gobreaker.StateOpen {
			message := fmt.Sprintf("[CIRCUIT BREAKER] The circuit %s State Change from %s to %s", name, from, to)
			g.logger.Warn().Msg(message)

		} else {
			message := fmt.Sprintf("[CIRCUIT BREAKER] The circuit %s State Change from %s to %s", name, from, to)
			g.logger.Info().Msg(message)
		}
	}

	circuitSettings := gobreaker.Settings{
		Name:          config.Name,
		MaxRequests:   config.CircuitBreaker.RequestsInOpenState,
		Interval:      time.Duration(config.CircuitBreaker.IntervalMillis) * time.Millisecond,
		Timeout:       time.Duration(config.CircuitBreaker.OpenStateDurationMillis) * time.Millisecond,
		ReadyToTrip:   readyToTripFunc,
		OnStateChange: onStateChangeFunc,
	}

	if g.circuitIsSuccessfulFn != nil {
		circuitSettings.IsSuccessful = g.circuitIsSuccessfulFn
	}

	g.circuitBreaker = gobreaker.NewCircuitBreaker(circuitSettings)

	g.circuitBreakerPrometheus = newCircuitBreakerPrometheus(config.Name)
	_ = prometheus.Register(g.circuitBreakerPrometheus.cbCounterVec)

	c := newPrometheusStatsCollector(g.config.Name, g.circuitBreaker)
	_ = prometheus.Register(c)

}

func withHttpClient(g *GoHttpBreaker) {

	config := g.config.Client

	pm := ClientPromMiddleware{
		name:             g.config.Name,
		labels:           g.prometheusCustomLabels,
		histogramBuckets: g.prometheusHistogramBuckets,
		transformUrlFn:   g.prometheusTransformUrlFn,
	}

	pmMiddlewareFn := pm.MiddlewareFunc()
	g.httpClient = resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(time.Duration(config.TimeOutMillis) * time.Millisecond).
		SetRetryWaitTime(time.Duration(config.WaitTimeMillis) * time.Millisecond).
		SetRetryMaxWaitTime(time.Duration(config.MaxWaitTimeMillis) * time.Millisecond).
		SetRetryCount(config.MaxRetries).
		OnAfterResponse(pmMiddlewareFn).
		// timeouts would have been skipped from metrics otherwise
		OnError(func(request *resty.Request, err error) {
			var resErr *resty.ResponseError
			if errors.As(err, &resErr) {
				midErr := pmMiddlewareFn(nil, resErr.Response)
				if midErr != nil {
					g.logger.Error().Err(midErr).Msg("error when calling Prometheus on error function")
				}
			}
		}).
		SetHeaders(g.clientLevelHeaders)

	if g.retryCondition != nil {
		g.httpClient.AddRetryCondition(g.retryCondition)
	}

	for _, m := range g.onBeforeRequest {
		g.httpClient.OnBeforeRequest(m)
	}

	for i := range g.onAfterResponseFunc {
		g.httpClient.OnAfterResponse(g.onAfterResponseFunc[i])
	}

	for i := range g.onErrorFunc {
		g.httpClient.OnError(g.onErrorFunc[i])
	}

	// MaxIdleConnsPerHost is too low considering DefaultMaxIdleConns is 100. When 0 default is 2.
	transportToBeUsed := http.DefaultTransport
	if transport, ok := transportToBeUsed.(*http.Transport); ok {
		httpTransport := transport.Clone()
		httpTransport.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost
		if config.MaxIdleConns > 0 {
			httpTransport.MaxIdleConns = config.MaxIdleConns
		}
		transportToBeUsed = httpTransport
	}

	if g.transport != nil {
		transportToBeUsed = g.transport
	}
	g.httpClient.SetTransport(transportToBeUsed)

	if g.tracer != nil {
		g.httpClient.SetTransport(otelhttp.NewTransport(transportToBeUsed))
	}
	if !g.withoutDatadogIntegration {

		datadogOpts := g.datadogTransportOptions
		if len(datadogOpts) == 0 {
			datadogOpts = []httpdatadogtrace.RoundTripperOption{
				httpdatadogtrace.RTWithResourceNamer(func(req *http.Request) string {
					url := req.Context().Value(requestUrlContextKey)
					if url, ok := url.(string); ok {
						return url
					}

					return req.URL.Path
				}),
			}
		}
		httpdatadogtrace.WrapClient(
			g.httpClient.GetClient(),
			datadogOpts...,
		)
	}

	if g.tlsConfig != nil {
		g.httpClient.SetTLSClientConfig(g.tlsConfig)
	}

}

func WithTLSClientConfig(config *tls.Config) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.tlsConfig = config

		return nil
	}
}

func WithOnBeforeRequest(middleware resty.RequestMiddleware) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.onBeforeRequest = append(ghb.onBeforeRequest, middleware)

		return nil
	}
}

func WithoutDatadogIntegration(_ ...httpdatadogtrace.RoundTripperOption) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.withoutDatadogIntegration = true

		return nil
	}
}

func WithOverrideDatadogTransportOptions(opts ...httpdatadogtrace.RoundTripperOption) Option {
	return func(ghb *GoHttpBreaker) error {
		ghb.datadogTransportOptions = opts

		return nil
	}
}
