package input

import (
	"context"
	"hamburgueria/internal/modules/payment/usecase/result"

	"github.com/google/uuid"
)

type ListPaymentStatusPort interface {
	ListPaymentStatus(ctx context.Context, paymentStatusId uuid.UUID) (*result.PaymentStatusProcessed, error)
}
