package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientUseCase interface {
	AddIngredient(ctx context.Context, ingredient entity.Ingredient) error
}
