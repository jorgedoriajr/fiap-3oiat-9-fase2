package middleware

import (
	"github.com/rs/zerolog"
	"hamburgueria/pkg/logger"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func GetLogCallsMiddlewareFunc() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			l := logger.Get()

			ctx := c.Request().Context()
			l.WithContext(ctx)
			headers := c.Request().Header

			l.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("X-B3-TraceId", strings.Join(headers["X-B3-TraceId"], ";"))
			})
			l.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("X-B3-SpanId", strings.Join(headers["X-B3-SpanId"], ";"))
			})

			l.Info().
				Str("method", c.Request().Method).
				Str("url", c.Request().RequestURI).
				Str("user_agent", c.Request().UserAgent()).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")

			return next(c)
		}
	}
}
