package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/usecase/result"
	"sync"
	"time"
)

type ProcessPaymentUseCase struct {
	orderPersistence        output.OrderPersistencePort
	orderHistoryPersistence output.OrderHistoryPersistencePort
}

func (p ProcessPaymentUseCase) ProcessPayment(ctx context.Context, orderReference uuid.UUID) (*result.PaymentProcessed, error) {
	//TODO transactional
	orderEntity, err := p.orderPersistence.FindById(ctx, orderReference)

	if orderEntity == nil {
		return nil, errors.New("order not found")
	}

	paymentData := result.PaymentProcessed{
		OrderReference: orderReference,
		PaymentId:      uuid.New(),
		PaymentData:    "not implemented",
	}
	err = p.orderPersistence.SavePaymentReference(ctx, paymentData)
	if err != nil {
		return nil, err
	}
	err = p.orderHistoryPersistence.Create(ctx, entity.OrderHistory{
		Id:        uuid.New(),
		OrderId:   orderEntity.Id,
		Status:    valueobject.PaymentCreated,
		ChangeBy:  "system",
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &paymentData, nil
}

var (
	processPaymentUseCase     ProcessPaymentUseCase
	processPaymentUseCaseOnce sync.Once
)

func GetProcessPaymentUseCase(
	orderPersistence output.OrderPersistencePort,
	orderHistoryPersistence output.OrderHistoryPersistencePort,
) input.ProcessPaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCase = ProcessPaymentUseCase{
			orderPersistence:        orderPersistence,
			orderHistoryPersistence: orderHistoryPersistence,
		}
	})
	return processPaymentUseCase
}
