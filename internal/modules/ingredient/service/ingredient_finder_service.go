package service

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/ingredient/usecase/result"
	"sync"
)

var (
	ingredientFinderServiceInstance *IngredientFinderService
	ingredientFinderServiceOnce     sync.Once
)

type IngredientFinderService struct {
	ingredientPersistence output.IngredientPersistencePort
}

func (p IngredientFinderService) FindAllIngredients(ctx context.Context) ([]result.FindIngredientResult, error) {
	ingredients, err := p.ingredientPersistence.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return toResult(ingredients), nil
}

func (p IngredientFinderService) FindIngredientByType(ctx context.Context, ingredientType string) ([]result.FindIngredientResult, error) {
	ingredients, err := p.ingredientPersistence.GetByType(ctx, ingredientType)
	if err != nil {
		return nil, err
	}
	return toResult(ingredients), nil
}

func (p IngredientFinderService) FindIngredientByID(ctx context.Context, id uuid.UUID) (*result.FindIngredientResult, error) {
	ingredient, err := p.ingredientPersistence.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if ingredient == nil {
		return nil, nil
	}
	findIngredientResult := result.FromEntity(*ingredient)
	return &findIngredientResult, nil
}

func (p IngredientFinderService) FindIngredientsByProductId(ctx context.Context, productID uuid.UUID) ([]read.FindIngredientQueryResult, error) {
	return p.ingredientPersistence.GetByProductID(ctx, productID)
}

func toResult(ingredients []entity.IngredientEntity) []result.FindIngredientResult {
	var results []result.FindIngredientResult
	for _, ingredient := range ingredients {
		results = append(results, result.FromEntity(ingredient))
	}
	return results
}

func NewIngredientFinderService(ingredientPersistence output.IngredientPersistencePort) *IngredientFinderService {
	ingredientFinderServiceOnce.Do(func() {
		ingredientFinderServiceInstance = &IngredientFinderService{
			ingredientPersistence: ingredientPersistence,
		}
	})
	return ingredientFinderServiceInstance
}
