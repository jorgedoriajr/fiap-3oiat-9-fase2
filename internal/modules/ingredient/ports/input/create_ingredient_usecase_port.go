package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/usecase/command"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type CreateIngredientUseCasePort interface {
	AddIngredient(ctx context.Context, command command.CreateIngredientCommand) (*result.CreateIngredientResult, error)
}
