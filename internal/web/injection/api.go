package injection

import (
	"hamburgueria/pkg/httpserver"
)

func GetAllApis(injection DependencyInjection) []httpserver.Controller {
	return []httpserver.Controller{
		injection.CustomerApi,
		injection.ProductApi,
		injection.IngredientApi,
		injection.OrderApi,
		injection.IngredientTypeApi,
		injection.ProductCategoryApi,
		injection.Swagger,
	}
}
