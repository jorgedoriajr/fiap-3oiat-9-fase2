package output

import (
	"context"
	"hamburgueria/internal/modules/product/domain"
)

type ProductIngredientPersistencePort interface {
	Create(ctx context.Context, productIngredient domain.ProductIngredient) error
}
