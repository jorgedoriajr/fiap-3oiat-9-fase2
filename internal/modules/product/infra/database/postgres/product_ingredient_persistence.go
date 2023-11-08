package postgres

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type ProductIngredientRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c ProductIngredientRepository) Create(ctx context.Context, productIngredient entity.ProductIngredientEntity) error {

	mapper := write.ToInsertProductIngredientQueryMapper(productIngredient)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertProductIngredientRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("ProductId", productIngredient.ProductId.String()).
			Msg("Failed to insert product_ingredient")
		return err
	}

	return nil
}

func NewProductIngredientRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *ProductIngredientRepository {
	return &ProductIngredientRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
