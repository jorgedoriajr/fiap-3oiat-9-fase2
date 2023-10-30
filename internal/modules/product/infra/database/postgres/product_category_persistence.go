package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/write"
	"hamburgueria/pkg/sql"
)

type ProductCategoryRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c ProductCategoryRepository) CreateProductCategory(ctx context.Context, category entity.ProductCategoryEntity) (*entity.ProductCategoryEntity, error) {
	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertProductCategoryRW, category)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("Name", category.Name).
			Msg("Failed to insert new product category")
		return nil, err
	}

	return &entity.ProductCategoryEntity{ID: category.ID, Name: category.Name}, nil
}

func (c ProductCategoryRepository) GetAllProductCategories(ctx context.Context) ([]entity.ProductCategoryEntity, error) {
	allProduct, allProductErr := sql.NewQuery[entity.ProductCategoryEntity](
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

	return allProduct, nil
}

func (c ProductCategoryRepository) GetProductCategoryByName(ctx context.Context, category string) (*entity.ProductCategoryEntity, error) {
	productCategoryByName, productCategoryByNameErr := sql.NewQuery[*entity.ProductCategoryEntity](
		ctx,
		c.readOnlyClient,
		read.FindProductByCategory,
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

func (c ProductCategoryRepository) GetProductCategoryById(ctx context.Context, id uuid.UUID) (*entity.ProductCategoryEntity, error) {

	result, err := sql.NewQuery[read.FindProductQueryResult](ctx, c.readOnlyClient, read.FindProductCategoryByID, id).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("categoryID", id.String()).
			Msg("Failed to get product category by id")
		return nil, err
	}

	return result.ToProductCategoryEntity(), nil
}

func NewProductCategoryRepository(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) *ProductCategoryRepository {
	return &ProductCategoryRepository{readWriteClient: readWriteClient, readOnlyClient: readOnlyClient, logger: logger}
}
