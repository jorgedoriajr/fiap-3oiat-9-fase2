package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientPort interface {
	AddIngredient(ctx context.Context, ingredient entity.Ingredient) error
}
