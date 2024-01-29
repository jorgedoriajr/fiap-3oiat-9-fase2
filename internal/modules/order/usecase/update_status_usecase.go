package usecase

import (
	"context"
	"errors"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/port/output"
	"sync"
	"time"

	"github.com/google/uuid"
)

type UpdateOrderUseCase struct {
	orderPersistenceGateway output.OrderPersistencePort
}

func (p UpdateOrderUseCase) Update(
	ctx context.Context,
	orderId uuid.UUID,
	status valueobject.OrderStatus,
	paymentId *uuid.UUID,
) error {
	order, err := p.orderPersistenceGateway.FindById(ctx, orderId)

	if err != nil {
		return err
	}

	if order == nil {
		return errors.New("order not found")
	}

	if paymentId != nil {
		order.PaymentId = *paymentId
	}

	order.Status = status
	order.UpdatedAt = time.Now()
	order.History = append(order.History, domain.OrderHistory{
		Id:        uuid.New(),
		OrderId:   order.Id,
		Status:    order.Status,
		ChangeBy:  "SYSTEM",
		CreatedAt: order.UpdatedAt,
	})

	return p.orderPersistenceGateway.Update(ctx, *order)
}

var (
	updateStatusUseCaseInstance input.UpdateOrderPort
	updateStatusUseCaseOnce     sync.Once
)

func GetUpdateOrderUseCase(
	orderPersistenceGateway output.OrderPersistencePort,
) input.UpdateOrderPort {
	updateStatusUseCaseOnce.Do(func() {
		updateStatusUseCaseInstance = UpdateOrderUseCase{
			orderPersistenceGateway: orderPersistenceGateway,
		}
	})
	return updateStatusUseCaseInstance
}
