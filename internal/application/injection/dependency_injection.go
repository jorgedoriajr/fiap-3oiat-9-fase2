package injection

import (
	"hamburgueria/internal/application/api/rest/v1/customer"
	"hamburgueria/internal/application/api/rest/v1/ingredient"
	"hamburgueria/internal/application/api/rest/v1/ingredienttype"
	"hamburgueria/internal/application/api/rest/v1/order"
	"hamburgueria/internal/application/api/rest/v1/product"
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
	CustomerController       *customer.CustomerController
	ProductController        *product.Controller
	IngredientController     *ingredient.Controller
	OrderController          *order.Controller
	IngredientTypeController *ingredienttype.Controller
	Swagger                  *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	ReadWriteClient, ReadOnlyClient := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerPersistence := database.GetCustomerPersistence(ReadWriteClient, ReadOnlyClient, logger.Get())

	productPersistence := postgres.NewProductRepository(
		ReadWriteClient,
		ReadOnlyClient,
		logger.Get(),
	)

	ingredientPersistence := ingredientPostgres.NewIngredientRepository(
		ReadWriteClient,
		ReadOnlyClient,
		logger.Get(),
	)

	ingredientTypePersistence := ingredientPostgres.NewIngredientTypeRepository(
		ReadWriteClient,
		ReadOnlyClient,
		logger.Get(),
	)

	productIngredientPersistence := postgres.NewProductIngredientRepository(ReadWriteClient, ReadOnlyClient, logger.Get())

	ingredientFinder := ingredientService.NewIngredientFinderService(ingredientPersistence)
	ingredientTypeFinder := ingredientService.GetIngredientTypeFinderService(ingredientTypePersistence)

	productUseCase := usecase.NewCreateProductUseCase(productPersistence, *ingredientFinder, productIngredientPersistence)

	orderHistoryPersistence := orderDatabase.GetOrderHistoryPersistence(ReadWriteClient, logger.Get())
	orderProductPersistence := orderDatabase.GetOrderProductPersistence(ReadWriteClient, logger.Get())
	orderPersistence := orderDatabase.GetOrderPersistence(ReadWriteClient, logger.Get())

	createOrderUseCase := orderUsecase.GetCreateOrderUseCase(
		productUseCase,
		productPersistence,
		orderPersistence,
		orderHistoryPersistence,
		orderProductPersistence,
	)

	return DependencyInjection{
		CustomerController: &customer.CustomerController{
			CreateCustomerUseCase: customerUseCase.GetCreateCustomerUseCase(customerPersistence),
			GetCustomerUseCase:    customerUseCase.GetGetCustomerUseCase(customerPersistence),
		},
		ProductController: &product.Controller{
			CreateProductUseCase: productUseCase,
			ProductFinderService: service.NewProductFinderService(productPersistence, *ingredientFinder),
		},
		OrderController: &order.Controller{
			CreateOrderUseCase: createOrderUseCase,
		},
		IngredientController: &ingredient.Controller{
			CreateIngredientUseCase: ingredientUsecase.NewCreateIngredientUseCase(ingredientPersistence),
			IngredientFinderService: ingredientFinder,
		},
		IngredientTypeController: &ingredienttype.Controller{IngredientTypeFinderService: ingredientTypeFinder},
		Swagger:                  &swagger.Swagger{},
	}
}
