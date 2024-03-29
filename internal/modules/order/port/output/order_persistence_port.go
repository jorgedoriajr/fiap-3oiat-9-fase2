package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain"
)

type OrderPersistencePort interface {
	Create(ctx context.Context, order domain.Order) error
	FindAll(ctx context.Context) ([]domain.Order, error)
	FindByStatus(ctx context.Context, status string) ([]domain.Order, error)
	FindByNumber(ctx context.Context, number int) (*domain.Order, error)
	FindById(ctx context.Context, orderId uuid.UUID) (*domain.Order, error)
	Update(ctx context.Context, order domain.Order) error
}
