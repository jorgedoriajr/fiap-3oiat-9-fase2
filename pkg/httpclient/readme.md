`auto generated documentation`

# http
--
    import "."


## Usage

#### func  NewRequest

```go
func NewRequest[T any](ctx context.Context, client Client, path string) *request[T]
```

#### type CircuitBreakerConfig

```go
type CircuitBreakerConfig struct {
	Enabled                       bool
	RequestsInOpenState           uint32
	IntervalMilliseconds          uint32
	OpenStateDurationMilliseconds uint32
	MinRequestsToOpen             uint32
	FailureAllowedRatio           float64
}
```

CircuitBreakerConfig represents http circuit breaker configuration

#### type Client

```go
type Client interface {
	Send(ctx context.Context, method string, path string, body any, headers map[string]string, pathParams map[string]string, queryParams map[string]string) (*resty.Response, error)
}
```


#### func  NewClient

```go
func NewClient(name string, conf Config) (Client, error)
```

#### type Config

```go
type Config struct {
	BaseUrl                      string
	TimeOutMilliseconds          int
	MaxRetries                   int
	RetryWaitTimeMilliseconds    int
	RetryMaxWaitTimeMilliseconds int
	DefaultHeaders               map[string]string
	CircuitBreaker               CircuitBreakerConfig
}
```

Config represents http client configuration

#### type EmptyResponse

```go
type EmptyResponse struct{}
```


#### type RequestError

```go
type RequestError struct {
	StatusCode int
	Message    string
}
```


#### func (RequestError) Error

```go
func (e RequestError) Error() string
```

#### type Response

```go
type Response[T any] struct {
	Result     *T
	StatusCode int
	Headers    map[string]string
}
```


#### func (*Response[T]) Body

```go
func (r *Response[T]) Body() []byte
```
