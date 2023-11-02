package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/payment/usecase/result"
)

type ProcessPaymentPort interface {
	ProcessPayment(ctx context.Context, orderReference uuid.UUID) (*result.PaymentProcessed, error)
}
