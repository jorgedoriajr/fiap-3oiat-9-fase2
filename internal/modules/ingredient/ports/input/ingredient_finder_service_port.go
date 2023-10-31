package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type IngredientFinderServicePort interface {
	FindIngredientByID(ctx context.Context, id uuid.UUID) (*result.FindIngredientResult, error)

	FindIngredientsByProductId(ctx context.Context, productID uuid.UUID) ([]read.FindIngredientQueryResult, error)

	FindIngredientByType(ctx context.Context, ingredientType string) ([]result.FindIngredientResult, error)

	FindAllIngredients(ctx context.Context) ([]result.FindIngredientResult, error)
}
