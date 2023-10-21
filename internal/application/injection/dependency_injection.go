package injection

import (
	"hamburgueria/internal/application/api/rest"
	"hamburgueria/internal/modules/customer/usecase"
)

type DependencyInjection struct {
	CustomerController *rest.CustomerController
}

func NewDependencyInjection() DependencyInjection {
	return DependencyInjection{
		CustomerController: &rest.CustomerController{
			CreateCustomerUseCase: usecase.CreateCustomerUseCase{},
			GetCustomerUseCase:    usecase.GetCustomerUseCase{},
		},
	}
}
