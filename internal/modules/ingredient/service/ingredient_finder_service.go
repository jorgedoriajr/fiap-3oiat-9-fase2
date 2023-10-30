package service

import (
	"context"
	"github.com/google/uuid"
	entity2 "hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"sync"
)

var (
	ingredientFinderServiceInstance *IngredientFinderService
	ingredientFinderServiceOnce     sync.Once
)

type IngredientFinderService struct {
	ingredientPersistence output.IngredientPersistencePort
}

func (p IngredientFinderService) FindAllIngredients(ctx context.Context) ([]entity2.IngredientEntity, error) {
	return p.ingredientPersistence.GetAll(ctx)
}

func (p IngredientFinderService) FindIngredientByType(ctx context.Context, ingredientType valueobject.IngredientType) ([]entity2.IngredientEntity, error) {
	return p.ingredientPersistence.GetByType(ctx, ingredientType)
}

func (p IngredientFinderService) FindIngredientByID(ctx context.Context, id uuid.UUID) (*entity2.IngredientEntity, error) {
	return p.ingredientPersistence.GetByID(ctx, id)
}

func NewIngredientFinderService(ingredientPersistence output.IngredientPersistencePort) *IngredientFinderService {
	ingredientFinderServiceOnce.Do(func() {
		ingredientFinderServiceInstance = &IngredientFinderService{
			ingredientPersistence: ingredientPersistence,
		}
	})
	return ingredientFinderServiceInstance
}
