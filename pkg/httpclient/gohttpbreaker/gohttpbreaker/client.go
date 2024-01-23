package gohttpbreaker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"moul.io/http2curl/v2"
)

type contextKey string

const (
	requestUrlContextKey contextKey = "resty_request_url"
)

// Client method returns the current `resty.Client` used by the gohttpbreaker client.
func (ghb *GoHttpBreaker) Client() *resty.Client {
	return ghb.httpClient
}

// NetHttpClient method returns the current `http.Client` used by the resty client.
func (ghb *GoHttpBreaker) NetHttpClient() *http.Client {
	return ghb.Client().GetClient()
}

func (ghb *GoHttpBreaker) R() *resty.Request {
	return ghb.httpClient.R()
}

func (ghb *GoHttpBreaker) RWithCtx(ctx context.Context) *resty.Request {
	return ghb.httpClient.R().SetContext(ctx)
}

func (ghb *GoHttpBreaker) Get(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodGet, url)
}

func (ghb *GoHttpBreaker) Post(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodPost, url)
}

func (ghb *GoHttpBreaker) Head(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodHead, url)
}

func (ghb *GoHttpBreaker) Put(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodPut, url)
}

func (ghb *GoHttpBreaker) Patch(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodPatch, url)
}

func (ghb *GoHttpBreaker) Options(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodOptions, url)
}

func (ghb *GoHttpBreaker) Delete(request *resty.Request, url string) (*resty.Response, error) {
	return ghb.request(request, resty.MethodDelete, url)
}

// Execute performs http request on specified method
func (ghb *GoHttpBreaker) Execute(request *resty.Request, method string, url string) (*resty.Response, error) {
	return ghb.request(request, method, url)
}

func (ghb *GoHttpBreaker) request(request *resty.Request, method string, url string) (*resty.Response, error) {
	request.SetContext(context.WithValue(request.Context(), requestUrlContextKey, url))
	if !ghb.config.CircuitBreaker.Enabled {
		return request.Execute(method, url) //nolint:wrapcheck
	}

	res, cberr := ghb.circuitBreaker.Execute(func() (interface{}, error) {
		response, err := request.Execute(method, url)
		if err == nil && response.IsError() && ghb.statusCheckFn(response.StatusCode()) {
			err = fmt.Errorf("failed with status code %d. %s", response.StatusCode(), response.String()) //nolint:goerr113
		}

		return response, err
	})

	return ghb.processResponse(res, cberr)
}

func (ghb *GoHttpBreaker) processResponse(res interface{}, cberr error) (*resty.Response, error) {
	ghb.circuitBreakerPrometheus.RegisterEvent(cberr)
	if res == nil {
		return &resty.Response{}, cberr
	}

	restyResponse := res.(*resty.Response)

	if ghb.config.Debug {
		request := restyResponse.Request.RawRequest
		requestBody, err := json.Marshal(restyResponse.Request.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not process response")
		}
		request.Body = io.NopCloser(strings.NewReader(string(requestBody)))
		_, _ = http2curl.GetCurlCommand(request)
	}

	return restyResponse, cberr
}

// true means status is an error that should be accounted for
func (ghb *GoHttpBreaker) defaultStatusCheck(statusCode int) bool {
	for _, code := range ghb.httpCodesToIgnore {
		if statusCode == code {
			return false
		}
	}

	return true
}
