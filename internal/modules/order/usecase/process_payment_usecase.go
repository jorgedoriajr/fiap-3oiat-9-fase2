package usecase

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/result"
	paymentInput "hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/usecase/command"
	"sync"
	"time"
)

type ProcessPaymentUseCase struct {
	orderPersistenceGateway output.OrderPersistencePort
	createPaymentUseCase    paymentInput.CreatePaymentPort
}

func (p ProcessPaymentUseCase) ProcessPayment(ctx context.Context, order domain.Order) (*result.PaymentCreatedResult, error) {

	paymentData, err := p.createPaymentUseCase.CreatePayment(ctx, command.CreatePaymentCommand{Amount: order.Amount})
	if err != nil {
		return nil, err
	}

	order.PaymentId = paymentData.PaymentId
	order.Status = valueobject.PaymentCreated
	order.UpdatedAt = time.Now()
	order.History = append(order.History, domain.OrderHistory{
		Id:        uuid.New(),
		OrderId:   order.Id,
		Status:    order.Status,
		ChangeBy:  "SYSTEM",
		CreatedAt: order.UpdatedAt,
	})

	err = p.orderPersistenceGateway.Update(ctx, order)
	if err != nil {
		return nil, err
	}
	return &result.PaymentCreatedResult{PaymentData: paymentData.PaymentData}, nil
}

var (
	processPaymentUseCaseInstance input.ProcessPaymentPort
	processPaymentUseCaseOnce     sync.Once
)

func GetProcessPaymentUseCase(
	orderPersistenceGateway output.OrderPersistencePort,
	processPaymentUseCase paymentInput.CreatePaymentPort,
) input.ProcessPaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCaseInstance = ProcessPaymentUseCase{
			orderPersistenceGateway: orderPersistenceGateway,
			createPaymentUseCase:    processPaymentUseCase,
		}
	})
	return processPaymentUseCaseInstance
}
