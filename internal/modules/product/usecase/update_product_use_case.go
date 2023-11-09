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
	"sync"
	"time"
)

type UpdateProductUseCase struct {
	productPersistencePort    output.ProductPersistencePort
	ingredientPersistencePort ingredientOutput.IngredientPersistencePort
}

func (u UpdateProductUseCase) UpdateProduct(ctx context.Context, command command.UpdateProductCommand) error {

	product, err := u.productPersistencePort.GetByNumber(ctx, command.Number)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New(fmt.Sprintf("product %d not found", command.Number))
	}

	product.UpdatedAt = time.Now()
	if command.Ingredients != nil {
		amount, productIngredients, err := u.buildIngredients(ctx, command, product.ID)
		if err != nil {
			return err
		}
		product.Amount = amount
		product.Ingredients = productIngredients
	}

	if command.Name != nil {
		product.Name = *command.Name
	}

	if command.Category != nil {
		product.Category.Name = *command.Category
	}

	if command.ImgPath != nil {
		product.ImgPath = *command.ImgPath
	}

	if command.Menu != nil {
		product.Menu = *command.Menu
	}

	if command.Description != nil {
		product.Description = *command.Description
	}

	err = u.productPersistencePort.Update(ctx, *product)
	if err != nil {
		return err
	}

	return nil
}

func (u UpdateProductUseCase) buildIngredients(
	ctx context.Context,
	command command.UpdateProductCommand,
	productID uuid.UUID,
) (int, []domain.ProductIngredient, error) {
	var amount int
	var productIngredients []domain.ProductIngredient
	for _, ingredient := range command.Ingredients {
		ingredientDomain, err := u.ingredientPersistencePort.GetByNumber(ctx, ingredient.Number)
		if err != nil {
			return 0, nil, err
		}
		if ingredientDomain == nil {
			return 0, nil, errors.New(fmt.Sprintf("ingredient %d not found", ingredient.Number))
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
	return amount, productIngredients, nil
}

var (
	updateProductUseCaseInstance input.UpdateProductUseCasePort
	updateProductUseCaseOnce     sync.Once
)

func GetUpdateProductUseCase(
	productPersistencePort output.ProductPersistencePort,
	ingredientPersistencePort ingredientOutput.IngredientPersistencePort,
) input.UpdateProductUseCasePort {
	updateProductUseCaseOnce.Do(func() {
		updateProductUseCaseInstance = UpdateProductUseCase{
			productPersistencePort:    productPersistencePort,
			ingredientPersistencePort: ingredientPersistencePort,
		}
	})
	return updateProductUseCaseInstance
}
