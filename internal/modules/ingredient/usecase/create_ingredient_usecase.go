package usecase

import (
	"context"
	"fmt"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/ingredient/usecase/command"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"sync"
)

var (
	createIngredientUseCaseInstance *CreateIngredientUseCase
	createIngredientUseCaseOnce     sync.Once
)

type CreateIngredientUseCase struct {
	ingredientPersistence output.IngredientPersistencePort
}

func (c CreateIngredientUseCase) AddIngredient(ctx context.Context, command command.CreateIngredientCommand) (result.CreateIngredientResult, error) {
	ingredient := command.ToIngredientEntity()

	fmt.Printf("creating new ingredient: [%v]", ingredient)
	err := c.ingredientPersistence.Create(ctx, *ingredient)
	if err != nil {
		return result.CreateIngredientResult{}, err
	}
	return result.ToCreateIngredientResultFrom(*ingredient), nil
}

func NewCreateIngredientUseCase(ingredientPersistence output.IngredientPersistencePort) *CreateIngredientUseCase {
	createIngredientUseCaseOnce.Do(func() {
		createIngredientUseCaseInstance = &CreateIngredientUseCase{
			ingredientPersistence: ingredientPersistence,
		}
	})
	return createIngredientUseCaseInstance
}
