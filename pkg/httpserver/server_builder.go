package httpserver

import (
	"github.com/labstack/echo/v4"
	"hamburgueria/config"
	"hamburgueria/pkg/healthcheck"
)

type ServerBuilder interface {
	WithConfig(serverConfig config.HttpServerConfig) ServerBuilder
	WithControllers(controllers ...Controller) ServerBuilder
	WithHealthCheck(checkers ...healthcheck.HealthChecker) ServerBuilder
	WithMiddleware(m ...echo.MiddlewareFunc) ServerBuilder
	WithHTTPErrorHandler(handler echo.HTTPErrorHandler) ServerBuilder
	WithValidator(validator echo.Validator) ServerBuilder
	Build() Server
}

type serverBuilder struct {
	controllers    []Controller
	serverConfig   config.HttpServerConfig
	healthCheckers []healthcheck.HealthChecker
	middleware     []echo.MiddlewareFunc
	errorHandler   echo.HTTPErrorHandler
	validator      echo.Validator
}

func Builder() ServerBuilder {
	return &serverBuilder{}
}

func (sb serverBuilder) WithConfig(serverConfig config.HttpServerConfig) ServerBuilder {
	sb.serverConfig = serverConfig
	return &sb
}

func (sb serverBuilder) WithControllers(controllers ...Controller) ServerBuilder {
	sb.controllers = append(sb.controllers, controllers...)
	return &sb
}

func (sb serverBuilder) WithHealthCheck(checkers ...healthcheck.HealthChecker) ServerBuilder {
	sb.healthCheckers = append(sb.healthCheckers, checkers...)
	return &sb
}

func (sb serverBuilder) WithMiddleware(m ...echo.MiddlewareFunc) ServerBuilder {
	sb.middleware = append(sb.middleware, m...)
	return &sb
}

func (sb serverBuilder) WithHTTPErrorHandler(h echo.HTTPErrorHandler) ServerBuilder {
	sb.errorHandler = h
	return &sb
}

func (sb serverBuilder) WithValidator(v echo.Validator) ServerBuilder {
	sb.validator = v
	return &sb
}

func (sb serverBuilder) Build() Server {
	return &server{
		builder: &sb,
	}
}
