package injection

import (
	"hamburgueria/internal/application/api/rest/v1/customer"
	"hamburgueria/internal/application/api/rest/v1/product"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/infra/database"
	"hamburgueria/internal/modules/customer/usecase/create"
	"hamburgueria/internal/modules/customer/usecase/get"
	"hamburgueria/internal/modules/product/infra/database/postgres"
	"hamburgueria/internal/modules/product/service"
	"hamburgueria/internal/modules/product/usecase"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerController *customer.CustomerController
	ProductController  *product.Controller
	Swagger            *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	customerPersistence := database.CustomerRepository{
		ReadWriteClient: sql.GetClient("readWrite"),
		ReadOnlyClient:  sql.GetClient("readOnly"),
		Logger:          logger.Get(),
	}

	productPersistence := postgres.NewProductRepository(
		sql.GetClient("readWrite"),
		sql.GetClient("readOnly"),
		logger.Get(),
	)

	return DependencyInjection{
		CustomerController: &customer.CustomerController{
			CreateCustomerUseCase: create.CreateCustomerUseCase{CustomerPersistence: customerPersistence},
			GetCustomerUseCase:    get.GetCustomerUseCase{CustomerPersistence: customerPersistence},
		},
		ProductController: &product.Controller{
			CreateProductUseCase: usecase.NewCreateProductUseCase(productPersistence),
			ProductFinderService: service.NewProductFinderService(productPersistence),
		},
		Swagger: &swagger.Swagger{},
	}
}
