package usecase

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/result"

	"sync"

	"github.com/google/uuid"
)

type FindPaymentStatusUseCase struct {
	paymentStatusPersistenceGateway output.PaymentStatusPersistencePort
}

func (ps FindPaymentStatusUseCase) ListPaymentStatus(ctx context.Context, paymentStatusId uuid.UUID) (*result.PaymentStatusProcessed, error) {

	paymentStatusData, err := ps.paymentStatusPersistenceGateway.FindPaymentStatus(ctx, paymentStatusId)
	if err != nil {
		return nil, err
	}
	return mapperPaymentStatusEntityToPaymentProcessed(paymentStatusData), nil
}

var (
	findPaymentStatusUseCase     FindPaymentStatusUseCase
	findPaymentStatusUseCaseOnce sync.Once
)

func GetFindPaymentStatusUseCase(paymentStatusPersistenceGateway output.PaymentStatusPersistencePort) input.ListPaymentStatusPort {
	findPaymentStatusUseCaseOnce.Do(func() {
		findPaymentStatusUseCase = FindPaymentStatusUseCase{paymentStatusPersistenceGateway: paymentStatusPersistenceGateway}

	})
	return findPaymentStatusUseCase
}

func mapperPaymentStatusEntityToPaymentProcessed(paymentStatus *domain.PaymentStatus) *result.PaymentStatusProcessed {
	return &result.PaymentStatusProcessed{
		Id:                   paymentStatus.Id,
		PaymentId:            paymentStatus.PaymentId,
		PaymentIntegrationId: paymentStatus.PaymentIntegrationId,
		Status:               paymentStatus.PaymentStatus,
	}
}
