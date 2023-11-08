package postgres

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/pkg/sql"

	"github.com/rs/zerolog"
)

type IngredientTypeRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c IngredientTypeRepository) GetTypeByName(ctx context.Context, name string) (*entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeByName, name).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("ingredientType", name).
			Msg("Failed to get ingredient")
		return nil, err
	}

	if result.Name == "" {
		return nil, nil
	}
	ingredientTypeResponse := result.ToEntity()
	return &ingredientTypeResponse, nil
}

func (c IngredientTypeRepository) GetByProductCategory(ctx context.Context, productCategoryName string) ([]entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeByProductCategory, productCategoryName).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get all ingredient")
		return nil, err
	}

	return read.ToIngredientTypeEntityList(result), nil
}

func (c IngredientTypeRepository) GetAll(ctx context.Context) ([]entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeAll).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get all ingredient")
		return nil, err
	}

	return read.ToIngredientTypeEntityList(result), nil
}

func NewIngredientTypeRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *IngredientTypeRepository {
	return &IngredientTypeRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
