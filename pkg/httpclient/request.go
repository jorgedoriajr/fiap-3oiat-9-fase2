package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
)

type request[T any] struct {
	ctx         context.Context
	path        string
	client      Client
	headers     map[string]string
	pathParams  map[string]string
	queryParams map[string]string
}

func NewRequest[T any](ctx context.Context, client Client, path string) *request[T] {
	return &request[T]{
		ctx:         ctx,
		path:        path,
		client:      client,
		headers:     map[string]string{},
		pathParams:  map[string]string{},
		queryParams: map[string]string{},
	}
}

func (r *request[T]) WithHeaders(headers map[string]string) *request[T] {
	r.headers = headers
	return r
}

func (r *request[T]) WithPathParams(pathParams map[string]string) *request[T] {
	r.pathParams = pathParams
	return r
}

func (r *request[T]) WithQueryParams(queryParams map[string]string) *request[T] {
	r.queryParams = queryParams
	return r
}

func (r request[T]) Get() (*Response[T], error) {
	return r.send(http.MethodGet, nil)
}

func (r request[T]) Delete() (*Response[T], error) {
	return r.send(http.MethodDelete, nil)
}

func (r request[T]) Options() (*Response[T], error) {
	return r.send(http.MethodOptions, nil)
}

func (r request[T]) Head() (*Response[T], error) {
	return r.send(http.MethodHead, nil)
}

func (r request[T]) Post(body any) (*Response[T], error) {
	return r.send(http.MethodPost, body)
}

func (r request[T]) Put(body any) (*Response[T], error) {
	return r.send(http.MethodPut, body)
}

func (r request[T]) Patch(body any) (*Response[T], error) {
	return r.send(http.MethodPatch, body)
}

func (r request[T]) send(method string, body any) (*Response[T], error) {

	response, err := r.client.Send(r.ctx, method, r.path, body, r.headers, r.pathParams, r.queryParams)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccess() {
		return &Response[T]{
				Result:     nil,
				StatusCode: response.StatusCode(),
				Headers:    joinHeaderValues(response.Header()),
				body:       response.Body(),
			}, &RequestError{
				StatusCode: response.StatusCode(),
				Message:    string(response.Body()),
			}
	}

	var decodedResponse T
	if err := json.Unmarshal(response.Body(), &decodedResponse); err != nil {
		return nil, err
	}

	return &Response[T]{
		Result:     &decodedResponse,
		StatusCode: response.StatusCode(),
		Headers:    joinHeaderValues(response.Header()),
		body:       response.Body(),
	}, nil
}
