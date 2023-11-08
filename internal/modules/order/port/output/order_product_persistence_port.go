package output

import (
	"context"
	"hamburgueria/internal/modules/order/domain/entity"
)

type OrderProductPersistencePort interface {
	Create(ctx context.Context, order entity.OrderProduct) error
}
