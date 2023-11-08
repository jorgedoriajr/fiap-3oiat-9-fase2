package middleware

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/pkg/starter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockHandlerFunction(c echo.Context) error {
	return nil
}

func TestTraceCallsMiddleware(t *testing.T) {
	t.Run(`should log calls`, func(t *testing.T) {
		starter.Initialize()
		traceCallsMiddlewareFunction := GetTraceCallsMiddlewareFunc()
		traceCallsMiddlewareHandlerFunction := traceCallsMiddlewareFunction(mockHandlerFunction)

		echoServer := echo.New()
		testUrl := "http://localhost:8080/customers/10609328220"
		bodyContent := `{"token": "NjIxMTUwMTIzNDU2Nzg5MTAxMTU1MTMzMg=="}`
		body := strings.NewReader(bodyContent)
		req, _ := http.NewRequest(http.MethodPost, testUrl, body)
		rec := httptest.NewRecorder()
		echoContext := echoServer.NewContext(req, rec)
		_ = traceCallsMiddlewareHandlerFunction(echoContext)

		traceHeader := echoContext.Request().Header.Get("X-B3-TraceId")
		spanHeader := echoContext.Request().Header.Get("X-B3-SpanId")

		assert.NotNil(t, traceHeader)
		assert.NotNil(t, spanHeader)
		assert.NotEqual(t, traceHeader, "")
		assert.NotEqual(t, spanHeader, "")
	})
}
