package output

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type IngredientTypePersistencePort interface {
	GetTypeByName(ctx context.Context, name string) (*entity.IngredientType, error)
	GetAll(ctx context.Context) ([]entity.IngredientType, error)
}
