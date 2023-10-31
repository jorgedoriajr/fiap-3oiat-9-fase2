package service

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/service"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	productFinderServiceInstance *ProductFinderService
	productFinderServiceOnce     sync.Once
)

type ProductFinderService struct {
	productPersistencePort  output.ProductPersistencePort
	ingredientFinderService service.IngredientFinderService
}

func (p ProductFinderService) FindAllProducts(ctx context.Context) ([]entity.ProductEntity, error) {
	return p.productPersistencePort.GetAll(ctx)
}

func (p ProductFinderService) FindByCategory(ctx context.Context, category string) ([]entity.ProductEntity, error) {
	return p.productPersistencePort.GetByCategory(ctx, category)
}

func (p ProductFinderService) FindByID(ctx context.Context, id uuid.UUID) (*entity.ProductEntity, error) {
	return p.productPersistencePort.GetByID(ctx, id)
}

func (p ProductFinderService) FindByIDWithIngredients(ctx context.Context, id uuid.UUID) (*result.FindProductWithIngredientsResult, error) {
	product, pErr := p.FindByID(ctx, id)
	if pErr != nil {
		return nil, pErr
	}
	ingredients, iErr := p.ingredientFinderService.FindIngredientsByProductId(ctx, id)
	if iErr != nil {
		return nil, iErr
	}

	r := &result.FindProductWithIngredientsResult{
		ID:          product.ID,
		Name:        product.Name,
		Number:      product.Number,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		Ingredients: ingredients,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	r.CalculateIngredientsAmount()
	return r, nil
}

func (p ProductFinderService) FindByNumber(ctx context.Context, number int) (*entity.ProductEntity, error) {
	return p.productPersistencePort.GetByNumber(ctx, number)
}

func NewProductFinderService(
	productPersistence output.ProductPersistencePort,
	ingredientFinder service.IngredientFinderService,
) *ProductFinderService {
	productFinderServiceOnce.Do(func() {
		productFinderServiceInstance = &ProductFinderService{
			productPersistencePort:  productPersistence,
			ingredientFinderService: ingredientFinder,
		}
	})
	return productFinderServiceInstance
}
