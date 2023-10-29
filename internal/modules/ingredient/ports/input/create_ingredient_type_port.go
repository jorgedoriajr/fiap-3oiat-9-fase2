package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientTypePort interface {
	AddIngredientType(ctx context.Context, ingredientType entity.IngredientType) error
}
