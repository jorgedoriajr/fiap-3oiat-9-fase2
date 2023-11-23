package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/internal/modules/product/infra/database/model"
	"hamburgueria/internal/modules/product/ports/output"
	"strconv"
	"sync"
)

type ProductRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []model.Product
	tx := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Find(&products)
	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Msg("Failed to find all products")
		return nil, tx.Error
	}

	var domainProducts []domain.Product
	for _, product := range products {
		domainProducts = append(domainProducts, *product.ToDomain())
	}

	return domainProducts, nil
}

func (c ProductRepository) GetByCategory(ctx context.Context, category string) ([]domain.Product, error) {
	var products []model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Joins("ProductCategory", c.readOnlyClient.Where(&model.ProductCategory{Name: category})).
		Find(&products).Error

	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("category", category).
			Msg("Failed to find products by category")
		return nil, err
	}

	var domainProducts []domain.Product
	for _, product := range products {
		domainProducts = append(domainProducts, *product.ToDomain())
	}

	return domainProducts, nil
}

func (c ProductRepository) Create(ctx context.Context, product domain.Product) error {

	var ingredients []model.ProductIngredient
	for _, ingredient := range product.Ingredients {
		ingredients = append(ingredients, model.ProductIngredientFromDomain(ingredient))
	}

	err := c.readWriteClient.
		Create(&model.Product{
			ID:          product.ID,
			Number:      product.Number,
			Name:        product.Name,
			Amount:      product.Amount,
			Description: product.Description,
			ProductCategory: model.ProductCategory{
				Name:         product.Category.Name,
				AcceptCustom: product.Category.AcceptCustom,
			},
			Menu:        product.Menu,
			ImgPath:     product.ImgPath,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
			Ingredients: ingredients,
		}).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("product", product.Name).
			Msg("Failed to insert product")
		return err
	}
	return nil
}

func (c ProductRepository) Update(ctx context.Context, product domain.Product) error {
	productModel := model.ProductFromDomain(product)
	err := c.readWriteClient.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Omit("Category").
		Save(&productModel).
		Error

	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productId", product.ID.String()).
			Str("product", product.Name).
			Msg("Failed to update product")
		return err
	}

	return nil
}

func (c ProductRepository) GetByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	var product model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		First(&product, productID).
		Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productId", productID.String()).
			Msg("Failed to find product by ID")
		return nil, err
	}
	return product.ToDomain(), nil
}

func (c ProductRepository) GetByNumber(ctx context.Context, productNumber int) (*domain.Product, error) {
	var product model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Where("number = ?", productNumber).
		First(&product).
		Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productNumber", strconv.Itoa(productNumber)).
			Msg("Failed to find product by number")
		return nil, err
	}
	return product.ToDomain(), nil
}

func (c ProductRepository) ProductAlreadyExists(ctx context.Context, product domain.Product) (bool, error) {
	var count int64
	query := "product.name = ?"
	args := []interface{}{product.Name}

	for _, ingredient := range product.Ingredients {

		query += `
			AND EXISTS (
				SELECT 1 FROM product_ingredient JOIN ingredient ON product_ingredient.ingredient_id = ingredient.id 
				WHERE product.id = product_ingredient.product_id AND ingredient.name = ? AND product_ingredient.quantity = ?)
			)
		`
		args = append(args, ingredient.Ingredient.Name, ingredient.Quantity)
	}

	err := c.readOnlyClient.Table("product").
		Where(query, args...).
		Count(&count).Error

	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productName", product.Name).
			Msg("Failed to find if product already exists")
		return false, err
	}

	return count > 0, nil
}

var (
	productRepositoryInstance output.ProductPersistencePort
	productRepositoryOnce     sync.Once
)

func GetProductRepository(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.ProductPersistencePort {
	productRepositoryOnce.Do(func() {
		productRepositoryInstance = ProductRepository{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return productRepositoryInstance
}
