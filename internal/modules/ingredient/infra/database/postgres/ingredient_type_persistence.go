package postgres

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"

	"github.com/rs/zerolog"
)

type IngredientTypeRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c IngredientTypeRepository) Create(ctx context.Context, ingredientType entity.IngredientType) error {

	mapper := write.ToInsertIngredientTypeQueryMapper(ingredientType)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertIngredientTypeRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("name", string(ingredientType.Name)).
			Msg("Failed to insert ingredient")
		return err
	}

	return nil
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
