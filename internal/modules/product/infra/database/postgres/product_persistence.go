package postgres

import (
	"context"
	"github.com/google/uuid"
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

func (c ProductRepository) GetAll(ctx context.Context) ([]entity.ProductEntity, error) {
	allProduct, allProductErr := sql.NewQuery[entity.ProductEntity](
		ctx,
		c.readOnlyClient,
		read.FindAllProducts,
	).Many()

	if allProductErr != nil {
		c.logger.Error().
			Err(allProductErr).
			Msg("Failed to get products")
		return []entity.ProductEntity{}, allProductErr
	}

	return allProduct, nil
}

func (c ProductRepository) GetByCategory(ctx context.Context, category string) ([]entity.ProductEntity, error) {
	productByCategory, productByCategoryErr := sql.NewQuery[entity.ProductEntity](
		ctx,
		c.readOnlyClient,
		read.FindProductByCategory,
		category,
	).Many()

	if productByCategoryErr != nil {
		c.logger.Error().
			Err(productByCategoryErr).
			Str("category", category).
			Msg("Failed to get products by category")
		return nil, productByCategoryErr
	}
	return productByCategory, nil
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

func (c ProductRepository) GetByID(ctx context.Context, productID uuid.UUID) (*entity.ProductEntity, error) {

	result, err := sql.NewQuery[read.FindProductQueryResult](ctx, c.readOnlyClient, read.FindProductByID, productID).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("productID", productID.String()).
			Msg("Failed to get product by id")
		return nil, err
	}

	return result.ToEntity(), nil
}

func (c ProductRepository) GetByNumber(ctx context.Context, productNumber int) (*entity.ProductEntity, error) {
	result, err := sql.NewQuery[read.FindProductQueryResult](ctx, c.readOnlyClient, read.FindProductByID, productNumber).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Int("productNumber", productNumber).
			Msg("Failed to get product by number")
		return nil, err
	}

	return result.ToEntity(), nil
}

func NewProductRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *ProductRepository {
	return &ProductRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
