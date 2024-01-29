package usecase

import (
	"context"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/port/input"
	"hamburgueria/internal/modules/order/usecase/result"
	paymentInput "hamburgueria/internal/modules/payment/port/input"
	"hamburgueria/internal/modules/payment/usecase/command"
	"sync"
)

type ProcessPaymentUseCase struct {
	updateOrderUseCase   input.UpdateOrderPort
	createPaymentUseCase paymentInput.CreatePaymentPort
}

func (p ProcessPaymentUseCase) ProcessPayment(ctx context.Context, order domain.Order) (*result.PaymentCreatedResult, error) {

	var orderItems []command.OrderItem
	for _, orderProduct := range order.Products {
		orderItems = append(orderItems, command.OrderItem{
			Name:        orderProduct.Product.Name,
			Amount:      orderProduct.Product.Amount,
			Quantity:    orderProduct.Quantity,
			TotalAmount: orderProduct.Amount,
		})
	}
	paymentData, err := p.createPaymentUseCase.CreatePayment(ctx, command.CreatePaymentCommand{
		Amount:     order.Amount,
		OrderId:    order.Id,
		OrderItems: orderItems,
	})
	if err != nil {
		return nil, err
	}

	err = p.updateOrderUseCase.Update(ctx, order.Id, valueobject.PaymentCreated, &paymentData.PaymentId)
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
	updateOrderUseCase input.UpdateOrderPort,
	processPaymentUseCase paymentInput.CreatePaymentPort,
) input.ProcessPaymentPort {
	processPaymentUseCaseOnce.Do(func() {
		processPaymentUseCaseInstance = ProcessPaymentUseCase{
			updateOrderUseCase:   updateOrderUseCase,
			createPaymentUseCase: processPaymentUseCase,
		}
	})
	return processPaymentUseCaseInstance
}
