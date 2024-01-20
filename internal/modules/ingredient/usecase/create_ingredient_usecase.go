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
	ingredientPersistenceGateway     output.IngredientPersistencePort
	ingredientTypePersistenceGateway output.IngredientTypePersistencePort
}

func (c CreateIngredientUseCase) AddIngredient(
	ctx context.Context,
	command command.CreateIngredientCommand,
) (*result.CreateIngredientResult, error) {
	ingredientType, err := c.ingredientTypePersistenceGateway.GetByName(ctx, command.Type)

	if err != nil {
		return nil, err
	}

	ingredient := command.ToIngredientEntity(*ingredientType)

	fmt.Printf("creating new ingredient: [%v]", ingredient)
	err = c.ingredientPersistenceGateway.Create(ctx, *ingredient)
	if err != nil {
		return nil, err
	}
	return result.ToCreateIngredientResultFrom(*ingredient), nil
}

func NewCreateIngredientUseCase(
	ingredientPersistenceGateway output.IngredientPersistencePort,
	ingredientTypePersistenceGateway output.IngredientTypePersistencePort,
) *CreateIngredientUseCase {
	createIngredientUseCaseOnce.Do(func() {
		createIngredientUseCaseInstance = &CreateIngredientUseCase{
			ingredientPersistenceGateway:     ingredientPersistenceGateway,
			ingredientTypePersistenceGateway: ingredientTypePersistenceGateway,
		}
	})
	return createIngredientUseCaseInstance
}
