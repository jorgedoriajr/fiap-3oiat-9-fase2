package output

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"

	"github.com/google/uuid"
)

type PaymentPersistencePort interface {
	Create(ctx context.Context, payment domain.Payment) error
	FindById(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error)
}
