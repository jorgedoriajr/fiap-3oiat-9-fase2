package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/server/api/rest/v1/customer/response"
	"sync"
)

type GetCustomerUseCase struct {
	customerPersistence output.CustomerPersistencePort
}

func (c GetCustomerUseCase) GetCustomer(ctx context.Context, document string) (*response.Customer, error) {
	customer, err := c.customerPersistence.Get(ctx, document)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, err
	}
	return &response.Customer{
		Document: customer.Document,
		Name:     customer.Name,
		Phone:    customer.Phone,
		Email:    customer.Email,
	}, err
}

var (
	getCustomerUseCaseInstance GetCustomerUseCase
	getCustomerUseCaseOnce     sync.Once
)

func GetGetCustomerUseCase(
	CustomerPersistence output.CustomerPersistencePort,
) GetCustomerUseCase {
	getCustomerUseCaseOnce.Do(func() {
		getCustomerUseCaseInstance = GetCustomerUseCase{
			customerPersistence: CustomerPersistence,
		}
	})
	return getCustomerUseCaseInstance
}
