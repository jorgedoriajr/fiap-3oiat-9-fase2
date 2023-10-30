package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/customer/usecase/command"
	"sync"
	"time"
)

type CreateCustomerUseCase struct {
	customerPersistence output.CustomerPersistencePort
}

func (c CreateCustomerUseCase) AddCustomer(ctx context.Context, customer command.CreateCustomerCommand) error {
	return c.customerPersistence.Create(
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

var (
	createCustomerUseCaseInstance CreateCustomerUseCase
	createCustomerUseCaseOnce     sync.Once
)

func GetCreateCustomerUseCase(
	CustomerPersistence output.CustomerPersistencePort,
) CreateCustomerUseCase {
	createCustomerUseCaseOnce.Do(func() {
		createCustomerUseCaseInstance = CreateCustomerUseCase{
			customerPersistence: CustomerPersistence,
		}
	})
	return createCustomerUseCaseInstance
}
