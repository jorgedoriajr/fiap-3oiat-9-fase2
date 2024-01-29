package input

import (
	"context"
	"hamburgueria/internal/modules/payment/usecase/command"
)

type CreatePaymentStatusPort interface {
	AddPaymentStatus(ctx context.Context, createOrderCommand command.CreatePaymentStatusCommand) error
}
