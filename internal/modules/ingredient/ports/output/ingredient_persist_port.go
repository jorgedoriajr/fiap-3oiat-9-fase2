package output

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientPersistencePort interface {
	Create(context.Context, entity.Ingredient) error
	GetByID(id string) (context.Context, entity.Ingredient, error)
}
