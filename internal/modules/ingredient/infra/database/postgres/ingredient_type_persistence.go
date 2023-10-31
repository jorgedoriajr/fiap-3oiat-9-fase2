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

func (c IngredientTypeRepository) GetByName(ctx context.Context, ingredientTypeName string) (*entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeByName, ingredientTypeName).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("ingredientType", ingredientTypeName).
			Msg("Failed to get ingredient")
		return nil, err
	}

	return result.ToEntity(), nil
}

func (c IngredientTypeRepository) GetByProductCategory(ctx context.Context, productCategoryName string) ([]*entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeByProductCategory, productCategoryName).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get all ingredient")
		return nil, err
	}

	return read.ToIngredientTypeEntityList(result), nil
}

func (c IngredientTypeRepository) GetAll(ctx context.Context) ([]*entity.IngredientType, error) {

	result, err := sql.NewQuery[read.FindIngredientTypeQueryResult](ctx, c.readOnlyClient, read.FindIngredientTypeAll).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get all ingredient")
		return nil, err
	}

	return read.ToIngredientTypeEntityList(result), nil
}
