package controller

import (
	"gorm.io/gorm"
	ingredientDatabase "hamburgueria/internal/modules/ingredient/infra/database"
	"hamburgueria/internal/modules/ingredient/ports/input"
	"hamburgueria/internal/modules/ingredient/usecase"
	"hamburgueria/pkg/logger"
	"sync"
)

type IngredientUseCaseController struct {
	CreateIngredientUseCase   input.CreateIngredientUseCasePort
	FindIngredientTypeUseCase input.FindIngredientTypeUseCasePort
	FindIngredientUseCase     input.FindIngredientUseCasePort
}

var (
	ingredientUseCaseControllerInstance *IngredientUseCaseController
	ingredientUseCaseControllerOnce     sync.Once
)

func GetIngredientUseCaseController(readWriteDB, readOnlyDB *gorm.DB) *IngredientUseCaseController {
	ingredientUseCaseControllerOnce.Do(func() {
		ingredientTypePersistence := ingredientDatabase.GetIngredientTypePersistenceGateway(readWriteDB, readOnlyDB, logger.Get())
		ingredientPersistence := ingredientDatabase.GetIngredientPersistenceGateway(readWriteDB, readOnlyDB, logger.Get())

		ingredientUseCaseControllerInstance = &IngredientUseCaseController{
			CreateIngredientUseCase:   usecase.NewCreateIngredientUseCase(ingredientPersistence, ingredientTypePersistence),
			FindIngredientTypeUseCase: usecase.GetIngredientTypeUseCase(ingredientTypePersistence),
			FindIngredientUseCase:     usecase.NewFindIngredientUseCase(ingredientPersistence),
		}
	})

	return ingredientUseCaseControllerInstance
}
