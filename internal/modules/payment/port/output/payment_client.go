package output

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/usecase/command"
)

type PaymentClient interface {
	GenerateQrCode(ctx context.Context, command command.CreatePaymentCommand) (*domain.Payment, error)
}
