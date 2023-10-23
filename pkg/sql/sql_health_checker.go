package sql

import (
	"context"
	"fmt"
	"hamburgueria/pkg/healthcheck"
	"time"
)

type healthCheckOptions struct {
	isCritical bool
	name       string
}

/*
HealthCheckIsCritical overrides [healthcheck.StandardChecker.Critical] value

default: true
*/
func HealthCheckIsCritical(isCritical bool) Opt {
	ensureNotInitialized()
	return func(o *options) error {
		o.healthCheckOptions.isCritical = isCritical
		return nil
	}
}

/*
HealthCheckName overrides [healthcheck.StandardChecker.Name] value

default: SQL Database
*/
func HealthCheckName(name string) Opt {
	ensureNotInitialized()
	return func(o *options) error {
		o.healthCheckOptions.name = name
		return nil
	}
}

// GetHealthChecker returns standard database health checker
func GetHealthChecker() healthcheck.HealthChecker {
	ensureInitialized()
	return stdHealthChecker
}

func initHealthChecker() {
	stdHealthChecker.Name = opts.healthCheckOptions.name
	stdHealthChecker.Critical = opts.healthCheckOptions.isCritical
	stdHealthChecker.Checker = check
}

func check(ctx context.Context) []healthcheck.Result {
	result := make([]healthcheck.Result, 0)
	for name, client := range clients {

		err := func(c Client) error {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()
			return c.Ping(ctx)
		}(client)

		if err != nil {
			result = append(result, healthcheck.Result{
				Status:  healthcheck.StatusDown,
				Service: fmt.Sprintf("Database: %s", name),
				Message: err.Error(),
			})
		} else {
			result = append(result, healthcheck.Result{
				Status:  healthcheck.StatusUp,
				Service: fmt.Sprintf("Database: %s", name),
			})
		}
	}

	return result
}
