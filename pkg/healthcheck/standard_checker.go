package healthcheck

import "context"

// StandardChecker implements [HealthChecker] interface
type StandardChecker struct {
	Name     string
	Critical bool
	Checker  func(context.Context) []Result
}

// Check implements [HealthChecker.Check(context.Context)]
func (c *StandardChecker) Check(ctx context.Context) []Result {
	return c.Checker(ctx)
}

// GetName implements [HealthChecker.GetName()]
func (c *StandardChecker) GetName() string {
	return c.Name
}

// IsCritical implements [HealthChecker.IsCritical()]
func (c *StandardChecker) IsCritical() bool {
	return c.Critical
}
