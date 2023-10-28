package injection

import (
	"hamburgueria/pkg/httpserver"
)

func GetAllControllers(injection DependencyInjection) []httpserver.Controller {
	return []httpserver.Controller{
		injection.CustomerController,
		injection.ProductController,
		injection.Swagger,
	}
}
