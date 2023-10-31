package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductIngredientPersistencePort interface {
	Create(ctx context.Context, productIngredient entity.ProductIngredientEntity) error
}
