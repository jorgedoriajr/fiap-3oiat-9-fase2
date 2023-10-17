package httpserver

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/internal/util/healthcheck"
	"net/http"
)

func livenessCheck() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, healthcheck.Result{Status: "UP"})
	}
}

func readinessCheck(checkers ...healthcheck.HealthChecker) func(c echo.Context) error {

	return func(c echo.Context) error {
		resultStatus, result := healthcheck.Check(checkers...)
		return c.JSON(statusCode(resultStatus), result)
	}
}

func statusCode(r healthcheck.HealthStatus) int {
	if r == healthcheck.StatusUp {
		return http.StatusOK
	}

	return http.StatusServiceUnavailable
}
