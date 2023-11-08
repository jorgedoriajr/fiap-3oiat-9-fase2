package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type FindIngredientUseCasePort interface {
	FindIngredientByNumber(ctx context.Context, number int) (*result.FindIngredientResult, error)

	FindIngredientByType(ctx context.Context, ingredientType string) ([]result.FindIngredientResult, error)

	FindAllIngredients(ctx context.Context) ([]result.FindIngredientResult, error)
}
