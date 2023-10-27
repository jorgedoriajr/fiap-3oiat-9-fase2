package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/customer/usecase/command"
	"time"
)

type CreateCustomerUseCase struct {
	CustomerPersistence output.CustomerPersistencePort
}

func (c CreateCustomerUseCase) AddCustomer(ctx context.Context, customer command.CreateCustomerCommand) error {
	return c.CustomerPersistence.Create(
		ctx,
		entity.Customer{
			Document:       customer.Document,
			Name:           customer.Name,
			Phone:          customer.Phone,
			Email:          customer.Email,
			OptInPromotion: customer.OptInPromotion,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	)
}
