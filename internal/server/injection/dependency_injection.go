package injection

import (
	"hamburgueria/internal/modules/customer/infra/database"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	ingredientDatabase "hamburgueria/internal/modules/ingredient/infra/database"
	ingredientUsecase "hamburgueria/internal/modules/ingredient/usecase"
	orderDatabase "hamburgueria/internal/modules/order/infra/database"
	orderUsecase "hamburgueria/internal/modules/order/usecase"
	paymentUseCase "hamburgueria/internal/modules/payment/usecase"
	productDatabase "hamburgueria/internal/modules/product/infra/database"
	"hamburgueria/internal/modules/product/usecase"
	"hamburgueria/internal/server/api/rest/v1/customer"
	"hamburgueria/internal/server/api/rest/v1/ingredient"
	"hamburgueria/internal/server/api/rest/v1/ingredienttype"
	"hamburgueria/internal/server/api/rest/v1/order"
	"hamburgueria/internal/server/api/rest/v1/product"
	"hamburgueria/internal/server/api/rest/v1/productcategory"
	"hamburgueria/internal/server/api/swagger"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerController        *customer.Controller
	ProductController         *product.Controller
	IngredientController      *ingredient.Controller
	OrderController           *order.Controller
	IngredientTypeController  *ingredienttype.Controller
	ProductCategoryController *productcategory.Controller
	Swagger                   *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	readWriteDB, readOnlyDB := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerPersistence := database.GetCustomerPersistence(readWriteDB, readOnlyDB, logger.Get())

	ingredientTypePersistence := ingredientDatabase.GetIngredientTypeRepository(readWriteDB, readOnlyDB, logger.Get())
	ingredientPersistence := ingredientDatabase.GetIngredientRepository(readWriteDB, readOnlyDB, logger.Get())

	productCategoryPersistence := productDatabase.GetProductCategoryRepository(readWriteDB, readOnlyDB, logger.Get())

	productPersistence := productDatabase.GetProductRepository(readWriteDB, readOnlyDB, logger.Get())
	findProductCategoryUseCase := usecase.NewGetProductCategoryUseCase(productCategoryPersistence)
	createProductUseCase := usecase.GetCreateProductUseCase(productPersistence, ingredientPersistence, productCategoryPersistence)
	deleteProductUseCase := usecase.GetDeleteProductUseCase(productPersistence)
	updateProductUseCase := usecase.GetUpdateProductUseCase(productPersistence, ingredientPersistence)
	findProductUseCase := usecase.NewFindProductUseCase(productPersistence)

	orderPersistence := orderDatabase.GetOrderPersistence(readWriteDB, readOnlyDB, logger.Get())

	createPaymentUseCase := paymentUseCase.GetCreatePaymentUseCase()
	processPaymentUseCase := orderUsecase.GetProcessPaymentUseCase(orderPersistence, createPaymentUseCase)

	createOrderUseCase := orderUsecase.GetCreateOrderUseCase(
		productPersistence,
		orderPersistence,
		processPaymentUseCase,
		customerPersistence,
	)

	return DependencyInjection{
		CustomerController: &customer.Controller{
			CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
			GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
		},
		ProductController: &product.Controller{
			CreateProductUseCase:  createProductUseCase,
			FindProductUseCase:    findProductUseCase,
			DeleteProductUseCase:  deleteProductUseCase,
			UpdatedProductUseCase: updateProductUseCase,
		},
		OrderController: &order.Controller{
			CreateOrderUseCase: createOrderUseCase,
			ListOrderUseCase:   orderUsecase.GetListOrderUseCase(orderPersistence),
		},
		IngredientController: &ingredient.Controller{
			CreateIngredientUseCase: ingredientUsecase.NewCreateIngredientUseCase(ingredientPersistence, ingredientTypePersistence),
			FindIngredientUseCase:   ingredientUsecase.NewFindIngredientUseCase(ingredientPersistence),
		},
		ProductCategoryController: &productcategory.Controller{GetProductCategoryUseCase: findProductCategoryUseCase},
		IngredientTypeController: &ingredienttype.Controller{
			FindIngredientTypeUseCase: ingredientUsecase.GetIngredientTypeUseCase(ingredientTypePersistence),
		},
		Swagger: &swagger.Swagger{},
	}
}
