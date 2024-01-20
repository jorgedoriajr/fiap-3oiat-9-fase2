package usecase

import (
	"context"
	"errors"
	"fmt"
	"hamburgueria/internal/modules/product/ports/input"
	"hamburgueria/internal/modules/product/ports/output"
	"sync"
)

var (
	deleteProductUseCaseInstance DeleteProductUseCase
	deleteProductUseCaseOnce     sync.Once
)

type DeleteProductUseCase struct {
	productPersistenceGateway output.ProductPersistencePort
}

func (d DeleteProductUseCase) Inactive(ctx context.Context, number int) error {
	product, err := d.productPersistenceGateway.GetByNumber(ctx, number)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New(fmt.Sprintf("product %d not found", number))
	}
	product.Active = false
	return d.productPersistenceGateway.Update(ctx, *product)
}

func GetDeleteProductUseCase(productPersistenceGateway output.ProductPersistencePort) input.DeleteProductUseCasePort {
	deleteProductUseCaseOnce.Do(func() {
		deleteProductUseCaseInstance = DeleteProductUseCase{
			productPersistenceGateway: productPersistenceGateway,
		}
	})
	return deleteProductUseCaseInstance
}
