package httpserver

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/pkg/healthcheck"
	"net/http"
)

func livenessCheck(checkers ...healthcheck.HealthChecker) func(c echo.Context) error {
	return func(c echo.Context) error {
		resultStatus, _ := healthcheck.Check(checkers...)

		return c.JSON(statusCode(resultStatus), result(resultStatus))
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

func result(r healthcheck.HealthStatus) healthcheck.Result {
	if r == healthcheck.StatusUp {
		return healthcheck.Result{Status: "UP"}
	}

	return healthcheck.Result{Status: "DOWN"}
}
