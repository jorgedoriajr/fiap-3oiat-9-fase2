package injection

import (
	"hamburgueria/internal/application/api/rest/v1/customer"
	"hamburgueria/internal/application/api/rest/v1/ingredient"
	"hamburgueria/internal/application/api/rest/v1/ingredienttype"
	"hamburgueria/internal/application/api/rest/v1/order"
	"hamburgueria/internal/application/api/rest/v1/product"
	"hamburgueria/internal/application/api/rest/v1/productcategory"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/infra/database"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	ingredientPostgres "hamburgueria/internal/modules/ingredient/infra/database/postgres"
	ingredientService "hamburgueria/internal/modules/ingredient/service"
	ingredientUsecase "hamburgueria/internal/modules/ingredient/usecase"
	orderDatabase "hamburgueria/internal/modules/order/infra/database"
	orderUsecase "hamburgueria/internal/modules/order/usecase"
	"hamburgueria/internal/modules/product/infra/database/postgres"
	"hamburgueria/internal/modules/product/service"
	"hamburgueria/internal/modules/product/usecase"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerController        *customer.CustomerController
	ProductController         *product.Controller
	IngredientController      *ingredient.Controller
	OrderController           *order.Controller
	IngredientTypeController  *ingredienttype.Controller
	ProductCategoryController *productcategory.Controller
	Swagger                   *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	readWriteClient, readOnlyClient := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerPersistence := database.GetCustomerPersistence(readWriteClient, readOnlyClient, logger.Get())

	productPersistence := postgres.NewProductRepository(
		readWriteClient,
		readOnlyClient,
		logger.Get(),
	)

	ingredientPersistence := ingredientPostgres.NewIngredientRepository(
		readWriteClient,
		readOnlyClient,
		logger.Get(),
	)

	ingredientTypePersistence := ingredientPostgres.NewIngredientTypeRepository(
		readWriteClient,
		readOnlyClient,
		logger.Get(),
	)

	productIngredientPersistence := postgres.NewProductIngredientRepository(readWriteClient, readOnlyClient, logger.Get())
	productCategoryPersistence := postgres.NewProductCategoryRepository(readWriteClient, readOnlyClient, logger.Get())

	ingredientFinder := ingredientService.NewIngredientFinderService(ingredientPersistence)
	ingredientTypeFinder := ingredientService.GetIngredientTypeFinderService(ingredientTypePersistence)

	productUseCase := usecase.NewCreateProductUseCase(productPersistence, *ingredientFinder, productIngredientPersistence)
	productFinder := service.NewProductFinderService(productPersistence, *ingredientFinder)

	orderHistoryPersistence := orderDatabase.GetOrderHistoryPersistence(readWriteClient, logger.Get())
	orderProductPersistence := orderDatabase.GetOrderProductPersistence(readWriteClient, logger.Get())
	orderPersistence := orderDatabase.GetOrderPersistence(readWriteClient, readOnlyClient, logger.Get())

	createOrderUseCase := orderUsecase.GetCreateOrderUseCase(
		*productFinder,
		orderPersistence,
		orderHistoryPersistence,
		orderProductPersistence,
	)

	getProductCategoryUseCase := usecase.NewGetProductCategoryUseCase(productCategoryPersistence)

	return DependencyInjection{
		CustomerController: &customer.CustomerController{
			CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
			GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
		},
		ProductController: &product.Controller{
			CreateProductUseCase: productUseCase,
			ProductFinderService: productFinder,
		},
		OrderController: &order.Controller{
			CreateOrderUseCase: createOrderUseCase,
			ListOrderUseCase:   orderUsecase.GetListOrderUseCase(orderPersistence),
		},
		IngredientController: &ingredient.Controller{
			CreateIngredientUseCase: ingredientUsecase.NewCreateIngredientUseCase(ingredientPersistence),
			IngredientFinderService: ingredientFinder,
		},
		ProductCategoryController: &productcategory.Controller{GetProductCategoryUseCase: getProductCategoryUseCase},
		IngredientTypeController:  &ingredienttype.Controller{IngredientTypeFinderService: ingredientTypeFinder},
		Swagger:                   &swagger.Swagger{},
	}
}
