package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type IngredientPersistencePort interface {
	Create(ctx context.Context, ingredient entity.IngredientEntity) error
	GetAll(ctx context.Context) ([]entity.IngredientEntity, error)
	GetByID(ctx context.Context, ingredientID uuid.UUID) (*entity.IngredientEntity, error)
	GetByType(ctx context.Context, ingredientType valueobject.IngredientType) ([]entity.IngredientEntity, error)
}
