package injection

import (
	"hamburgueria/internal/application/api/rest"
	"hamburgueria/internal/application/api/swagger"
	"hamburgueria/internal/modules/customer/usecase"
)

type DependencyInjection struct {
	CustomerController *rest.CustomerController
	Swagger            *swagger.Swagger
}

func NewDependencyInjection() DependencyInjection {
	return DependencyInjection{
		CustomerController: &rest.CustomerController{
			CreateCustomerUseCase: usecase.CreateCustomerUseCase{},
			GetCustomerUseCase:    usecase.GetCustomerUseCase{},
		},
		Swagger: &swagger.Swagger{},
	}
}
