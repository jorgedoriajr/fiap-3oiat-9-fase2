package usecase

import (
	"context"
	"fmt"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	createProductCategoryUseCaseInstance *CreateProductCategoryUseCase
	createProductCategoryUseCaseOnce     sync.Once
)

type CreateProductCategoryUseCase struct {
	productCategoryPersistencePort output.ProductCategoryPersistencePort
}

func (c CreateProductCategoryUseCase) AddProductCategory(ctx context.Context, command command.CreateProductCategoryCommand) (result.CreateProductCategoryResult, error) {
	category := command.ToEntity()
	fmt.Printf("creating new product category: [%v]", category)
	categoryCreated, err := c.productCategoryPersistencePort.CreateProductCategory(ctx, category)
	if err != nil {
		return result.CreateProductCategoryResult{}, err
	}
	return result.ToCreateProductCategoryResultFrom(*categoryCreated), nil
}

func NewCreateCategoryProductUseCase(productCategoryPersistence output.ProductCategoryPersistencePort) *CreateProductCategoryUseCase {
	createProductCategoryUseCaseOnce.Do(func() {
		createProductCategoryUseCaseInstance = &CreateProductCategoryUseCase{
			productCategoryPersistencePort: productCategoryPersistence,
		}
	})
	return createProductCategoryUseCaseInstance
}
