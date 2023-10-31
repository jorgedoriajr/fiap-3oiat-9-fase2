package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
	"hamburgueria/internal/modules/product/service"
	"sync"
	"time"
)

type CreateOrderUseCase struct {
	productFinderService    service.ProductFinderService
	orderPersistence        output.OrderPersistencePort
	orderHistoryPersistence output.OrderHistoryPersistencePort
	orderProductPersistence output.OrderProductPersistencePort
}

func (c CreateOrderUseCase) AddOrder(
	ctx context.Context,
	createOrderCommand command.CreateOrderCommand,
) (*result.CreateOrderResult, error) {

	var amount int
	var products []entity.OrderProduct
	orderId := uuid.New()

	for _, createProductCommand := range createOrderCommand.Products {
		var productAmount int
		if createProductCommand.Type == "default" {
			product, err := c.productFinderService.FindByID(ctx, createProductCommand.Id)
			if err != nil {
				return nil, err
			}
			productAmount = product.Amount * createProductCommand.Quantity
			products = append(products, entity.OrderProduct{
				Id:        uuid.New(),
				ProductId: product.ID,
				OrderId:   orderId,
				Quantity:  createProductCommand.Quantity,
				Amount:    productAmount,
			})
		} else {
			return nil, errors.New("not implemented")
		}

		amount = amount + productAmount

	}

	order := entity.Order{
		Id:         orderId,
		CustomerId: createOrderCommand.CustomerDocument,
		Products:   products,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "CREATED",
		Amount:     amount,
	}

	err := c.createOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &result.CreateOrderResult{
		Amount:      amount,
		PaymentData: "not implemented",
	}, err
}

// TODO need to create transaction
func (c CreateOrderUseCase) createOrder(ctx context.Context, order entity.Order) error {
	err := c.orderPersistence.Create(ctx, order)
	if err != nil {
		return err
	}

	err = c.createOrderHistory(ctx, order)
	if err != nil {
		return err
	}

	err = c.createOrderProducts(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (c CreateOrderUseCase) createOrderProducts(ctx context.Context, order entity.Order) error {
	for _, product := range order.Products {
		err := c.orderProductPersistence.Create(ctx, entity.OrderProduct{
			Id:        uuid.New(),
			ProductId: product.ProductId,
			OrderId:   order.Id,
			Quantity:  product.Quantity,
			Amount:    product.Amount,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c CreateOrderUseCase) createOrderHistory(ctx context.Context, order entity.Order) error {
	return c.orderHistoryPersistence.Create(ctx, entity.OrderHistory{
		Id:        uuid.New(),
		OrderId:   order.Id,
		Status:    order.Status,
		ChangeBy:  "user",
		CreatedAt: order.CreatedAt,
	})
}

var (
	createOrderUseCaseInstance CreateOrderUseCase
	createOrderUseCaseOnce     sync.Once
)

func GetCreateOrderUseCase(
	productFinderService service.ProductFinderService,
	orderPersistence output.OrderPersistencePort,
	orderHistoryPersistence output.OrderHistoryPersistencePort,
	orderProductPersistence output.OrderProductPersistencePort,
) CreateOrderUseCase {
	createOrderUseCaseOnce.Do(func() {
		createOrderUseCaseInstance = CreateOrderUseCase{
			productFinderService:    productFinderService,
			orderPersistence:        orderPersistence,
			orderHistoryPersistence: orderHistoryPersistence,
			orderProductPersistence: orderProductPersistence,
		}
	})
	return createOrderUseCaseInstance
}
