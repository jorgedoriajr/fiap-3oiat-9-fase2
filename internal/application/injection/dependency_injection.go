package injection

import (
	"hamburgueria/internal/application/api/rest/v1/customer"
	"hamburgueria/internal/application/api/rest/v1/ingredient"
	"hamburgueria/internal/application/api/rest/v1/product"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/infra/database"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	postgres2 "hamburgueria/internal/modules/ingredient/infra/database/postgres"
	service2 "hamburgueria/internal/modules/ingredient/service"
	usecase2 "hamburgueria/internal/modules/ingredient/usecase"
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

	ingredientPersistence := postgres2.NewIngredientRepository(
		ReadWriteClient,
		ReadOnlyClient,
		logger.Get(),
	)

	return DependencyInjection{
		CustomerController: &customer.CustomerController{
			CreateCustomerUseCase: customerUseCase.CreateCustomerUseCase{CustomerPersistence: customerPersistence},
			GetCustomerUseCase:    customerUseCase.GetCustomerUseCase{CustomerPersistence: customerPersistence},
		},
		ProductController: &product.Controller{
			CreateProductUseCase: usecase.NewCreateProductUseCase(productPersistence),
			ProductFinderService: service.NewProductFinderService(productPersistence),
		},

		IngredientController: &ingredient.Controller{
			CreateIngredientUseCase: usecase2.NewCreateIngredientUseCase(ingredientPersistence),
			IngredientFinderService: service2.NewIngredientFinderService(ingredientPersistence),
		},
		Swagger: &swagger.Swagger{},
	}
}
