package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	ingredientOutput "hamburgueria/internal/modules/ingredient/ports/output"
	"hamburgueria/internal/modules/product/domain"
	"hamburgueria/internal/modules/product/ports/input"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	createProductUseCaseInstance input.CreateProductUseCasePort
	createProductUseCaseOnce     sync.Once
)

type CreateProductUseCase struct {
	productPersistencePort    output.ProductPersistencePort
	productCategoryPort       output.ProductCategoryPersistencePort
	ingredientPersistencePort ingredientOutput.IngredientPersistencePort
}

func (c CreateProductUseCase) AddProduct(
	ctx context.Context,
	command command.CreateProductCommand,
) (*result.ProductResult, error) {
	productID := uuid.New()

	category, err := c.productCategoryPort.GetConfig(ctx, command.Category)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	mandatoryTypes := make(map[string]bool)
	for _, config := range category.ConfigByProductCategory {
		if !config.Optional {
			mandatoryTypes[config.IngredientType] = false
		}
	}

	amount, productIngredients, err := c.buildIngredients(ctx, command, productID, mandatoryTypes)
	if err != nil {
		return nil, err
	}

	product := command.ToProductDomain(productIngredients, amount, productID, *category)
	existentProduct, err := c.productPersistencePort.CheckProductExists(ctx, product)

	if existentProduct != nil {
		return result.FromDomain(*existentProduct), nil
	}
	if err != nil {
		return nil, err
	}

	err = c.productPersistencePort.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return result.FromDomain(product), nil
}

func (c CreateProductUseCase) buildIngredients(
	ctx context.Context,
	command command.CreateProductCommand,
	productID uuid.UUID,
	mandatoryTypes map[string]bool,
) (int, []domain.ProductIngredient, error) {
	var amount int
	var productIngredients []domain.ProductIngredient
	for _, ingredient := range command.Ingredients {
		ingredientDomain, err := c.ingredientPersistencePort.GetByNumber(ctx, ingredient.Number)
		if err != nil {
			return 0, nil, err
		}
		if ingredientDomain == nil {
			return 0, nil, errors.New(fmt.Sprintf("ingredient %d not found", ingredient.Number))
		}

		mandatoryTypes[ingredientDomain.Type.Name] = true

		var containsCategory = false
		for _, config := range ingredientDomain.Type.ConfigByProductCategory {
			if config.ProductCategory == command.Category {
				containsCategory = true
				if config.MaxQtd != 0 && ingredient.Quantity > config.MaxQtd {
					return 0, nil, errors.New(
						fmt.Sprintf("ingredient %d quantity exceeds the limit", ingredient.Number),
					)
				}
			}
		}

		if !containsCategory {
			return 0, nil, errors.New(
				fmt.Sprintf("ingredient %d is from a different category", ingredient.Number),
			)
		}

		productIngredients = append(productIngredients, domain.ProductIngredient{
			ID:         uuid.New(),
			ProductId:  productID,
			Ingredient: *ingredientDomain,
			Quantity:   ingredient.Quantity,
			Amount:     ingredientDomain.Amount * ingredient.Quantity,
		})

		amount = amount + ingredientDomain.Amount*ingredient.Quantity
	}

	for mandatoryType, contains := range mandatoryTypes {
		if contains == false {
			return 0, nil, errors.New(
				fmt.Sprintf("mandatory ingredient type %s missing", mandatoryType),
			)
		}
	}

	return amount, productIngredients, nil
}

func GetCreateProductUseCase(
	productPersistence output.ProductPersistencePort,
	ingredientPersistencePort ingredientOutput.IngredientPersistencePort,
	productCategoryPort output.ProductCategoryPersistencePort,
) input.CreateProductUseCasePort {
	createProductUseCaseOnce.Do(func() {
		createProductUseCaseInstance = CreateProductUseCase{
			productPersistencePort:    productPersistence,
			ingredientPersistencePort: ingredientPersistencePort,
			productCategoryPort:       productCategoryPort,
		}
	})
	return createProductUseCaseInstance
}
