package injection

import (
	"hamburgueria/internal/application/api/rest"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/infra/database"
	customerUseCase "hamburgueria/internal/modules/customer/usecase"
	orderDatabase "hamburgueria/internal/modules/order/infra/database"
	orderUseCase "hamburgueria/internal/modules/order/usecase"
	productDatabase "hamburgueria/internal/modules/product/infra/database/postgres"
	productUseCase "hamburgueria/internal/modules/product/usecase"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerController *rest.CustomerController
	OrderController    *rest.OrderController
	Swagger            *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	customerPersistence := database.CustomerRepository{
		ReadWriteClient: sql.GetClient("readWrite"),
		ReadOnlyClient:  sql.GetClient("readOnly"),
		Logger:          logger.Get(),
	}

	orderPersistence := orderDatabase.OrderRepository{
		ReadWriteClient: sql.GetClient("readWrite"),
		Logger:          logger.Get(),
	}

	productPersistence := productDatabase.ProductRepository{
		ReadWriteClient: sql.GetClient("readWrite"),
		ReadOnlyClient:  sql.GetClient("readOnly"),
		Logger:          logger.Get(),
	}

	createProductUseCase := productUseCase.CreateProductUseCase{
		ProductPersistencePort: productPersistence,
	}

	return DependencyInjection{
		CustomerController: &rest.CustomerController{
			CreateCustomerUseCase: customerUseCase.CreateCustomerUseCase{CustomerPersistence: customerPersistence},
			GetCustomerUseCase:    customerUseCase.GetCustomerUseCase{CustomerPersistence: customerPersistence},
		},
		OrderController: &rest.OrderController{
			CreateOrderUseCase: orderUseCase.CreateOrderUseCase{
				ProductUseCase:     createProductUseCase,
				ProductPersistence: productPersistence,
				OrderPersistence:   orderPersistence,
			},
		},
		Swagger: &swagger.Swagger{},
	}
}
