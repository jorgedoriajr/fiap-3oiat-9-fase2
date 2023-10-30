package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductCategoryPersistencePort interface {
	GetAllProductCategories(ctx context.Context) ([]entity.ProductCategoryEntity, error)
	GetProductCategoryByName(ctx context.Context, category string) (*entity.ProductCategoryEntity, error)
	GetProductCategoryById(ctx context.Context, cateroryId uuid.UUID) (*entity.ProductCategoryEntity, error)
}
