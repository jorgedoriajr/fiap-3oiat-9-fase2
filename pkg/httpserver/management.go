package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) startManagementServer() {
	if !s.isManagementEnabled() {
		return
	}

	go func() {
		managementPort := fmt.Sprintf(":%d", s.builder.serverConfig.Management.Port)

		if err := s.managementInstance.Start(managementPort); err != http.ErrServerClosed {
			logging.GetLogger().Fatal(context.Background(), "[Server] Error while starting the management server", err)
			panic(err)
		}
	}()
}

func (s *server) configureManagementServer() {

	if !s.isManagementEnabled() {
		return
	}

	s.managementInstance = echo.New()
	s.managementInstance.HideBanner = true

	if s.builder.serverConfig.Management.HealthCheck.Enabled {
		s.managementInstance.GET(
			s.builder.serverConfig.Management.HealthCheck.LivenessPath,
			livenessCheck(),
		)

		s.managementInstance.GET(
			s.builder.serverConfig.Management.HealthCheck.ReadinessPath,
			readinessCheck(s.builder.healthCheckers...),
		)
	}

	if s.builder.serverConfig.Management.Metrics.Enabled {
		c := metrics.MetricsMiddlewareConfig{
			SkippedPaths: []string{
				s.builder.serverConfig.Management.HealthCheck.LivenessPath,
				s.builder.serverConfig.Management.HealthCheck.ReadinessPath,
				s.builder.serverConfig.Management.Metrics.Path,
			},
			MetricsPath: s.builder.serverConfig.Management.Metrics.Path,
			Subsystem:   s.builder.serverConfig.Management.Metrics.Subsystem,
		}

		metricsMd := metrics.GetMetricsMiddleware(c)
		s.echoInstance.Use(metricsMd.HandlerFunc)
		metricsMd.SetMetricsPath(s.managementInstance)
	}
}

func (s *server) isManagementEnabled() bool {
	return s.builder.serverConfig.Management.HealthCheck.Enabled ||
		s.builder.serverConfig.Management.Metrics.Enabled
}
