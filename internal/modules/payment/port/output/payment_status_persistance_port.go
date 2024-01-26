package output

import (
	"context"
	"hamburgueria/internal/modules/payment/domain"

	"github.com/google/uuid"
)

type PaymentStatusPersistencePort interface {
	FindPaymentStatus(ctx context.Context, paymentStatusId uuid.UUID) (*domain.PaymentIntegrationLog, error)
	CreatePaymentStatus(ctx context.Context, paymentStatus domain.PaymentIntegrationLog) error
}
