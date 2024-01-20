package usecase

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/internal/modules/payment/usecase/result"

	"sync"
)

type CreatePaymentUseCase struct {
	paymentClientGateway output.PaymentClient
}

func (p CreatePaymentUseCase) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	paymentData, err := p.paymentClientGateway.CreatePayment(ctx, command)
	if err != nil {
		return nil, err
	}
	return mapperPaymentEntityToPaymentProcessed(paymentData), nil
}

var (
	processPaymentUseCase     CreatePaymentUseCase
	processPaymentUseCaseOnce sync.Once
)

func GetCreatePaymentUseCase(paymentClientGateway output.PaymentClient) input.CreatePaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCase = CreatePaymentUseCase{paymentClientGateway: paymentClientGateway}

	})
	return processPaymentUseCase
}

func mapperPaymentEntityToPaymentProcessed(payment *domain.Payment) *result.PaymentProcessed {
	return &result.PaymentProcessed{
		PaymentId:   payment.Id,
		PaymentData: payment.Data,
	}
}
