package input

import (
	"context"
	"hamburgueria/internal/modules/order/usecase/command"
	"hamburgueria/internal/modules/order/usecase/result"
)

type CreateOrderPort interface {
	AddOrder(ctx context.Context, createOrderCommand command.CreateOrderCommand) (*result.CreateOrderResult, error)
}
