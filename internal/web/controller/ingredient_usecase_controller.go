package controller

import (
	"gorm.io/gorm"
	ingredientDatabase "hamburgueria/internal/modules/ingredient/infra/database"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/usecase"
	"hamburgueria/pkg/logger"
)

type IngredientUseCaseController struct {
	CreateIngredientUseCase   input.CreateIngredientUseCasePort
	FindIngredientTypeUseCase input.FindIngredientTypeUseCasePort
	FindIngredientUseCase     input.FindIngredientUseCasePort
}

func NewIngredientUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *IngredientUseCaseController {
	ingredientTypePersistence := ingredientDatabase.GetIngredientTypePersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
	ingredientPersistence := ingredientDatabase.GetIngredientPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

	return &IngredientUseCaseController{
		CreateIngredientUseCase:   usecase.NewCreateIngredientUseCase(ingredientPersistence, ingredientTypePersistence),
		FindIngredientTypeUseCase: usecase.GetIngredientTypeUseCase(ingredientTypePersistence),
		FindIngredientUseCase:     usecase.NewFindIngredientUseCase(ingredientPersistence),
	}
}
