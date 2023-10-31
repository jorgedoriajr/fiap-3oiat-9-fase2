package output

import (
	"context"
	"hamburgueria/internal/modules/order/domain/entity"
)

type OrderHistoryPersistencePort interface {
	Create(ctx context.Context, order entity.OrderHistory) error
}
