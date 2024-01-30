package injection

import (
	"hamburgueria/internal/web/api/rest/v1/customer"
	"hamburgueria/internal/web/api/rest/v1/ingredient"
	"hamburgueria/internal/web/api/rest/v1/ingredienttype"
	"hamburgueria/internal/web/api/rest/v1/order"
	"hamburgueria/internal/web/api/rest/v1/payment"
	"hamburgueria/internal/web/api/rest/v1/product"
	"hamburgueria/internal/web/api/rest/v1/productcategory"
	"hamburgueria/internal/web/api/swagger"
	"hamburgueria/internal/web/controller"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerApi           *customer.Api
	ProductApi            *product.Api
	IngredientApi         *ingredient.Api
	OrderApi              *order.Api
	IngredientTypeApi     *ingredienttype.Api
	ProductCategoryApi    *productcategory.Api
	PaymentsStatusWebhook *payment.Webhook
	Swagger               *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	readWriteDB, readOnlyDB := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerUseCaseController := controller.GetCustomerUseCaseController(readWriteDB, readOnlyDB)
	orderUseCaseController := controller.GetOrderUseCaseController(readWriteDB, readOnlyDB)
	productUseCaseController := controller.GetProductUseCaseController(readWriteDB, readOnlyDB)
	ingredientUseCaseController := controller.GetIngredientUseCaseController(readWriteDB, readOnlyDB)
	paymentUseCaseController := controller.GetPaymentUseCaseController(readWriteDB, readOnlyDB)

	return DependencyInjection{
		CustomerApi: &customer.Api{
			CreateCustomerUseCase: customerUseCaseController.CreateCustomerUseCase,
			GetCustomerUseCase:    customerUseCaseController.GetCustomerUseCase,
		},
		ProductApi: &product.Api{
			CreateProductUseCase:  productUseCaseController.CreateProductUseCase,
			FindProductUseCase:    productUseCaseController.FindProductUseCase,
			DeleteProductUseCase:  productUseCaseController.DeleteProductUseCase,
			UpdatedProductUseCase: productUseCaseController.UpdateProductUseCase,
		},
		OrderApi: &order.Api{
			CreateOrderUseCase: orderUseCaseController.CreateOrderUseCase,
			ListOrderUseCase:   orderUseCaseController.ListOrderUseCase,
		},
		IngredientApi: &ingredient.Api{
			CreateIngredientUseCase: ingredientUseCaseController.CreateIngredientUseCase,
			FindIngredientUseCase:   ingredientUseCaseController.FindIngredientUseCase,
		},
		ProductCategoryApi: &productcategory.Api{
			GetProductCategoryUseCase: productUseCaseController.GetProductCategoryUseCase,
		},
		IngredientTypeApi: &ingredienttype.Api{
			FindIngredientTypeUseCase: ingredientUseCaseController.FindIngredientTypeUseCase,
		},
		PaymentsStatusWebhook: &payment.Webhook{
			CreatePaymentStatusUseCase: paymentUseCaseController.CreatePaymentStatus,
		},
		Swagger: &swagger.Swagger{},
	}
}
