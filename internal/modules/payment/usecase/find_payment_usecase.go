package usecase

import (
	"context"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/result"

	"sync"

	"github.com/google/uuid"
)

type FindPaymentUseCase struct {
	paymentClientGateway output.PaymentPersistencePort
}

func (p FindPaymentUseCase) FindPaymentById(ctx context.Context, paymentId uuid.UUID) (*result.PaymentProcessed, error) {
	paymentData, err := p.paymentClientGateway.FindById(ctx, paymentId)
	if err != nil {
		return nil, err
	}
	return mapperPaymentEntityToPaymentProcessed(paymentData), nil
}

var (
	findPaymentUseCase     FindPaymentUseCase
	findPaymentUseCaseOnce sync.Once
)

func GetFindPaymentUseCase(paymentClientGateway output.PaymentPersistencePort) input.FindPaymentPort {
	findPaymentUseCaseOnce.Do(func() {
		findPaymentUseCase = FindPaymentUseCase{paymentClientGateway: paymentClientGateway}

	})
	return findPaymentUseCase
}
