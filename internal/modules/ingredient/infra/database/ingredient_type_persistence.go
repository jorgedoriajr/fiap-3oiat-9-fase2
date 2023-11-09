package database

import (
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/ingredient/infra/database/model"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"sync"
)

type IngredientTypeRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c IngredientTypeRepository) GetByName(ctx context.Context, name string) (*domain.IngredientType, error) {
	var ingredientType model.IngredientType
	err := c.readOnlyClient.Preload("ConfigByProductCategory").Where("name = ?", name).Find(&ingredientType).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("name", name).
			Msg("Failed to find type by name")
		return nil, err
	}

	return ingredientType.ToDomain(), nil
}

func (c IngredientTypeRepository) GetAll(ctx context.Context) ([]domain.IngredientType, error) {
	var ingredientTypes []model.IngredientType
	err := c.readOnlyClient.Preload("ConfigByProductCategory").Find(&ingredientTypes).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Msg("Failed to find all types")
		return nil, err
	}

	var types []domain.IngredientType
	for _, ingredientType := range ingredientTypes {
		types = append(types, *ingredientType.ToDomain())
	}

	return types, nil
}

var (
	ingredientTypeInstance output.IngredientTypePersistencePort
	ingredientTypeOnce     sync.Once
)

func GetIngredientTypeRepository(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.IngredientTypePersistencePort {
	ingredientTypeOnce.Do(func() {
		ingredientTypeInstance = IngredientTypeRepository{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return ingredientTypeInstance
}
