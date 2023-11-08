package output

import (
	"context"
	"hamburgueria/internal/modules/order/domain"
)

type OrderProductPersistencePort interface {
	Create(ctx context.Context, order domain.OrderProduct) error
}
