package rest

import (
	"hamburgueria/internal/application/injection"
	"hamburgueria/pkg/httpserver"
)

func GetAllControllers(injection injection.DependencyInjection) []httpserver.Controller {
	return []httpserver.Controller{
		injection.CustomerController,
	}
}
