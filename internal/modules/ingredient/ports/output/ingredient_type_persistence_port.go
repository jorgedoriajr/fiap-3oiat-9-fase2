package output

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain"
)

type IngredientTypePersistencePort interface {
	GetByName(ctx context.Context, name string) (*domain.IngredientType, error)
	GetAll(ctx context.Context) ([]domain.IngredientType, error)
}
