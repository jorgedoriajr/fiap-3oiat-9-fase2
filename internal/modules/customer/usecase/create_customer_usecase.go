package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/internal/modules/customer/domain/request"
	"hamburgueria/internal/modules/customer/port/output"
	"time"
)

type CreateCustomerUseCase struct {
	customerPersistence output.CustomerPersistencePort
}

func (c CreateCustomerUseCase) AddCustomer(ctx context.Context, customer request.CreateCustomerCommand) error {
	return c.customerPersistence.Create(
		ctx,
		entity.Customer{
			Cpf:       customer.Cpf,
			Name:      customer.Name,
			Phone:     customer.Phone,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
}
