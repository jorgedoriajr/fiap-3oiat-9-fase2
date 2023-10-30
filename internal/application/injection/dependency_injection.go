package injection

import (
	"hamburgueria/internal/application/api/rest/v1/customer"
	"hamburgueria/internal/application/api/rest/v1/ingredient"
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
	CustomerController   *customer.CustomerController
	ProductController    *product.Controller
	IngredientController *ingredient.Controller
	OrderController      *order.Controller
	Swagger              *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	ReadWriteClient, ReadOnlyClient := sql.GetClient("readWrite"), sql.GetClient("readOnly")

	customerPersistence := database.CustomerRepository{
		ReadWriteClient: ReadWriteClient,
		ReadOnlyClient:  ReadOnlyClient,
		Logger:          logger.Get(),
	}

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

	productUseCase := usecase.NewCreateProductUseCase(productPersistence)

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
			CreateCustomerUseCase: customerUseCase.CreateCustomerUseCase{CustomerPersistence: customerPersistence},
			GetCustomerUseCase:    customerUseCase.GetCustomerUseCase{CustomerPersistence: customerPersistence},
		},
		ProductController: &product.Controller{
			CreateProductUseCase: productUseCase,
			ProductFinderService: service.NewProductFinderService(productPersistence),
		},
		OrderController: &order.Controller{
			CreateOrderUseCase: createOrderUseCase,
		},
		IngredientController: &ingredient.Controller{
			CreateIngredientUseCase: ingredientUsecase.NewCreateIngredientUseCase(ingredientPersistence),
			IngredientFinderService: ingredientService.NewIngredientFinderService(ingredientPersistence),
		},
		Swagger: &swagger.Swagger{},
	}
}
