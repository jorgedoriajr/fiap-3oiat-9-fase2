package database

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"hamburgueria/internal/modules/customer/domain"
	"hamburgueria/internal/modules/customer/infra/database/model"
	"hamburgueria/internal/modules/customer/port/output"
	"sync"
	"time"
)

type CustomerPersistenceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c CustomerPersistenceGateway) Create(ctx context.Context, customer domain.Customer) error {
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

func (c CustomerPersistenceGateway) Get(ctx context.Context, document string) (*domain.Customer, error) {

	var customer model.Customer
	err := c.readOnlyClient.First(&customer, document).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("document", customer.Cpf).
			Msg("Failed to get customer by document")
		return nil, err
	}

	if customer.Cpf == "" {
		return nil, nil
	}
	return customer.ToDomain(), nil
}

var (
	customerRepositoryInstance output.CustomerPersistencePort
	customerRepositoryOnce     sync.Once
)

func GetCustomerPersistence(
	ReadWriteClient *gorm.DB,
	ReadOnlyClient *gorm.DB,
	Logger zerolog.Logger,
) output.CustomerPersistencePort {
	customerRepositoryOnce.Do(func() {
		customerRepositoryInstance = CustomerPersistenceGateway{
			readWriteClient: ReadWriteClient,
			readOnlyClient:  ReadOnlyClient,
			logger:          Logger,
		}
	})
	return customerRepositoryInstance
}
