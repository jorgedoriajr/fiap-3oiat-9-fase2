package output

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientTypePersistencePort interface {
	Create(context.Context, entity.IngredientType) error
	GetByName(id string) (context.Context, entity.IngredientType, error)
}
