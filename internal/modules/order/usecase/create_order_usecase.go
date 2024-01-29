package usecase

import (
	"context"
	"errors"
	"fmt"
	customerOutput "hamburgueria/internal/modules/customer/port/output"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	paymentInput "hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
	productOutput "hamburgueria/internal/modules/product/ports/output"
	"sync"
	"time"

	"github.com/google/uuid"
)

type CreateOrderUseCase struct {
	customerPersistenceGateway customerOutput.CustomerPersistencePort
	productPersistenceGateway  productOutput.ProductPersistencePort
	orderPersistenceGateway    output.OrderPersistencePort
	processPaymentUseCase      paymentInput.ProcessPaymentPort
}

func (c CreateOrderUseCase) AddOrder(
	ctx context.Context,
	createOrderCommand command.CreateOrderCommand,
) (*result.CreateOrderResult, error) {

	customer, err := c.customerPersistenceGateway.Get(ctx, createOrderCommand.CustomerDocument)
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
			product, err := c.productPersistenceGateway.GetByNumber(ctx, createProductCommand.Number)
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

	err = c.orderPersistenceGateway.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	orderFound, err := c.orderPersistenceGateway.FindById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	paymentProcessed, err := c.processPaymentUseCase.ProcessPayment(ctx, *orderFound)
	if err != nil {
		return nil, err
	}

	return &result.CreateOrderResult{
		Number:      orderFound.Number,
		Amount:      amount,
		PaymentData: paymentProcessed.PaymentData,
	}, err
}

var (
	createOrderUseCaseInstance CreateOrderUseCase
	createOrderUseCaseOnce     sync.Once
)

func GetCreateOrderUseCase(
	productPersistenceGateway productOutput.ProductPersistencePort,
	orderPersistenceGateway output.OrderPersistencePort,
	processPaymentUseCase paymentInput.ProcessPaymentPort,
	customerPersistenceGateway customerOutput.CustomerPersistencePort,
) CreateOrderUseCase {
	createOrderUseCaseOnce.Do(func() {
		createOrderUseCaseInstance = CreateOrderUseCase{
			productPersistenceGateway:  productPersistenceGateway,
			orderPersistenceGateway:    orderPersistenceGateway,
			processPaymentUseCase:      processPaymentUseCase,
			customerPersistenceGateway: customerPersistenceGateway,
		}
	})
	return createOrderUseCaseInstance
}
