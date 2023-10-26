package output

import (
	"context"
	"hamburgueria/internal/modules/order/domain/entity"
)

type OrderPersistencePort interface {
	Create(ctx context.Context, order entity.Order) error
}
