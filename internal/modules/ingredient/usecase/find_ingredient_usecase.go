package usecase

import (
	"context"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"sync"
)

var (
	findIngredientUseCaseInstance input.FindIngredientUseCasePort
	findIngredientUseCaseOnce     sync.Once
)

type FindIngredientUseCase struct {
	ingredientPersistenceGateway output.IngredientPersistencePort
}

func (f FindIngredientUseCase) FindIngredientByNumber(ctx context.Context, number int) (*result.FindIngredientResult, error) {
	ingredient, err := f.ingredientPersistenceGateway.GetByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	ingredientResult := result.FromDomain(*ingredient)
	return &ingredientResult, nil
}

func (f FindIngredientUseCase) FindIngredientByType(ctx context.Context, ingredientType string) ([]result.FindIngredientResult, error) {
	ingredients, err := f.ingredientPersistenceGateway.GetByType(ctx, ingredientType)
	if err != nil {
		return nil, err
	}

	var ingredientsResult []result.FindIngredientResult
	for _, ingredient := range ingredients {
		ingredientsResult = append(ingredientsResult, result.FromDomain(ingredient))
	}

	return ingredientsResult, nil
}

func (f FindIngredientUseCase) FindAllIngredients(ctx context.Context) ([]result.FindIngredientResult, error) {
	ingredients, err := f.ingredientPersistenceGateway.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var ingredientsResult []result.FindIngredientResult
	for _, ingredient := range ingredients {
		ingredientsResult = append(ingredientsResult, result.FromDomain(ingredient))
	}

	return ingredientsResult, nil
}

func NewFindIngredientUseCase(
	ingredientPersistenceGateway output.IngredientPersistencePort,
) input.FindIngredientUseCasePort {
	findIngredientUseCaseOnce.Do(func() {
		findIngredientUseCaseInstance = FindIngredientUseCase{
			ingredientPersistenceGateway: ingredientPersistenceGateway,
		}
	})
	return findIngredientUseCaseInstance
}
