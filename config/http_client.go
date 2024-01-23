package config

type HttpClientConfig struct {
	BaseUrl                      string
	TimeOutMilliseconds          int
	MaxRetries                   int
	RetryWaitTimeMilliseconds    int
	RetryMaxWaitTimeMilliseconds int
	DefaultHeaders               map[string]string
	CircuitBreaker               CircuitBreakerConfig
}

type CircuitBreakerConfig struct {
	Enabled                       bool
	RequestsInOpenState           uint32
	IntervalMilliseconds          uint32
	OpenStateDurationMilliseconds uint32
	MinRequestsToOpen             uint32
	FailureAllowedRatio           float64
}
