package httpclient

// Config represents http client configuration
type Config struct {
	BaseUrl                      string
	TimeOutMilliseconds          int
	MaxRetries                   int
	RetryWaitTimeMilliseconds    int
	RetryMaxWaitTimeMilliseconds int
	DefaultHeaders               map[string]string
	CircuitBreaker               CircuitBreakerConfig
}

// CircuitBreakerConfig represents http circuit breaker configuration
type CircuitBreakerConfig struct {
	Enabled                       bool
	RequestsInOpenState           uint32
	IntervalMilliseconds          uint32
	OpenStateDurationMilliseconds uint32
	MinRequestsToOpen             uint32
	FailureAllowedRatio           float64
}

func mergeConfig(defaultConf, conf Config) Config {
	c := defaultConf

	if conf.BaseUrl != "" {
		c.BaseUrl = conf.BaseUrl
	}
	if conf.TimeOutMilliseconds > 0 {
		c.TimeOutMilliseconds = conf.TimeOutMilliseconds
	}
	if conf.MaxRetries > 0 {
		c.MaxRetries = conf.MaxRetries
	}
	if conf.RetryWaitTimeMilliseconds > 0 {
		c.RetryWaitTimeMilliseconds = conf.RetryWaitTimeMilliseconds
	}
	if conf.RetryMaxWaitTimeMilliseconds > 0 {
		c.RetryMaxWaitTimeMilliseconds = conf.RetryMaxWaitTimeMilliseconds
	}
	if len(conf.DefaultHeaders) > 0 {
		c.DefaultHeaders = conf.DefaultHeaders
	}

	c.CircuitBreaker = mergeCircuitBreakerConfig(c.CircuitBreaker, conf.CircuitBreaker)
	return c
}

func mergeCircuitBreakerConfig(defaultConf, conf CircuitBreakerConfig) CircuitBreakerConfig {
	c := defaultConf
	c.Enabled = conf.Enabled

	if conf.RequestsInOpenState > 0 {
		c.RequestsInOpenState = conf.RequestsInOpenState
	}
	if conf.IntervalMilliseconds > 0 {
		c.IntervalMilliseconds = conf.IntervalMilliseconds
	}
	if conf.OpenStateDurationMilliseconds > 0 {
		c.OpenStateDurationMilliseconds = conf.OpenStateDurationMilliseconds
	}
	if conf.MinRequestsToOpen > 0 {
		c.MinRequestsToOpen = conf.MinRequestsToOpen
	}
	if conf.FailureAllowedRatio > 0 {
		c.FailureAllowedRatio = conf.FailureAllowedRatio
	}

	return c
}
