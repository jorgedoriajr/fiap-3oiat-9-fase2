package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"hamburgueria/pkg/validation"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Listen()
	ShutDown()
}

type server struct {
	echoInstance       *echo.Echo
	managementInstance *echo.Echo
	builder            *serverBuilder
}

type Controller interface {
	RegisterEchoRoutes(e *echo.Echo)
}

func (s *server) Listen() {
	s.configureMainServer()
	s.configureManagementServer()
	s.startManagementServer()

	serverPort := fmt.Sprintf(":%d", s.builder.serverConfig.Port)
	if err := s.echoInstance.Start(serverPort); !errors.Is(err, http.ErrServerClosed) {
		panic("[Server] Fatal error while starting the http server")
	}
}

func (s *server) ShutDown() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.echoInstance.Shutdown(ctx); err != nil {
		panic("[Server] Fatal error while shutting down the server")
	}
}

func (s *server) configureMainServer() {
	s.echoInstance = echo.New()

	for _, m := range s.builder.middleware {
		s.echoInstance.Use(m)
	}

	for _, controller := range s.builder.controllers {
		controller.RegisterEchoRoutes(s.echoInstance)
	}

	if s.builder.errorHandler != nil {
		s.echoInstance.HTTPErrorHandler = s.builder.errorHandler
	}

	if s.builder.validator != nil {
		s.echoInstance.Validator = s.builder.validator
	}

	if s.builder.serverConfig.UseRequestValidator && s.builder.validator == nil {
		s.echoInstance.Validator = &validation.DefaultValidator{Validator: validator.New()}
	}
}
