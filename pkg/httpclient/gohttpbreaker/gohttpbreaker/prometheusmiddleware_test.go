package gohttpbreaker //nolint:testpackage

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	client := resty.New()
	resp := &resty.Response{
		RawResponse: &http.Response{
			Body:       nil,
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		Request: &resty.Request{
			URL:    "http://localhost:8080/storage/d99de142-a1d2-434b-8240-b79dcce55b85/000000045/modified/85396b87-d830-4ea2-8c94-02f3676d0685",
			Method: "GET",
		},
	}
	cpm := ClientPromMiddleware{
		name: "unit-test",
	}
	middlewareFn := cpm.MiddlewareFunc()
	cpm.reqs.Reset()
	assert.NoError(t, middlewareFn(client, resp))
	assert.Equal(t, float64(1), testutil.ToFloat64(cpm.reqs.WithLabelValues("200", "GET", "/storage/:id/000000045/modified/:id")))
}

func TestMiddlewareWithContext(t *testing.T) {
	ctx := context.Background()
	client := resty.New()
	req := &resty.Request{
		URL:    "http://localhost:8080/storage/d99de142-a1d2-434b-8240-b79dcce55b85/000000045/modified/85396b87-d830-4ea2-8c94-02f3676d0685",
		Method: "GET",
	}
	req.SetContext(context.WithValue(ctx, requestUrlContextKey, "/storage/{id0}/{id1}/modified/{id2}"))
	resp := &resty.Response{
		RawResponse: &http.Response{
			Body:       nil,
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		Request: req,
	}
	cpm := ClientPromMiddleware{
		name: "unit-test",
	}
	middlewareFn := cpm.MiddlewareFunc()
	cpm.reqs.Reset()
	assert.NoError(t, middlewareFn(client, resp))
	assert.Equal(t, float64(1), testutil.ToFloat64(cpm.reqs.WithLabelValues("200", "GET", "/storage/{id0}/{id1}/modified/{id2}")))
}

func TestMiddleware_WithUrlTransformer(t *testing.T) {
	client := resty.New()
	resp1 := &resty.Response{
		RawResponse: &http.Response{
			Body:       nil,
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		Request: &resty.Request{
			URL:    "http://localhost:8080/storage/d99de142-a1d2-434b-8240-b79dcce55b85/000000045/modified/85396b87-d830-4ea2-8c94-02f3676d0685",
			Method: "GET",
		},
	}
	resp2 := &resty.Response{
		RawResponse: &http.Response{
			Body:       nil,
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		Request: &resty.Request{
			URL:    "http://localhost:8080/merchant/85396b87-d830-4ea2-8c94-02f3676d0685",
			Method: "GET",
		},
	}
	cpm := ClientPromMiddleware{
		name: "unit-test",
		transformUrlFn: func(url string) string {
			if strings.HasPrefix(url, "http://localhost:8080/storage") {
				return "/storage"
			}

			return ApplyDefaultUrlTransformation(url)
		},
	}
	middlewareFn := cpm.MiddlewareFunc()
	cpm.reqs.Reset()
	assert.NoError(t, middlewareFn(client, resp1))
	assert.NoError(t, middlewareFn(client, resp2))
	assert.Equal(t, float64(1), testutil.ToFloat64(cpm.reqs.WithLabelValues("200", "GET", "/storage")))
	assert.Equal(t, float64(1), testutil.ToFloat64(cpm.reqs.WithLabelValues("200", "GET", "/merchant/:id")))
}
