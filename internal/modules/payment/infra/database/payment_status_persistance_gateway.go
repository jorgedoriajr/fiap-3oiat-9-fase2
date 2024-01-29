package database

import (
	"context"
	"errors"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/infra/database/model"
	"hamburgueria/internal/modules/payment/port/output"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type PaymentStatusPersistenceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (ps PaymentStatusPersistenceGateway) CreatePaymentStatus(ctx context.Context, paymentStatus domain.PaymentStatus) error {
	tx := ps.readWriteClient.Create(&model.PaymentStatus{
		Id:                   paymentStatus.Id,
		PaymentId:            paymentStatus.PaymentId,
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

func (ps PaymentStatusPersistenceGateway) FindPaymentStatus(ctx context.Context, paymentStatusId uuid.UUID) (*domain.PaymentStatus, error) {

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
	paymentStatusRepositoryInstance output.PaymentStatusPersistencePort
	paymentStatusRepositoryOnce     sync.Once
)

func GetPaymentStatusPersistenceGateway(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.PaymentStatusPersistencePort {
	paymentStatusRepositoryOnce.Do(func() {
		paymentStatusRepositoryInstance = PaymentStatusPersistenceGateway{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return paymentStatusRepositoryInstance
}
