package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type IngredientRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c IngredientRepository) GetAll(ctx context.Context) ([]entity.IngredientEntity, error) {
	allIngredients, allIngredientsErr := sql.NewQuery[entity.IngredientEntity](
		ctx,
		c.readOnlyClient,
		read.FindAllIngredients,
	).Many()

	if allIngredientsErr != nil {
		c.logger.Error().
			Err(allIngredientsErr).
			Msg("Failed to get ingredients")
		return []entity.IngredientEntity{}, allIngredientsErr
	}

	return allIngredients, nil
}

func (c IngredientRepository) GetByType(ctx context.Context, ingredientType string) ([]entity.IngredientEntity, error) {
	ingredientsByType, ingredientsByTypeErr := sql.NewQuery[entity.IngredientEntity](
		ctx,
		c.readOnlyClient,
		read.FindIngredientsByType,
		ingredientType,
	).Many()

	if ingredientsByTypeErr != nil {
		c.logger.Error().
			Err(ingredientsByTypeErr).
			Str("type", ingredientType).
			Msg("Failed to get ingredients by type")
		return nil, ingredientsByTypeErr
	}
	return ingredientsByType, nil
}

func (c IngredientRepository) Create(ctx context.Context, ingredient entity.IngredientEntity) error {

	mapper := write.ToInsertIngredientQueryMapper(ingredient)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertIngredientRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("name", ingredient.Name).
			Msg("Failed to insert ingredient")
		return err
	}

	return nil
}

func (c IngredientRepository) GetByID(ctx context.Context, ingredientId uuid.UUID) (*entity.IngredientEntity, error) {

	result, err := sql.NewQuery[read.FindIngredientQueryResult](ctx, c.readOnlyClient, read.FindIngredientByID, ingredientId).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("ingredientID", ingredientId.String()).
			Msg("Failed to get ingredient by id")
		return nil, err
	}

	return result.ToEntity(), nil
}

func (c IngredientRepository) GetByProductID(ctx context.Context, productID uuid.UUID) ([]read.FindIngredientQueryResult, error) {

	result, err := sql.NewQuery[read.FindIngredientQueryResult](ctx, c.readOnlyClient, read.FindIngredientsByProductID, productID).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("productID", productID.String()).
			Msg("Failed to get ingredient by productID")
		return nil, err
	}

	return result, nil
}

func NewIngredientRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *IngredientRepository {
	return &IngredientRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
