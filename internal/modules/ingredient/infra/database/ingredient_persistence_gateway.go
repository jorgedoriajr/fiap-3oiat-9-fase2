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

type IngredientPersistenceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c IngredientPersistenceGateway) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
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

func (c IngredientPersistenceGateway) GetByType(ctx context.Context, ingredientType string) ([]domain.Ingredient, error) {
	var ingredients []model.Ingredient
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Table("ingredient").
		Joins("JOIN ingredient_type ON ingredient_type.name = ingredient.type").
		Where("ingredient.type = ?", ingredientType).
		Find(&ingredients).Error
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

func (c IngredientPersistenceGateway) Create(ctx context.Context, ingredient domain.Ingredient) error {
	err := c.readWriteClient.
		Omit("Type").
		Create(&model.Ingredient{
			ID:     ingredient.ID,
			Number: ingredient.Number,
			Name:   ingredient.Name,
			Amount: ingredient.Amount,
			IngredientType: model.IngredientType{
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

func (c IngredientPersistenceGateway) GetByID(ctx context.Context, ingredientId uuid.UUID) (*domain.Ingredient, error) {
	var ingredient model.Ingredient
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("IngredientType.ConfigByProductCategory").
		First(&ingredient, ingredientId).
		Error
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

func (c IngredientPersistenceGateway) GetByNumber(ctx context.Context, number int) (*domain.Ingredient, error) {
	var ingredient model.Ingredient
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("IngredientType.ConfigByProductCategory").
		Where("number = ?", number).
		First(&ingredient).Error
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
	ingredientPersistenceGatewayInstance output.IngredientPersistencePort
	ingredientPersistenceGatewayOnce     sync.Once
)

func GetIngredientPersistenceGateway(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.IngredientPersistencePort {
	ingredientPersistenceGatewayOnce.Do(func() {
		ingredientPersistenceGatewayInstance = IngredientPersistenceGateway{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return ingredientPersistenceGatewayInstance
}
