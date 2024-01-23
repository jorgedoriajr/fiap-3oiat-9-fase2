package httpclient

import (
	"fmt"
	"hamburgueria/pkg/starter"
	"strings"
	"sync"
)

const defaultClientName = "_default"

var (
	clients        map[string]Client
	initializeOnce sync.Once
)

func Initialize() {
	initializeOnce.Do(func() {
		clients = make(map[string]Client)
		conf := starter.GetHttpClientsConfig()

		var defaultConf Config
		if c, exists := conf[defaultClientName]; exists {
			defaultConf = Config{
				BaseUrl:                      c.BaseUrl,
				TimeOutMilliseconds:          c.TimeOutMilliseconds,
				MaxRetries:                   c.MaxRetries,
				RetryWaitTimeMilliseconds:    c.RetryWaitTimeMilliseconds,
				RetryMaxWaitTimeMilliseconds: c.RetryMaxWaitTimeMilliseconds,
				DefaultHeaders:               c.DefaultHeaders,
				CircuitBreaker: CircuitBreakerConfig{
					Enabled:                       c.CircuitBreaker.Enabled,
					RequestsInOpenState:           c.CircuitBreaker.RequestsInOpenState,
					IntervalMilliseconds:          c.CircuitBreaker.IntervalMilliseconds,
					OpenStateDurationMilliseconds: c.CircuitBreaker.OpenStateDurationMilliseconds,
					MinRequestsToOpen:             c.CircuitBreaker.MinRequestsToOpen,
					FailureAllowedRatio:           c.CircuitBreaker.FailureAllowedRatio,
				},
			}
		} else {
			defaultConf = Config{}
		}

		for k, v := range conf {
			if k == defaultClientName {
				continue
			}

			mergedConfig := mergeConfig(defaultConf, Config{
				BaseUrl:                      v.BaseUrl,
				TimeOutMilliseconds:          v.TimeOutMilliseconds,
				MaxRetries:                   v.MaxRetries,
				RetryWaitTimeMilliseconds:    v.RetryWaitTimeMilliseconds,
				RetryMaxWaitTimeMilliseconds: v.RetryMaxWaitTimeMilliseconds,
				DefaultHeaders:               v.DefaultHeaders,
				CircuitBreaker: CircuitBreakerConfig{
					Enabled:                       v.CircuitBreaker.Enabled,
					RequestsInOpenState:           v.CircuitBreaker.RequestsInOpenState,
					IntervalMilliseconds:          v.CircuitBreaker.IntervalMilliseconds,
					OpenStateDurationMilliseconds: v.CircuitBreaker.OpenStateDurationMilliseconds,
					MinRequestsToOpen:             v.CircuitBreaker.MinRequestsToOpen,
					FailureAllowedRatio:           v.CircuitBreaker.FailureAllowedRatio,
				},
			})

			c, err := NewClient(k, mergedConfig)
			if err != nil {
				panic(err)
			}
			clients[strings.ToLower(k)] = c
		}
	})
}

func GetClient(name string) Client {
	if c, ok := clients[strings.ToLower(name)]; ok {
		return c
	}

	panic(fmt.Sprintf("no http client found for %q", name))
}
