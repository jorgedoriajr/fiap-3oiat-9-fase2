package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/payment/usecase/result"
)

type OrderPersistencePort interface {
	Create(ctx context.Context, order entity.Order) error
	FindAll(ctx context.Context) ([]entity.Order, error)
	FindByStatus(ctx context.Context, status string) ([]entity.Order, error)
	FindById(ctx context.Context, orderId uuid.UUID) (*entity.Order, error)
	SavePaymentReference(ctx context.Context, payment result.PaymentProcessed) error
}
