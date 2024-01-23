package httpclient

import "fmt"

type EmptyResponse struct{}

type Response[T any] struct {
	Result     *T
	StatusCode int
	Headers    map[string]string

	body []byte
}

func (r *Response[T]) Body() []byte {
	return r.body
}

type RequestError struct {
	StatusCode int
	Message    string
}

func (e RequestError) Error() string {
	return fmt.Sprintf("HTTP request failed. status: %d. message: %q", e.StatusCode, e.Message)
}
