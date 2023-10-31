package injection

import (
	"hamburgueria/pkg/httpserver"
)

func GetAllControllers(injection DependencyInjection) []httpserver.Controller {
	return []httpserver.Controller{
		injection.CustomerController,
		injection.ProductController,
		injection.IngredientController,
		injection.OrderController,
		injection.IngredientTypeController,
		injection.Swagger,
	}
}
