package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain"
)

type ProductCategoryPersistencePort interface {
	GetAll(ctx context.Context) ([]domain.ProductCategory, error)
	GetByName(ctx context.Context, category string) (*domain.ProductCategory, error)
}
