package postgres

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/read"
	"hamburgueria/pkg/sql"
)

type ProductCategoryRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c ProductCategoryRepository) GetAllProductCategories(ctx context.Context) ([]entity.ProductCategoryEntity, error) {
	allProduct, allProductErr := sql.NewQuery[read.FindProductCategoryQueryResult](
		ctx,
		c.readOnlyClient,
		read.FindAllProductCategories,
	).Many()

	if allProductErr != nil {
		c.logger.Error().
			Err(allProductErr).
			Msg("Failed to get product categories")
		return []entity.ProductCategoryEntity{}, allProductErr
	}

	return read.ToProductCategoryEntityList(allProduct), nil
}

func (c ProductCategoryRepository) GetProductCategoryByName(ctx context.Context, category string) (*entity.ProductCategoryEntity, error) {
	productCategoryByName, productCategoryByNameErr := sql.NewQuery[*entity.ProductCategoryEntity](
		ctx,
		c.readOnlyClient,
		read.FindProductCategoryByName,
		category,
	).One()

	if productCategoryByNameErr != nil {
		c.logger.Error().
			Err(productCategoryByNameErr).
			Str("category", category).
			Msg("Failed to get product categories by name")
		return nil, productCategoryByNameErr
	}
	return productCategoryByName, nil
}

func NewProductCategoryRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *ProductCategoryRepository {
	return &ProductCategoryRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
