package database

import (
	"context"
	"errors"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/infra/database/model"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type PaymentPersistenceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (p *PaymentPersistenceGateway) Create(ctx context.Context, payment domain.Payment) error {
	tx := p.readWriteClient.Create(&model.Payment{
		Id:      payment.Id,
		OrderId: payment.OrderId,
		Data:    payment.Data,
	})
	if tx.Error != nil {
		p.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Str("order", payment.OrderId.String()).
			Msg("Failed to insert payment")
		return tx.Error
	}
	return nil
}

func (p *PaymentPersistenceGateway) FindById(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error) {

	var payment model.Payment

	err := p.readOnlyClient.First(&payment, paymentId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		p.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("payment", paymentId.String()).
			Msg("Failed to get payment by id")
		return nil, err
	}

	if payment.Id.String() == "" {
		return nil, nil
	}
	return payment.ToDomain(), nil
}

var (
	paymentRepositoryInstance PaymentPersistenceGateway
	paymentRepositoryOnce     sync.Once
)

func GetPaymentPersistenceGateway(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) PaymentPersistenceGateway {
	paymentRepositoryOnce.Do(func() {
		paymentRepositoryInstance = PaymentPersistenceGateway{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return paymentRepositoryInstance
}
