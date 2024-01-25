package usecase

import (
	"context"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/internal/modules/payment/usecase/result"

	"sync"
)

type FindPaymentUseCase struct {
	paymentClientGateway output.PaymentClient
}

func (p FindPaymentUseCase) FindPayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	paymentData, err := p.paymentClientGateway.CreatePayment(ctx, command)
	if err != nil {
		return nil, err
	}
	return mapperPaymentEntityToPaymentProcessed(paymentData), nil
}

var (
	findPaymentUseCase     FindPaymentUseCase
	findPaymentUseCaseOnce sync.Once
)

func GetFindPaymentUseCase(paymentClientGateway output.PaymentClient) input.CreatePaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCase = CreatePaymentUseCase{paymentClientGateway: paymentClientGateway}

	})
	return processPaymentUseCase
}
