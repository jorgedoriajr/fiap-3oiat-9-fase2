package database

import (
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/internal/modules/product/infra/database/model"
	"hamburgueria/internal/modules/product/ports/output"
	"sync"
)

type ProductCategoryRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c ProductCategoryRepository) GetAll(ctx context.Context) ([]domain.ProductCategory, error) {
	var categories []model.ProductCategory
	err := c.readOnlyClient.Find(&categories).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Msg("Failed to find all categories")
		return nil, err
	}

	var domainCategories []domain.ProductCategory
	for _, category := range categories {
		domainCategories = append(domainCategories, *category.ToDomain())
	}

	return domainCategories, nil
}

func (c ProductCategoryRepository) GetByName(ctx context.Context, name string) (*domain.ProductCategory, error) {
	var category model.ProductCategory
	err := c.readOnlyClient.Find(&category, name).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("name", name).
			Msg("Failed to find category by name")
		return nil, err
	}

	return category.ToDomain(), nil
}

var (
	productCategoryOnce     sync.Once
	productCategoryInstance output.ProductCategoryPersistencePort
)

func GetProductCategoryRepository(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.ProductCategoryPersistencePort {
	productCategoryOnce.Do(func() {
		productCategoryInstance = ProductCategoryRepository{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return productCategoryInstance
}
