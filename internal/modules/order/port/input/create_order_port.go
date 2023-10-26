package input

import (
	"context"
	"hamburgueria/internal/modules/order/usecase/command"
)

type CreateOrderPort interface {
	AddOrder(ctx context.Context, createOrderCommand command.CreateOrderCommand) error
}
