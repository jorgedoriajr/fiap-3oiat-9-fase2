package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductCategoryPersistencePort interface {
	GetAllProductCategories(ctx context.Context) ([]entity.ProductCategoryEntity, error)
	GetProductCategoryByName(ctx context.Context, category string) (*entity.ProductCategoryEntity, error)
}
