package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductPersistencePort interface {
	Create(ctx context.Context, product entity.ProductEntity) error
	GetAll(ctx context.Context) ([]entity.ProductEntity, error)
	GetByID(ctx context.Context, productID int) (*entity.ProductEntity, error)
	GetByCategory(ctx context.Context, productID string) ([]entity.ProductEntity, error)
}

// produto por ingredient
