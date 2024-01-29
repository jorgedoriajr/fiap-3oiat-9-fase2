package database

import (
	"context"
	"errors"
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

type ProductPersistenceGateway struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c ProductPersistenceGateway) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []model.Product
	tx := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Table("product").
		Where("product.active = ?", true).
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

func (c ProductPersistenceGateway) GetByCategory(ctx context.Context, category string) ([]domain.Product, error) {
	var products []model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Table("product").
		Joins("JOIN product_category ON product_category.name = product.category").
		Where("product.category = ?", category).
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

func (c ProductPersistenceGateway) Create(ctx context.Context, product domain.Product) error {

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
			Active:      true,
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

func (c ProductPersistenceGateway) Update(ctx context.Context, product domain.Product) error {
	productModel := model.ProductFromDomain(product)
	err := c.readWriteClient.
		Session(&gorm.Session{FullSaveAssociations: false}).
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

func (c ProductPersistenceGateway) GetByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	var product model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		First(&product, productID).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productId", productID.String()).
			Msg("Failed to find product by ID")
		return nil, err
	}
	return product.ToDomain(), nil
}

func (c ProductPersistenceGateway) GetByNumber(ctx context.Context, productNumber int) (*domain.Product, error) {
	var product model.Product
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Preload("Ingredients.Ingredient.IngredientType").
		Where("number = ?", productNumber).
		First(&product).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productNumber", strconv.Itoa(productNumber)).
			Msg("Failed to find product by number")
		return nil, err
	}
	return product.ToDomain(), nil
}

func (c ProductPersistenceGateway) CheckProductExists(ctx context.Context, product domain.Product) (*domain.Product, error) {
	query := "product.name = ?"
	args := []interface{}{product.Name}

	for _, ingredient := range product.Ingredients {

		query += `
			AND EXISTS (
				SELECT 1 FROM product_ingredient JOIN ingredient ON product_ingredient.ingredient_id = ingredient.id 
				WHERE product.id = product_ingredient.product_id AND ingredient.name = ? AND product_ingredient.quantity = ?
			)
		`
		args = append(args, ingredient.Ingredient.Name, ingredient.Quantity)
	}

	var existentProduct model.Product
	err := c.readOnlyClient.Table("product").
		Where(query, args...).
		First(&existentProduct).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("productName", product.Name).
			Msg("Failed to find if product already exists")
		return nil, err
	}

	return existentProduct.ToDomain(), nil
}

var (
	productPersistenceGatewayInstance output.ProductPersistencePort
	productPersistenceGatewayOnce     sync.Once
)

func GetProductPersistenceGateway(
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
	logger zerolog.Logger,
) output.ProductPersistencePort {
	productPersistenceGatewayOnce.Do(func() {
		productPersistenceGatewayInstance = ProductPersistenceGateway{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return productPersistenceGatewayInstance
}
