package usecase

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/command"

	"sync"
)

type CreatePaymentStatusUseCase struct {
	paymentStatusPersistenceGateway output.PaymentStatusPersistencePort
}

func (ps CreatePaymentStatusUseCase) AddPaymentStatus(ctx context.Context, command command.CreatePaymentStatusCommand) error {
	paymentStatus := mapperPaymentStatusCommandToEntityPaymentStatus(command)
	err := ps.paymentStatusPersistenceGateway.CreatePaymentStatus(ctx, paymentStatus)
	if err != nil {
		return err
	}

	return nil
}

var (
	processPaymentStatusUseCase     CreatePaymentStatusUseCase
	processPaymentStatusUseCaseOnce sync.Once
)

func GetCreatePaymentStatusUseCase(paymentStatusPersistanceGateway output.PaymentStatusPersistencePort) input.CreatePaymentStatusPort {
	processPaymentStatusUseCaseOnce.Do(func() {
		processPaymentStatusUseCase = CreatePaymentStatusUseCase{paymentStatusPersistenceGateway: paymentStatusPersistanceGateway}

	})
	return processPaymentStatusUseCase
}

func mapperPaymentStatusCommandToEntityPaymentStatus(command command.CreatePaymentStatusCommand) domain.PaymentStatus {
	return domain.PaymentStatus{
		Id:                   command.Id,
		PaymentId:            command.PaymentId,
		PaymentIntegrationId: command.ExternalReference,
		PaymentStatus:        command.Status,
	}
}
