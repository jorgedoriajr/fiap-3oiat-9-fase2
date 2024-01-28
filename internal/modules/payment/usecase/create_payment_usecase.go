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

var (
	processPaymentUseCase     CreatePaymentUseCase
	processPaymentUseCaseOnce sync.Once
)

type CreatePaymentUseCase struct {
	paymentClientGateway      output.PaymentClient
	paymentPersistanceGateway output.PaymentPersistencePort
}

func (p CreatePaymentUseCase) CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	paymentData, err := p.paymentClientGateway.CreatePayment(ctx, command)
	if err != nil {
		return nil, err
	}

	errPersistance := p.paymentPersistanceGateway.Create(ctx, paymentData)
	if errPersistance != nil {
		return nil, errPersistance
	}

	return mapperPaymentEntityToPaymentProcessed(&paymentData), nil
}

func GetCreatePaymentUseCase(paymentClientGateway output.PaymentClient, paymentPersistanceGateway output.PaymentPersistencePort) input.CreatePaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCase = CreatePaymentUseCase{paymentClientGateway: paymentClientGateway, paymentPersistanceGateway: paymentPersistanceGateway}

	})
	return processPaymentUseCase
}

func mapperPaymentEntityToPaymentProcessed(payment *domain.Payment) *result.PaymentProcessed {
	return &result.PaymentProcessed{
		PaymentId:   payment.Id,
		OrderId:     payment.OrderId,
		PaymentData: payment.Data,
		CreatedAt:   payment.CreatedAt,
	}
}
