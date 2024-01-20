package usecase

import (
	"context"
	"hamburgueria/internal/modules/customer/domain"
	"hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/customer/usecase/command"
	"sync"
	"time"
)

type CreateCustomerUseCase struct {
	customerPersistenceGateway output.CustomerPersistencePort
}

func (c CreateCustomerUseCase) AddCustomer(ctx context.Context, customer command.CreateCustomerCommand) error {
	return c.customerPersistenceGateway.Create(
		ctx,
		domain.Customer{
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
	CustomerPersistenceGateway output.CustomerPersistencePort,
) CreateCustomerUseCase {
	createCustomerUseCaseOnce.Do(func() {
		createCustomerUseCaseInstance = CreateCustomerUseCase{
			customerPersistenceGateway: CustomerPersistenceGateway,
		}
	})
	return createCustomerUseCaseInstance
}
