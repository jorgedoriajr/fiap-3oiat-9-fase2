package healthcheck

import (
	"context"
	"hamburgueria/internal/util/slice"
)

// HealthStatus represents a resource health status
type HealthStatus string

const (
	// StatusUp indicates a healthy resource
	StatusUp HealthStatus = "UP"
	// StatusDown indicates a unhealthy resource
	StatusDown HealthStatus = "DOWN"
	// StatusUnknown indicates a resource in a unknown state
	StatusUnknown HealthStatus = "UNKNOWN"
)

// Result represents a resource check output
type Result struct {
	// Status (optional) resource status
	Status HealthStatus `json:"status"`
	// Service (optional) service description. Used to identify a resource in a resource list.
	Service string `json:"service,omitempty"`
	// Message (optional) message to help understand the result of check
	Message string `json:"message,omitempty"`
}

// HealthChecker provides methods to check resource availability
type HealthChecker interface {
	// Check checks if a resource is healthy
	Check(context.Context) []Result
	// GetName name of resource being checked
	GetName() string
	// IsCritical indicates when a resource is critical to the application.
	// If true and healthcheck fails, application should be marked as unhealthy.
	// Otherwise, application should be marked as healthy regardless healthcheck result
	IsCritical() bool
}

// Check performs a health check based on given health checkers
func Check(checkers ...HealthChecker) (HealthStatus, []Result) {
	result := make([]Result, 0)
	resultStatus := StatusUp

	// TODO: goroutine
	// TODO: timeout
	// TODO: Unknown state
	for _, checker := range checkers {
		checkedResult := checker.Check(context.TODO())

		if slice.Any(checkedResult, func(r Result) bool {
			return r.Status != StatusUp && checker.IsCritical()
		}) {
			// TODO: log a non critical failed check
			resultStatus = StatusDown
		}

		result = append(result, checkedResult...)
	}

	return resultStatus, result
}
