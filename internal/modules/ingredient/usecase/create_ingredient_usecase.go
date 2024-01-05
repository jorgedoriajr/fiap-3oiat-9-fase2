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
	ingredientPersistence     output.IngredientPersistencePort
	ingredientTypePersistence output.IngredientTypePersistencePort
}

func (c CreateIngredientUseCase) AddIngredient(ctx context.Context, command command.CreateIngredientCommand) (*result.CreateIngredientResult, error) {
	ingredientType, err := c.ingredientTypePersistence.GetByName(ctx, command.Type)

	if err != nil {
		return nil, err
	}

	ingredient := command.ToIngredientEntity(*ingredientType)

	fmt.Printf("creating new ingredient: [%v]", ingredient)
	err = c.ingredientPersistence.Create(ctx, *ingredient)
	if err != nil {
		return nil, err
	}
	return result.ToCreateIngredientResultFrom(*ingredient), nil
}

func NewCreateIngredientUseCase(
	ingredientPersistence output.IngredientPersistencePort,
	ingredientTypePersistence output.IngredientTypePersistencePort,
) *CreateIngredientUseCase {
	createIngredientUseCaseOnce.Do(func() {
		createIngredientUseCaseInstance = &CreateIngredientUseCase{
			ingredientPersistence:     ingredientPersistence,
			ingredientTypePersistence: ingredientTypePersistence,
		}
	})
	return createIngredientUseCaseInstance
}
