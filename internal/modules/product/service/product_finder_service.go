package service

import (
	"context"
	"github.com/google/uuid"
	readIngredient "hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/ingredient/service"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/infra/database/postgres/sql/read"
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

func (p ProductFinderService) FindAllProducts(ctx context.Context) ([]*result.FindProductWithIngredientsResult, error) {

	products, err := p.productPersistencePort.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return p.getIngredientsForProducts(ctx, products)
}

func (p ProductFinderService) FindByCategory(ctx context.Context, category string) ([]*result.FindProductWithIngredientsResult, error) {
	products, err := p.productPersistencePort.GetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return p.getIngredientsForProducts(ctx, products)
}

func (p ProductFinderService) FindByOrderID(ctx context.Context, orderId uuid.UUID) ([]*result.FindProductWithIngredientsResult, error) {
	products, err := p.productPersistencePort.GetByOrderID(ctx, orderId)
	if err != nil {
		return nil, err
	}
	var productsWithIngredient []*result.FindProductWithIngredientsResult

	for _, product := range products {
		ingredients, iErr := p.ingredientFinderService.FindIngredientsByProductId(ctx, product.ID)
		if iErr != nil {
			return nil, iErr
		}
		r := convertFindProductOrderQueryResultWithIngredientsResult(ingredients, product)

		productsWithIngredient = append(productsWithIngredient, r)
	}
	return productsWithIngredient, nil

}

func convertFindProductOrderQueryResultWithIngredientsResult(ingredients []readIngredient.FindProductIngredientQueryResult, product read.FindProductOrderQueryResult) *result.FindProductWithIngredientsResult {
	var ingredientsResult []result.FindProductsIngredientsResult
	for _, ingredient := range ingredients {
		ingredientsResult = append(ingredientsResult, ingredient.ToResult())
	}
	r := &result.FindProductWithIngredientsResult{
		ID:          product.ID,
		Name:        product.Name,
		Number:      product.Number,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		ImgPath:     product.ImgPath,
		Amount:      product.Amount,
		Ingredients: ingredientsResult,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return r
}

func convertWithIngredientsResult(ingredients []readIngredient.FindProductIngredientQueryResult, product *entity.ProductEntity) *result.FindProductWithIngredientsResult {
	var ingredientsResult []result.FindProductsIngredientsResult
	for _, ingredient := range ingredients {
		ingredientsResult = append(ingredientsResult, ingredient.ToResult())
	}
	r := &result.FindProductWithIngredientsResult{
		ID:          product.ID,
		Name:        product.Name,
		Number:      product.Number,
		Description: product.Description,
		Category:    product.Category,
		Menu:        product.Menu,
		ImgPath:     product.ImgPath,
		Amount:      product.Amount,
		Ingredients: ingredientsResult,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return r
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

	return convertWithIngredientsResult(ingredients, product), nil
}

func (p ProductFinderService) FindByNumber(ctx context.Context, number int) (*result.FindProductWithIngredientsResult, error) {
	product, err := p.productPersistencePort.GetByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	ingredients, iErr := p.ingredientFinderService.FindIngredientsByProductId(ctx, product.ID)
	if iErr != nil {
		return nil, iErr
	}

	return convertWithIngredientsResult(ingredients, product), nil
}

func (p ProductFinderService) getIngredientsForProducts(ctx context.Context, products []entity.ProductEntity) ([]*result.FindProductWithIngredientsResult, error) {
	var productsWithIngredient []*result.FindProductWithIngredientsResult

	for _, product := range products {

		ingredients, iErr := p.ingredientFinderService.FindIngredientsByProductId(ctx, product.ID)
		if iErr != nil {
			return nil, iErr
		}

		productsWithIngredient = append(productsWithIngredient, convertWithIngredientsResult(ingredients, &product))
	}

	return productsWithIngredient, nil
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
