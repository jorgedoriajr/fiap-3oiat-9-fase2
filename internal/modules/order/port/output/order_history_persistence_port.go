package output

import (
	"context"
	"hamburgueria/internal/modules/order/domain"
)

type OrderHistoryPersistencePort interface {
	Create(ctx context.Context, order domain.OrderHistory) error
}
