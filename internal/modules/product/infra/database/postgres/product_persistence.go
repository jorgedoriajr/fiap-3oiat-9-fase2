package postgres

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type ProductRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c ProductRepository) Create(ctx context.Context, product entity.ProductEntity) error {

	mapper := write.ToInsertProductQueryMapper(product)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertProductRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("name", product.Name).
			Msg("Failed to insert product")
		return err
	}

	return nil
}

func (c ProductRepository) GetByID(ctx context.Context, productID string) (*entity.ProductEntity, error) {

	result, err := sql.NewQuery[read.FindProductQueryResult](ctx, c.readOnlyClient, read.FindProductByID, productID).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("productID", productID).
			Msg("Failed to get product")
		return nil, err
	}

	return result.ToEntity(), nil
}
