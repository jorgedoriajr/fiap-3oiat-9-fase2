package usecase

import (
	"context"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/result"
	"hamburgueria/internal/modules/product/ports/input"
	"sync"
)

type ListOrderUseCase struct {
	orderPersistence output.OrderPersistencePort
	productFinder    input.ProductFinderServicePort
}

func (c ListOrderUseCase) FindAllOrders(ctx context.Context) ([]result.ListOrderResult, error) {
	orders, err := c.orderPersistence.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var resultOrders []result.ListOrderResult
	for _, order := range orders {
		products, err := c.productFinder.FindByOrderID(ctx, order.Id)
		if err != nil {
			return nil, err
		}
		resultOrders = append(resultOrders, result.ListOrderResult{
			OrderId:    order.Id,
			Status:     string(order.Status),
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
			Products:   products,
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
		products, err := c.productFinder.FindByOrderID(ctx, order.Id)
		if err != nil {
			return nil, err
		}
		resultOrders = append(resultOrders, result.ListOrderResult{
			OrderId:    order.Id,
			Status:     string(order.Status),
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
			Products:   products,
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
	productFinder input.ProductFinderServicePort,
) ListOrderUseCase {
	listOrderUseCaseOnce.Do(func() {
		listOrderUseCaseInstance = ListOrderUseCase{
			orderPersistence: orderPersistence,
			productFinder:    productFinder,
		}
	})
	return listOrderUseCaseInstance
}
