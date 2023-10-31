package usecase

import (
	"context"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/result"
	"sync"
)

type ListOrderUseCase struct {
	orderPersistence output.OrderPersistencePort
}

func (c ListOrderUseCase) FindAllOrders(ctx context.Context) ([]result.ListOrderResult, error) {
	orders, err := c.orderPersistence.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var resultOrders []result.ListOrderResult
	for _, order := range orders {
		resultOrders = append(resultOrders, result.ListOrderResult{
			OrderId:    order.Id,
			Status:     order.Status,
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
		})
	}
	return resultOrders, nil
}

func (c ListOrderUseCase) FindByStatus(ctx context.Context, status string) ([]result.ListOrderResult, error) {
	orders, err := c.orderPersistence.FindByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var resultOrders []result.ListOrderResult
	for _, order := range orders {
		resultOrders = append(resultOrders, result.ListOrderResult{
			OrderId:    order.Id,
			Status:     order.Status,
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
		})
	}
	return resultOrders, nil
}

var (
	listOrderUseCaseInstance ListOrderUseCase
	listOrderUseCaseOnce     sync.Once
)

func GetListOrderUseCase(
	orderPersistence output.OrderPersistencePort,
) ListOrderUseCase {
	listOrderUseCaseOnce.Do(func() {
		listOrderUseCaseInstance = ListOrderUseCase{
			orderPersistence: orderPersistence,
		}
	})
	return listOrderUseCaseInstance
}
