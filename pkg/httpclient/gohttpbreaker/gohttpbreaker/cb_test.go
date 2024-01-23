package gohttpbreaker //nolint:testpackage

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/sony/gobreaker"
	"github.com/stretchr/testify/assert"
)

func TestClosedCircuitWhenDefaultConfig(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, _ := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)

	circuitOpen := false

	for i := 0; i < 1000; i++ {
		req := client.R()
		_, err := client.Get(req, "test")

		if errors.Is(err, gobreaker.ErrOpenState) {
			circuitOpen = true

			break
		}
	}
	assert.False(t, circuitOpen, "circuit not opened")
}

func TestOpenCircuitWhenDefaultConfig(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, _ := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)

	circuitOpen := false

	for i := 0; i < 1000; i++ {
		req := client.R()
		_, err := client.Get(req, "test")

		if errors.Is(err, gobreaker.ErrOpenState) {
			circuitOpen = true

			break
		}
	}
	assert.True(t, circuitOpen, "circuit not opened")
}

func TestClosedCircuitWhenOverride(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, _ := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithOverrideHttpCodesToIgnore(400),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)

	circuitOpen := false

	for i := 0; i < 1000; i++ {
		req := client.R()
		_, err := client.Get(req, "test")

		if errors.Is(err, gobreaker.ErrOpenState) {
			circuitOpen = true

			break
		}
	}
	assert.False(t, circuitOpen, "circuit not opened")
}

func TestNewGoHttpBreakerWithCustomPrometheusLabelsButNoLabelValues(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, err = client.Get(client.R(), "test")
	assert.NoError(t, err)
}

func TestNewGoHttpBreakerWithCustomPrometheusLabelsButMissingLabelValues(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := client.R()
	req.SetContext(ContextWithCustomPrometheusLabelValues(req.Context(), "val1"))

	_, err = client.Get(req, "test")
	assert.NoError(t, err)
}

func TestNewGoHttpBreakerWithCustomPrometheusLabelsButExceedingLabelValues(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := client.R()
	req.SetContext(ContextWithCustomPrometheusLabelValues(req.Context(), "val1", "val2", "val3"))

	_, err = client.Get(req, "test")
	assert.NoError(t, err)
}

func TestNewGoHttpBreakerWithCustomPrometheusLabelsAndCorrectLabelValues(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithCustomPrometheusLabels("label1", "label2"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := client.R()
	req.SetContext(ContextWithCustomPrometheusLabelValues(req.Context(), "val1", "val2"))

	_, err = client.Get(req, "test")

	assert.NoError(t, err)
}

func TestWithPrometheusUrlTransformFn(t *testing.T) {
	expectedTransformation := "yes"
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = "localhost"
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithPrometheusUrlTransformFn(func(s string) string {
			return expectedTransformation
		}),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.NotNil(t, client.prometheusTransformUrlFn)
	assert.Equal(t, expectedTransformation, client.prometheusTransformUrlFn("no"))
}

func TestNewGoHttpBreaker_WithClientLevelHeader(t *testing.T) {
	testHeaderKey := "test-on-before-header-key"
	testHeaderValue := "test-on-before-header-value"
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, testHeaderValue, r.Header.Get(testHeaderKey))
		_, _ = w.Write([]byte("ok"))
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithClientLevelHeader(testHeaderKey, testHeaderValue),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreaker_WithBeforeRequest(t *testing.T) {
	testHeaderKey := "test-on-before-header-key"
	testHeaderValue := "test-on-before-header-value"
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, testHeaderValue, r.Header.Get(testHeaderKey))
		_, _ = w.Write([]byte("ok"))
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithOnBeforeRequest(func(client *resty.Client, r *resty.Request) error {
			r.SetHeader(testHeaderKey, testHeaderValue)

			return nil
		}),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreaker_WithoutRequesterToken(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.Header.Get("x-requester-token"))
		_, _ = w.Write([]byte("ok"))
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreaker_WithRequesterToken(t *testing.T) {
	requesterToken := "secret"
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, requesterToken, r.Header.Get("x-requester-token"))
		_, _ = w.Write([]byte("ok"))
	}))
	t.Setenv("REQUESTER_TOKEN", requesterToken)
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreakerWithRetryCondition(t *testing.T) {
	attemptCount := 0
	maxRetries := 3
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	defaultClientConfig.Client.MaxRetries = maxRetries
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
		WithRetryCondition(func(response *resty.Response, err error) bool {
			return response.StatusCode() == 500
		}),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	_, err = client.Get(client.R(), "test")
	assert.Error(t, err)
	assert.Equal(t, maxRetries+1, attemptCount)
}

func TestNewGoHttpBreakerWithoutRetryCondition(t *testing.T) {
	attemptCount := 0
	maxRetries := 3
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000
	defaultClientConfig.Client.MaxRetries = maxRetries
	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithRequesterTokenFromEnv(),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	_, err = client.Get(client.R(), "test")
	assert.Error(t, err)
	assert.Equal(t, 1, attemptCount)
}

func TestNewGoHttpBreaker_WithTransport(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 10
	transport.MaxConnsPerHost = 10
	transport.MaxIdleConnsPerHost = 10

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithTransport(transport),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreaker_WithHttpCodesToIgnore(t *testing.T) {
	setupTestClient := func(writeResponse string, statusResponse int, httpCodesToIgnore ...int) (*GoHttpBreaker, error) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusResponse)
			_, _ = w.Write([]byte(writeResponse))
		}))

		defaultClientConfig := DefaultConfig()
		defaultClientConfig.Name = "unit-test"
		defaultClientConfig.Client.BaseURL = testServer.URL
		defaultClientConfig.Client.TimeOutMillis = 5000

		client, err := NewGoHttpBreaker(
			WithConfig(defaultClientConfig),
			WithOverrideHttpCodesToIgnore(httpCodesToIgnore...),
		)

		return client, err
	}
	t.Run("SuccessfulRequest", func(t *testing.T) {
		client, err := setupTestClient("ok", 200)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "ok", string(resp.Body()))
	})

	t.Run("ShouldIgnoreCodeWhenValueIsConfigured", func(t *testing.T) {
		client, err := setupTestClient("bad request", 400, 400)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "bad request", string(resp.Body()))
	})

	t.Run("ShouldNotIgnoreCodeWhenDifferentValueIsConfigured", func(t *testing.T) {
		client, err := setupTestClient("unauthorized", 401, 400)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.Error(t, err)
		assert.Equal(t, "unauthorized", string(resp.Body()))
	})

	t.Run("ShouldNotIgnoreCodeWhenNoValueIsConfigured", func(t *testing.T) {
		client, err := setupTestClient("bad request", 400)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.Error(t, err)
		assert.Equal(t, "bad request", string(resp.Body()))
	})

	t.Run("ShouldIgnoreCodeWhenNoValueIsConfiguredButErrorIs404", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("not found"))
		}))

		defaultClientConfig := DefaultConfig()
		defaultClientConfig.Name = "unit-test"
		defaultClientConfig.Client.BaseURL = testServer.URL
		defaultClientConfig.Client.TimeOutMillis = 5000

		client, _ := NewGoHttpBreaker(
			WithConfig(defaultClientConfig),
		)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "not found", string(resp.Body()))
	})
}

func TestNewGoHttpBreaker_WithCustomStatusCheck(t *testing.T) {
	setupTestClient := func(
		writeResponse string,
		statusResponse int,
		statusCheckFn func(statusCode int) bool,
	) (*GoHttpBreaker, error) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusResponse)
			_, _ = w.Write([]byte(writeResponse))
		}))

		defaultClientConfig := DefaultConfig()
		defaultClientConfig.Name = "unit-test"
		defaultClientConfig.Client.BaseURL = testServer.URL
		defaultClientConfig.Client.TimeOutMillis = 5000

		client, err := NewGoHttpBreaker(
			WithConfig(defaultClientConfig),
			WithCustomStatusCheck(statusCheckFn),
		)

		return client, err
	}

	t.Run("SuccessfulRequest", func(t *testing.T) {
		client, err := setupTestClient("ok", 200, func(statusCode int) bool {
			return false
		})
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "ok", string(resp.Body()))
	})

	match4xx := func(statusCode int) bool {
		return statusCode > 399 && statusCode < 500
	}

	t.Run("ShouldFailOn4xxMatcherFn", func(t *testing.T) {
		client, err := setupTestClient("bad request", 400, match4xx)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.Error(t, err)
		assert.Equal(t, "bad request", string(resp.Body()))
	})

	t.Run("ShouldIgnoreOn4xxMatcherFn", func(t *testing.T) {
		client, err := setupTestClient("internal server error", 500, match4xx)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "internal server error", string(resp.Body()))
	})

	matchAllErrStatuses := func(statusCode int) bool {
		return statusCode >= http.StatusBadRequest
	}

	t.Run("ShouldFailOnAllErrMatcherFn", func(t *testing.T) {
		client, err := setupTestClient("bad gateway", 502, matchAllErrStatuses)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.Error(t, err)
		assert.Equal(t, "bad gateway", string(resp.Body()))
	})

	t.Run("ShouldIgnoreOnAllErrMatcherFn", func(t *testing.T) {
		client, err := setupTestClient("temporary redirect", 307, matchAllErrStatuses)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		resp, err := client.Get(client.R(), "test")
		assert.NoError(t, err)
		assert.Equal(t, "temporary redirect", string(resp.Body()))
	})
}

func TestNewGoHttpBreaker_CustomStatusCheckPrecedence(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("unauthorized"))
	}))

	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithOverrideHttpCodesToIgnore(401),
		WithCustomStatusCheck(func(statusCode int) bool {
			return statusCode == 401
		}),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.Get(client.R(), "test")
	assert.Error(t, err)
	assert.Equal(t, "unauthorized", string(resp.Body()))
}

func TestNewGoHttpBreaker_WithTLS(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defaultClientConfig := DefaultConfig()
	defaultClientConfig.Name = "unit-test"
	defaultClientConfig.Client.BaseURL = testServer.URL
	defaultClientConfig.Client.TimeOutMillis = 5000

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	client, err := NewGoHttpBreaker(
		WithConfig(defaultClientConfig),
		WithTLSClientConfig(config),
	)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.Get(client.R(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "ok", string(resp.Body()))
}

func TestNewGoHttpBreaker_CustomCircuitIsSuccessful(t *testing.T) {
	setupTestClient := func(customFn func(err error) bool) (*GoHttpBreaker, error) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("success"))
		}))

		defaultClientConfig := DefaultConfig()
		defaultClientConfig.Name = "unit-test"
		defaultClientConfig.Client.BaseURL = testServer.URL
		defaultClientConfig.Client.TimeOutMillis = 5000

		return NewGoHttpBreaker(
			WithConfig(defaultClientConfig),
			WithCustomCircuitIsSuccessful(customFn),
		)
	}

	t.Run("ShouldNotIncreaseFailuresOnContextCanceled", func(t *testing.T) {
		client, err := setupTestClient(func(err error) bool {
			return err == nil || errors.Is(err, context.Canceled)
		})
		assert.NoError(t, err)
		assert.NotNil(t, client)

		// a cancelled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		resp, err := client.Get(client.RWithCtx(ctx), "test")
		assert.Error(t, context.Canceled, err)
		assert.Empty(t, resp.Body())
		assert.Zero(t, client.circuitBreaker.Counts().TotalFailures)
	})

	t.Run("ShouldIncreaseFailuresOnContextCanceled", func(t *testing.T) {
		client, err := setupTestClient(func(err error) bool {
			return err == nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, client)

		// a cancelled context
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		resp, err := client.Get(client.RWithCtx(ctx), "test")
		assert.Error(t, context.Canceled, err)
		assert.Empty(t, resp.Body())
		assert.Equal(t, uint32(1), client.circuitBreaker.Counts().TotalFailures)
	})
}
