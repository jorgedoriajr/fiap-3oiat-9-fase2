package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type IngredientPersistencePort interface {
	Create(ctx context.Context, ingredient domain.Ingredient) error
	GetAll(ctx context.Context) ([]domain.Ingredient, error)
	GetByID(ctx context.Context, ingredientID uuid.UUID) (*domain.Ingredient, error)
	GetByType(ctx context.Context, ingredientType string) ([]domain.Ingredient, error)
	GetByNumber(ctx context.Context, number int) (*domain.Ingredient, error)
}
