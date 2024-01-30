package usecase

import (
	"context"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/result"
	"sync"
)

type ListOrderUseCase struct {
	orderPersistenceGateway output.OrderPersistencePort
}

func (c ListOrderUseCase) FindAllOrders(ctx context.Context) ([]result.ListOrderResult, error) {
	orders, err := c.orderPersistenceGateway.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var resultOrders []result.ListOrderResult
	for _, order := range orders {
		resultOrders = append(resultOrders, *result.ListOrderResultFromDomain(order))
	}
	return resultOrders, nil
}

func (c ListOrderUseCase) FindByStatus(ctx context.Context, status string) ([]result.ListOrderResult, error) {
	orders, err := c.orderPersistenceGateway.FindByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var resultOrders []result.ListOrderResult
	for _, order := range orders {
		resultOrders = append(resultOrders, *result.ListOrderResultFromDomain(order))
	}
	return resultOrders, nil
}

func (c ListOrderUseCase) FindByNumber(ctx context.Context, number int) (*result.ListOrderResult, error) {
	order, err := c.orderPersistenceGateway.FindByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return result.ListOrderResultFromDomain(*order), nil
}

var (
	listOrderUseCaseInstance input.ListOrderPort
	listOrderUseCaseOnce     sync.Once
)

func GetListOrderUseCase(
	orderPersistenceGateway output.OrderPersistencePort,
) input.ListOrderPort {
	listOrderUseCaseOnce.Do(func() {
		listOrderUseCaseInstance = ListOrderUseCase{
			orderPersistenceGateway: orderPersistenceGateway,
		}
	})
	return listOrderUseCaseInstance
}
