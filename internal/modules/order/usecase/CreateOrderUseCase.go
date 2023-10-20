package usecase

import (
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/internal/modules/payment/port/input"
)

type CreateOrderUseCase struct {
	orderPersistence output.OrderPersistencePort
	paymentGateway   input.CreatePaymentPort
}

func (c CreateOrderUseCase) CreateOrder(createOrderCommand any) (createOrderResult any, err error) {
	// TODO validate command
	// TODO persistence new order
	// TODO make payment

	_, err = c.orderPersistence.Create("persistence cmd")
	if err != nil {
		return nil, err
	}
	_, err = c.paymentGateway.CreatePayment("payment cmd")
	if err != nil {
		return nil, err
	}
	panic("implement me")
}
