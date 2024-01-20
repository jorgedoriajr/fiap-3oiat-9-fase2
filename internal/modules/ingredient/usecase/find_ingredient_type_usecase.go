package usecase

import (
	"context"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"sync"
)

var (
	findIngredientTypeUseCaseInstance input.FindIngredientTypeUseCasePort
	findIngredientTypeUseCaseOnce     sync.Once
)

type FindIngredientTypeUseCase struct {
	ingredientTypePersistenceGateway output.IngredientTypePersistencePort
}

func (p FindIngredientTypeUseCase) FindAll(ctx context.Context) ([]result.IngredientTypeResult, error) {
	ingredients, err := p.ingredientTypePersistenceGateway.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var ingredientTypeResult []result.IngredientTypeResult
	for _, ingredientType := range ingredients {
		ingredientTypeResult = append(ingredientTypeResult, result.IngredientTypeResult{
			Name: ingredientType.Name,
		})
	}
	return ingredientTypeResult, nil
}

func GetIngredientTypeUseCase(
	IngredientTypePersistenceGateway output.IngredientTypePersistencePort,
) input.FindIngredientTypeUseCasePort {
	findIngredientTypeUseCaseOnce.Do(func() {
		findIngredientTypeUseCaseInstance = FindIngredientTypeUseCase{
			ingredientTypePersistenceGateway: IngredientTypePersistenceGateway,
		}
	})
	return findIngredientTypeUseCaseInstance
}
