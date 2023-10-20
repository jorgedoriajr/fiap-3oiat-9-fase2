package httpserver

import (
	"errors"
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

		if err := s.managementInstance.Start(managementPort); !errors.Is(err, http.ErrServerClosed) {
			panic("[Server] Error while starting the management server")
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
}

func (s *server) isManagementEnabled() bool {
	return s.builder.serverConfig.Management.HealthCheck.Enabled ||
		s.builder.serverConfig.Management.Metrics.Enabled
}
