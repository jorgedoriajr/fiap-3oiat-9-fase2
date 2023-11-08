package usecase

import (
	"context"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/result"
	productResult "hamburgueria/internal/modules/product/usecase/result"
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
		var products []productResult.FindProductResult
		for _, product := range order.Products {
			products = append(products, productResult.FromProductDomain(product.Product))
		}
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
		var products []productResult.FindProductResult
		for _, product := range order.Products {
			products = append(products, productResult.FromProductDomain(product.Product))
		}
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
) ListOrderUseCase {
	listOrderUseCaseOnce.Do(func() {
		listOrderUseCaseInstance = ListOrderUseCase{
			orderPersistence: orderPersistence,
		}
	})
	return listOrderUseCaseInstance
}
