package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	ingredientService "hamburgueria/internal/modules/ingredient/service"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	createProductUseCaseInstance *CreateProductUseCase
	createProductUseCaseOnce     sync.Once
)

type CreateProductUseCase struct {
	productPersistencePort           output.ProductPersistencePort
	ingredientFinderService          ingredientService.IngredientFinderService
	productIngredientPersistencePort output.ProductIngredientPersistencePort
}

func (c CreateProductUseCase) AddProduct(ctx context.Context, command command.CreateProductCommand) (result.CreateProductResult, error) {
	productId := uuid.New()
	amount, productIngredients, err := c.buildIngredients(ctx, command, productId)
	if err != nil {
		return result.CreateProductResult{}, err
	}

	product := command.ToProductEntity(productId, productIngredients, amount)
	err = c.productPersistencePort.Create(ctx, product)
	if err != nil {
		return result.CreateProductResult{}, err
	}

	err = c.createProductIngredients(ctx, productIngredients)
	if err != nil {
		return result.CreateProductResult{}, err
	}

	return result.FromEntity(product), nil
}

func (c CreateProductUseCase) createProductIngredients(ctx context.Context, productIngredients []entity.ProductIngredientEntity) error {
	for _, productIngredient := range productIngredients {
		err := c.productIngredientPersistencePort.Create(ctx, productIngredient)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c CreateProductUseCase) buildIngredients(ctx context.Context, command command.CreateProductCommand, productId uuid.UUID) (int, []entity.ProductIngredientEntity, error) {
	var amount int
	var productIngredients []entity.ProductIngredientEntity
	for _, ingredient := range command.Ingredients {
		ingredientEntity, err := c.ingredientFinderService.FindIngredientByID(ctx, uuid.MustParse(ingredient.ID))
		if err != nil {
			return 0, nil, err
		}
		if ingredientEntity == nil {
			return 0, nil, errors.New("ingredient not found")
		}

		productIngredients = append(productIngredients, entity.ProductIngredientEntity{
			ID:           uuid.New(),
			ProductId:    productId,
			IngredientId: uuid.MustParse(ingredient.ID),
			Quantity:     ingredient.Quantity,
			Amount:       ingredientEntity.Amount * ingredient.Quantity,
		})

		amount = amount + ingredientEntity.Amount*ingredient.Quantity
	}
	return amount, productIngredients, nil
}

func NewCreateProductUseCase(
	productPersistence output.ProductPersistencePort,
	ingredientFinderService ingredientService.IngredientFinderService,
	productIngredientPersistencePort output.ProductIngredientPersistencePort,
) *CreateProductUseCase {
	createProductUseCaseOnce.Do(func() {
		createProductUseCaseInstance = &CreateProductUseCase{
			productPersistencePort:           productPersistence,
			ingredientFinderService:          ingredientFinderService,
			productIngredientPersistencePort: productIngredientPersistencePort,
		}
	})
	return createProductUseCaseInstance
}
