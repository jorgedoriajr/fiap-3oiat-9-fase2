package gohttpbreaker

// Config includes http client and circuit breaker configuration
//
// RequestsInOpenState is the maximum number of requests allowed to pass through
// when the CircuitBreaker is half-open.
// If RequestsInOpenState is 0, the CircuitBreaker allows only 1 request.
//
// IntervalMillis is the cyclic period of the closed state
// for the CircuitBreaker to clear the internal Counts.
// If IntervalMillis is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
//
// OpenStateDurationMillis is the period of the open state,
// after which the state of the CircuitBreaker becomes half-open.
// If OpenStateDurationMillis is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
//
// MinRequestsToOpen is the minimum numbers of requests required to open state.
//
// Failure allowed ratio is the threshold ratio (between 0 and 1) to open state.
//
// MaxRetries, if greater than 0. The client will retry the number of times.
//
// WaitTimeMillis method sets default wait time to sleep before retrying
// request.
//
// MaxWaitTimeMillis method sets max wait time to sleep before retrying
// request.
type Config struct {
	Name           string `mapstructure:"name" default:""`
	Debug          bool   `mapstructure:"debug" default:"false"`
	CircuitBreaker struct {
		Enabled                 bool    `mapstructure:"enabled" default:"true"`
		RequestsInOpenState     uint32  `mapstructure:"requests-in-open-state" default:"10"`
		IntervalMillis          uint32  `mapstructure:"interval-ms" default:"30000"`
		OpenStateDurationMillis uint32  `mapstructure:"open-state-duration-ms" default:"60000"`
		MinRequestsToOpen       uint32  `mapstructure:"min-requests-to-open" default:"30"`
		FailureAllowedRatio     float64 `mapstructure:"failure-allowed-ratio" default:"0.30"`
	} `mapstructure:"circuit-breaker"`
	Client struct {
		BaseURL             string `mapstructure:"base-url" default:""`
		TimeOutMillis       int    `mapstructure:"timeout-ms" default:"1000"`
		MaxRetries          int    `mapstructure:"max-retries" default:"0"`
		WaitTimeMillis      int    `mapstructure:"wait-time-ms" default:"100"`
		MaxWaitTimeMillis   int    `mapstructure:"max-wait-time-ms" default:"2000"`
		MaxIdleConnsPerHost int    `mapstructure:"max-idle-conns-per-host" default:"20"`
		MaxIdleConns        int    `mapstructure:"max-idle-conns" default:"100"`
	} `mapstructure:"client"`
}

func DefaultConfig() *Config {
	config := &Config{}

	config.Name = ""
	config.Debug = false

	config.CircuitBreaker.Enabled = true
	config.CircuitBreaker.RequestsInOpenState = 10
	config.CircuitBreaker.IntervalMillis = 30000
	config.CircuitBreaker.OpenStateDurationMillis = 60000
	config.CircuitBreaker.MinRequestsToOpen = 30
	config.CircuitBreaker.FailureAllowedRatio = 0.30

	config.Client.BaseURL = ""
	config.Client.TimeOutMillis = 1000
	config.Client.MaxRetries = 0
	config.Client.WaitTimeMillis = 100
	config.Client.MaxWaitTimeMillis = 2000
	config.Client.MaxIdleConns = 100
	config.Client.MaxIdleConnsPerHost = 20

	return config
}
