package database

import (
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"hamburgueria/internal/modules/customer/domain"
	"hamburgueria/internal/modules/customer/infra/database/model"
	"sync"
	"time"
)

type CustomerRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c CustomerRepository) Create(ctx context.Context, customer domain.Customer) error {
	tx := c.readWriteClient.Create(&model.Customer{
		Cpf:            customer.Document,
		Name:           customer.Name,
		Phone:          customer.Phone,
		Email:          customer.Email,
		OptInPromotion: customer.OptInPromotion,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Str("document", customer.Document).
			Msg("Failed to insert customer")
		return tx.Error
	}
	return nil
}

func (c CustomerRepository) Get(ctx context.Context, document string) (customerResult *domain.Customer, err error) {

	var customer model.Customer
	tx := c.readOnlyClient.First(&customer, document)

	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Str("document", customer.Cpf).
			Msg("Failed to insert customer")
		return nil, tx.Error
	}

	if customer.Cpf == "" {
		return nil, nil
	}
	return customer.ToDomain(), nil
}

var (
	customerRepositoryInstance CustomerRepository
	customerRepositoryOnce     sync.Once
)

func GetCustomerPersistence(
	ReadWriteClient *gorm.DB,
	ReadOnlyClient *gorm.DB,
	Logger zerolog.Logger,
) CustomerRepository {
	customerRepositoryOnce.Do(func() {
		customerRepositoryInstance = CustomerRepository{
			readWriteClient: ReadWriteClient,
			readOnlyClient:  ReadOnlyClient,
			logger:          Logger,
		}
	})
	return customerRepositoryInstance
}
