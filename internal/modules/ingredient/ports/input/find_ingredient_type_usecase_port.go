package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type FindIngredientTypeUseCasePort interface {
	FindAll(ctx context.Context) ([]result.IngredientTypeResult, error)
}
