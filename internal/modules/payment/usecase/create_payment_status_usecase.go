package usecase

import (
	"context"
	voOrder "hamburgueria/internal/modules/order/domain/valueobject"
	inputOrder "hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/domain/valueobjects"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/port/output"

	"hamburgueria/internal/modules/payment/usecase/command"

	"sync"

	"github.com/rs/zerolog"
)

type CreatePaymentStatusUseCase struct {
	paymentStatusPersistenceGateway output.PaymentStatusPersistencePort
	updateOrderStatus               inputOrder.UpdateOrderPort
	logger                          zerolog.Logger
}

func (ps CreatePaymentStatusUseCase) AddPaymentStatus(ctx context.Context, command command.CreatePaymentStatusCommand) error {
	paymentStatus := mapperPaymentStatusCommandToEntityPaymentStatus(command)

	switch paymentStatus.PaymentStatus {
	case valueobjects.Approved:
		ps.updateOrderStatus.Update(ctx, command.ExternalReference, voOrder.Started, nil)
	case valueobjects.Rejected:
		ps.updateOrderStatus.Update(ctx, command.ExternalReference, voOrder.PaymentRefused, nil)
	default:
		ps.logger.Info().Ctx(ctx).Str("Status", string(paymentStatus.PaymentStatus)).Msg("Status ignorado!")
	}
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

func GetCreatePaymentStatusUseCase(paymentStatusPersistanceGateway output.PaymentStatusPersistencePort, updateOrderStatus inputOrder.UpdateOrderPort, logger zerolog.Logger) input.CreatePaymentStatusPort {
	processPaymentStatusUseCaseOnce.Do(func() {
		processPaymentStatusUseCase = CreatePaymentStatusUseCase{paymentStatusPersistenceGateway: paymentStatusPersistanceGateway, updateOrderStatus: updateOrderStatus, logger: logger}
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
