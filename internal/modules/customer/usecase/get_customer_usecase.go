package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/response"
	"hamburgueria/internal/modules/customer/port/output"
)

type GetCustomerUseCase struct {
	CustomerPersistence output.CustomerPersistencePort
}

func (c GetCustomerUseCase) GetCustomer(ctx context.Context, document string) (*response.CustomerResponse, error) {
	customer, err := c.CustomerPersistence.Get(ctx, document)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, err
	}
	return &response.CustomerResponse{
		Document: customer.Document,
		Name:     customer.Name,
		Phone:    customer.Phone,
		Email:    customer.Email,
	}, err
}
