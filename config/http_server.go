package config

type HttpServerConfig struct {
	Port                int
	UseRequestValidator bool
	Management          ManagementConfig
}

type ManagementConfig struct {
	Port        int
	HealthCheck HealthCheckConfig
	Metrics     ManagementMetricsConfig
}

type HealthCheckConfig struct {
	Enabled       bool
	LivenessPath  string
	ReadinessPath string
}

type ManagementMetricsConfig struct {
	Enabled   bool
	Path      string
	Subsystem string
}
