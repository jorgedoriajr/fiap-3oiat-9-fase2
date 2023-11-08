package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	customerOutput "hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
	"hamburgueria/internal/modules/payment/port/input"
	productOutput "hamburgueria/internal/modules/product/ports/output"
	"sync"
	"time"
)

type CreateOrderUseCase struct {
	customerPersistence   customerOutput.CustomerPersistencePort
	productPersistence    productOutput.ProductPersistencePort
	orderPersistence      output.OrderPersistencePort
	processPaymentUseCase input.ProcessPaymentPort
}

func (c CreateOrderUseCase) AddOrder(
	ctx context.Context,
	createOrderCommand command.CreateOrderCommand,
) (*result.CreateOrderResult, error) {

	customer, err := c.customerPersistence.Get(ctx, createOrderCommand.CustomerDocument)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	var amount int
	var products []domain.OrderProduct
	orderId := uuid.New()

	for _, createProductCommand := range createOrderCommand.Products {
		var productAmount int
		if createProductCommand.Type == "default" {
			product, err := c.productPersistence.GetByNumber(ctx, createProductCommand.Number)
			if err != nil {
				return nil, err
			}
			if product == nil {
				return nil, errors.New(fmt.Sprintf("product %d not found", createProductCommand.Number))
			}
			productAmount = product.Amount * createProductCommand.Quantity
			products = append(products, domain.OrderProduct{
				Id:       uuid.New(),
				Product:  *product,
				OrderId:  orderId,
				Quantity: createProductCommand.Quantity,
				Amount:   productAmount,
			})
		} else {
			return nil, errors.New("not implemented")
		}

		amount = amount + productAmount

	}

	order := domain.Order{
		Id:         orderId,
		CustomerId: createOrderCommand.CustomerDocument,
		Products:   products,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     valueobject.Created,
		Amount:     amount,
	}

	err = c.orderPersistence.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	paymentProcessed, err := c.processPaymentUseCase.ProcessPayment(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return &result.CreateOrderResult{
		Amount:      amount,
		PaymentData: paymentProcessed.PaymentData,
	}, err
}

var (
	createOrderUseCaseInstance CreateOrderUseCase
	createOrderUseCaseOnce     sync.Once
)

func GetCreateOrderUseCase(
	productPersistence productOutput.ProductPersistencePort,
	orderPersistence output.OrderPersistencePort,
	processPaymentUseCase input.ProcessPaymentPort,
	customerPersistence customerOutput.CustomerPersistencePort,
) CreateOrderUseCase {
	createOrderUseCaseOnce.Do(func() {
		createOrderUseCaseInstance = CreateOrderUseCase{
			productPersistence:    productPersistence,
			orderPersistence:      orderPersistence,
			processPaymentUseCase: processPaymentUseCase,
			customerPersistence:   customerPersistence,
		}
	})
	return createOrderUseCaseInstance
}
