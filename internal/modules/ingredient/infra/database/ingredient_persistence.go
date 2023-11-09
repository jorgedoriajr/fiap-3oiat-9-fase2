package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/ingredient/infra/database/model"
	"hamburgueria/internal/modules/ingredient/ports/output"
	"strconv"
	"sync"
)

type IngredientRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c IngredientRepository) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
	var ingredients []model.Ingredient
	err := c.readOnlyClient.Preload(clause.Associations).Find(&ingredients).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Msg("Failed to find all ingredient")
		return nil, err
	}

	var domainIngredients []domain.Ingredient
	for _, ingredient := range ingredients {
		domainIngredients = append(domainIngredients, *ingredient.ToDomain())
	}

	return domainIngredients, nil
}

func (c IngredientRepository) GetByType(ctx context.Context, ingredientType string) ([]domain.Ingredient, error) {
	var ingredients []model.Ingredient
	err := c.readOnlyClient.Preload(clause.Associations).Joins("Type").Where("type.name = ?", ingredientType).Find(&ingredients).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("type", ingredientType).
			Msg("Failed to find ingredient by type")
		return nil, err
	}

	var domainIngredients []domain.Ingredient
	for _, ingredient := range ingredients {
		domainIngredients = append(domainIngredients, *ingredient.ToDomain())
	}

	return domainIngredients, nil
}

func (c IngredientRepository) Create(ctx context.Context, ingredient domain.Ingredient) error {
	err := c.readWriteClient.
		Omit("Type").
		Create(&model.Ingredient{
			ID:     ingredient.ID,
			Number: ingredient.Number,
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
			Type: model.IngredientType{
				Name: ingredient.Name,
			},
		}).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("ingredient", ingredient.Name).
			Msg("Failed to insert ingredient")
		return err
	}
	return nil
}

func (c IngredientRepository) GetByID(ctx context.Context, ingredientId uuid.UUID) (*domain.Ingredient, error) {
	var ingredient model.Ingredient
	err := c.readOnlyClient.Preload(clause.Associations).First(&ingredient, ingredientId).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("ingredientId", ingredientId.String()).
			Msg("Failed to find ingredient by id")
		return nil, err
	}

	return ingredient.ToDomain(), nil
}

func (c IngredientRepository) GetByNumber(ctx context.Context, number int) (*domain.Ingredient, error) {
	var ingredient model.Ingredient
	err := c.readOnlyClient.Preload(clause.Associations).Where("number = ?", number).First(&ingredient).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("number", strconv.Itoa(number)).
			Msg("Failed to find ingredient by number")
		return nil, err
	}

	return ingredient.ToDomain(), nil
}

var (
	ingredientRepositoryInstance output.IngredientPersistencePort
	ingredientRepositoryOnce     sync.Once
)

func GetIngredientRepository(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.IngredientPersistencePort {
	ingredientRepositoryOnce.Do(func() {
		ingredientRepositoryInstance = IngredientRepository{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return ingredientRepositoryInstance
}
