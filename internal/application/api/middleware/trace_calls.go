package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opencensus.io/trace"
	"hamburgueria/pkg/starter"
)

func GetTraceCallsMiddlewareFunc() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, parent := trace.StartSpan(c.Request().Context(), starter.GetAppConfig().Name)

			c.Request().Header.Set("X-B3-TraceId", parent.SpanContext().TraceID.String())
			c.Request().Header.Set("X-B3-SpanId", parent.SpanContext().SpanID.String())

			return next(c)
		}
	}
}
