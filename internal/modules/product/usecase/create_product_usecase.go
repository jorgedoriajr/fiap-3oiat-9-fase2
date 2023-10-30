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
	createProductUseCaseInstance *CreateProductUseCase
)

type CreateProductUseCase struct {
	productPersistencePort output.ProductPersistencePort
}

func (c CreateProductUseCase) AddProduct(ctx context.Context, command command.CreateProductCommand) (result.CreateProductResult, error) {
	product := command.ToProductEntity()
	fmt.Printf("creating new product: [%v]", product)
	err := c.productPersistencePort.Create(ctx, product)
	if err != nil {
		return result.CreateProductResult{}, err
	}
	fmt.Printf("returning new product: [%v]", product)

	return result.ToCreateProductResultFrom(product), nil
}

func NewCreateProductUseCase(productPersistence output.ProductPersistencePort) *CreateProductUseCase {
	sync.OnceValue[*CreateProductUseCase](func() *CreateProductUseCase {
		createProductUseCaseInstance = &CreateProductUseCase{
			productPersistencePort: productPersistence,
		}
		return createProductUseCaseInstance
	})
	return createProductUseCaseInstance
}
