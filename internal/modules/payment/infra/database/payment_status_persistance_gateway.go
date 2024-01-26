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

type PaymentStatusPersistanceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (ps *PaymentStatusPersistanceGateway) CreatePaymentStatus(ctx context.Context, paymentStatus domain.PaymentIntegrationLog) error {
	tx := ps.readWriteClient.Create(&model.PaymentStatus{
		Id:                   paymentStatus.Id,
		PaymentIntegrationId: paymentStatus.PaymentIntegrationId,
		PaymentStatus:        string(paymentStatus.PaymentStatus),
	})
	if tx.Error != nil {
		ps.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Str("order", paymentStatus.Id.String()).
			Msg("Failed to insert payment status")
		return tx.Error
	}
	return nil
}

func (ps *PaymentStatusPersistanceGateway) FindPaymentStatus(ctx context.Context, paymentStatusId uuid.UUID) (*domain.PaymentIntegrationLog, error) {

	var paymentStatus model.PaymentStatus

	err := ps.readOnlyClient.Last(&paymentStatus, paymentStatusId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		ps.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("payment", paymentStatusId.String()).
			Msg("Failed to get payment status by id")
		return nil, err
	}

	if paymentStatus.Id.String() == "" {
		return nil, nil
	}
	return paymentStatus.ToDomain(), nil
}

var (
	paymentStatusRepositoryInstance PaymentStatusPersistanceGateway
	paymentStatusRepositoryOnce     sync.Once
)

func GetPaymentStatusPersistenceGateway(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) PaymentStatusPersistanceGateway {
	paymentStatusRepositoryOnce.Do(func() {
		paymentStatusRepositoryInstance = PaymentStatusPersistanceGateway{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return paymentStatusRepositoryInstance
}
