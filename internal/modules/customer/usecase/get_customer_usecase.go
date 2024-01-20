package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/customer/usecase/result"
	"sync"
)

type GetCustomerUseCase struct {
	customerPersistenceGateway output.CustomerPersistencePort
}

func (c GetCustomerUseCase) GetCustomer(ctx context.Context, document string) (*result.CustomerCreated, error) {
	customer, err := c.customerPersistenceGateway.Get(ctx, document)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, err
	}
	return &result.CustomerCreated{
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
	CustomerPersistenceGateway output.CustomerPersistencePort,
) GetCustomerUseCase {
	getCustomerUseCaseOnce.Do(func() {
		getCustomerUseCaseInstance = GetCustomerUseCase{
			customerPersistenceGateway: CustomerPersistenceGateway,
		}
	})
	return getCustomerUseCaseInstance
}
