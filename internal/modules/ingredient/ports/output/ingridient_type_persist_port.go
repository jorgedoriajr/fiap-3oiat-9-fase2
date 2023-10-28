package output

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientTypePersistencePort interface {
	Create(context.Context, entity.IngredientType) error
	GetByName(name string) (context.Context, entity.IngredientType, error)
	GetAll() (context.Context, []entity.IngredientType, error)
}
