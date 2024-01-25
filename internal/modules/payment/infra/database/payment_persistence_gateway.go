package database

import (
	"context"
	"errors"
	"hamburgueria/internal/modules/payment/domain"
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

func NewPaymentPersistenceGateway(readWriteClient, readOnlyClient *gorm.DB, logger zerolog.Logger) *PaymentPersistenceGateway {
	return &PaymentPersistenceGateway{
		readWriteClient: readWriteClient,
		readOnlyClient:  readOnlyClient,
		logger:          logger,
	}
}

func (p *PaymentPersistenceGateway) Create(ctx context.Context, payment domain.Payment) error {
	if err := p.readWriteClient.Create(&payment).Error; err != nil {
		p.logger.Error().Err(err).Msg("Failed to create payment")
		return err
	}
	return nil
}

func (p *PaymentPersistenceGateway) GetAll(ctx context.Context) ([]domain.Payment, error) {
	var payments []domain.Payment
	if err := p.readOnlyClient.Find(&payments).Error; err != nil {
		p.logger.Error().Err(err).Msg("Failed to get all payments")
		return nil, err
	}
	return payments, nil
}

func (p *PaymentPersistenceGateway) FindByStatus(ctx context.Context, status string) ([]domain.Payment, error) {
	var payments []domain.Payment
	if err := p.readOnlyClient.Where("status = ?", status).Find(&payments).Error; err != nil {
		p.logger.Error().Err(err).Msg("Failed to find payments by status")
		return nil, err
	}
	return payments, nil
}

func (p *PaymentPersistenceGateway) FindById(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error) {
	var payment domain.Payment
	if err := p.readOnlyClient.Where("id = ?", paymentId).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		p.logger.Error().Err(err).Msg("Failed to find payment by ID")
		return nil, err
	}
	return &payment, nil
}

func (p *PaymentPersistenceGateway) Update(ctx context.Context, payment domain.Payment) error {
	if err := p.readWriteClient.Save(&payment).Error; err != nil {
		p.logger.Error().Err(err).Msg("Failed to update payment")
		return err
	}
	return nil
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
