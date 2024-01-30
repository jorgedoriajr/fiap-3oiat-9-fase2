package httpclient

import (
	"context"
	"hamburgueria/pkg/httpclient/gohttpbreaker/gohttpbreaker"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	Send(ctx context.Context, method string, path string, body any, headers map[string]string, pathParams map[string]string, queryParams map[string]string) (*resty.Response, error)
}

type client struct {
	httpBreaker *gohttpbreaker.GoHttpBreaker
}

func NewClient(name string, conf Config) (Client, error) {

	clientConfig := getClientConfig(name, conf)

	opts := gohttpbreaker.Options(
		gohttpbreaker.WithConfig(clientConfig),
		gohttpbreaker.WithRequesterTokenFromEnv(),
	)

	if len(conf.DefaultHeaders) > 0 {
		for k, v := range conf.DefaultHeaders {
			opts = append(opts, gohttpbreaker.WithClientLevelHeader(k, v))
		}
	}

	c, err := gohttpbreaker.NewGoHttpBreaker(
		opts...,
	)

	if err != nil {
		return nil, err
	}

	return &client{
		httpBreaker: c,
	}, nil
}

func (c client) Send(ctx context.Context, method string, path string, body any, headers map[string]string, pathParams map[string]string, queryParams map[string]string) (*resty.Response, error) {

	req := c.httpBreaker.RWithCtx(ctx)

	if body != nil {
		req.SetBody(body)
	}

	req.SetHeaders(headers)
	req.SetPathParams(pathParams)
	req.SetQueryParams(queryParams)

	switch method {
	case http.MethodPost:
		return c.httpBreaker.Post(req, path)
	case http.MethodHead:
		return c.httpBreaker.Head(req, path)
	case http.MethodPut:
		return c.httpBreaker.Put(req, path)
	case http.MethodPatch:
		return c.httpBreaker.Patch(req, path)
	case http.MethodOptions:
		return c.httpBreaker.Options(req, path)
	case http.MethodDelete:
		return c.httpBreaker.Delete(req, path)
	default:
		return c.httpBreaker.Get(req, path)
	}
}

func getClientConfig(name string, conf Config) *gohttpbreaker.Config {

	c := &gohttpbreaker.Config{}
	c.Name = name
	c.Debug = false

	c.CircuitBreaker.Enabled = conf.CircuitBreaker.Enabled
	c.CircuitBreaker.RequestsInOpenState = conf.CircuitBreaker.RequestsInOpenState
	c.CircuitBreaker.IntervalMillis = conf.CircuitBreaker.IntervalMilliseconds
	c.CircuitBreaker.OpenStateDurationMillis = conf.CircuitBreaker.OpenStateDurationMilliseconds
	c.CircuitBreaker.MinRequestsToOpen = conf.CircuitBreaker.MinRequestsToOpen
	c.CircuitBreaker.FailureAllowedRatio = conf.CircuitBreaker.FailureAllowedRatio

	c.Client.BaseURL = conf.BaseUrl
	c.Client.TimeOutMillis = conf.TimeOutMilliseconds
	c.Client.MaxRetries = conf.MaxRetries
	c.Client.WaitTimeMillis = conf.RetryWaitTimeMilliseconds
	c.Client.MaxWaitTimeMillis = conf.RetryMaxWaitTimeMilliseconds

	return c
}
