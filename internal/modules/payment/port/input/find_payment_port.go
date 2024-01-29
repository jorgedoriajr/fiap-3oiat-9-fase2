package input

import (
	"context"
	"hamburgueria/internal/modules/payment/usecase/result"

	"github.com/google/uuid"
)

type FindPaymentPort interface {
	FindPaymentById(ctx context.Context, paymentId uuid.UUID) (*result.PaymentProcessed, error)
}
