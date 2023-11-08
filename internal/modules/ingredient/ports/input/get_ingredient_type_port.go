package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type IngredientTypeFinderServicePort interface {
	FindAllIngredientType(ctx context.Context) ([]result.IngredientTypeResult, error)
}
