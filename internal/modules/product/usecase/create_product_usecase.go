package usecase

import (
	"context"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
)

var (
	name = "CreateProductUseCase"
)

type CreateProductUseCase struct {
	productPersistencePort output.ProductPersistencePort
}

func (c CreateProductUseCase) AddProduct(ctx context.Context, command command.CreateProductCommand) (result.CreateProductResult, error) {
	err := c.productPersistencePort.Persist(ctx, command.ToProductEntity())
	if err != nil {
		return result.CreateProductResult{}, err
	}
	panic("implement me")
}
