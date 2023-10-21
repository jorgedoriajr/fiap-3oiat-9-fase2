package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/response"
	"hamburgueria/internal/modules/customer/port/output"
)

type GetCustomerUseCase struct {
	customerPersistence output.CustomerPersistencePort
}

func (c GetCustomerUseCase) GetCustomer(ctx context.Context, document string) (*response.CustomerResponse, error) {
	customer, err := c.customerPersistence.Get(ctx, document)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, err
	}
	return &response.CustomerResponse{
		Cpf:   customer.Cpf,
		Name:  customer.Name,
		Phone: customer.Phone,
	}, err
}
