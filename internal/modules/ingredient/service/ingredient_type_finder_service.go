package service

import (
	"context"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"sync"
)

var (
	ingredientTypeFinderServiceInstance IngredientTypeFinderService
	ingredientTypeFinderServiceOnce     sync.Once
)

type IngredientTypeFinderService struct {
	ingredientTypePersistence output.IngredientTypePersistencePort
}

func (p IngredientTypeFinderService) FindAllIngredientType(ctx context.Context) ([]result.IngredientTypeResult, error) {
	ingredients, err := p.ingredientTypePersistence.GetAll(ctx)
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

func GetIngredientTypeFinderService(IngredientTypePersistence output.IngredientTypePersistencePort) IngredientTypeFinderService {
	ingredientTypeFinderServiceOnce.Do(func() {
		ingredientTypeFinderServiceInstance = IngredientTypeFinderService{
			ingredientTypePersistence: IngredientTypePersistence,
		}
	})
	return ingredientTypeFinderServiceInstance
}
