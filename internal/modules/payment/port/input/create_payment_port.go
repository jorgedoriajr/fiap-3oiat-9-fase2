package input

import (
	"context"
	"hamburgueria/internal/modules/payment/usecase/command"
	"hamburgueria/internal/modules/payment/usecase/result"
)

type CreatePaymentPort interface {
	CreatePayment(ctx context.Context, command command.CreatePaymentCommand) (*result.PaymentProcessed, error)
}
