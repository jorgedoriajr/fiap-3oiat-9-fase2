package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductPersistencePort interface {
	Create(ctx context.Context, product entity.ProductEntity) error
	GetByID(ctx context.Context, productID string) (entity.ProductEntity, error)
}
