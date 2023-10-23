package injection

import (
	"hamburgueria/internal/application/api/rest"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/infra/database"
	"hamburgueria/internal/modules/customer/usecase"
	"hamburgueria/pkg/logger"
	"hamburgueria/pkg/sql"
)

type DependencyInjection struct {
	CustomerController *rest.CustomerController
	Swagger            *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {

	customerPersistence := database.CustomerRepository{
		ReadWriteClient: sql.GetClient("readWrite"),
		ReadOnlyClient:  sql.GetClient("readOnly"),
		Logger:          logger.Get(),
	}

	return DependencyInjection{
		CustomerController: &rest.CustomerController{
			CreateCustomerUseCase: usecase.CreateCustomerUseCase{CustomerPersistence: customerPersistence},
			GetCustomerUseCase:    usecase.GetCustomerUseCase{CustomerPersistence: customerPersistence},
		},
		Swagger: &swagger.Swagger{},
	}
}
