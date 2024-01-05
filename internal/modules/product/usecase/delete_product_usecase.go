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
	productPersistencePort output.ProductPersistencePort
}

func (d DeleteProductUseCase) Inactive(ctx context.Context, number int) error {
	product, err := d.productPersistencePort.GetByNumber(ctx, number)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New(fmt.Sprintf("product %d not found", number))
	}
	product.Active = false
	return d.productPersistencePort.Update(ctx, *product)
}

func GetDeleteProductUseCase(productPersistencePort output.ProductPersistencePort) input.DeleteProductUseCasePort {
	deleteProductUseCaseOnce.Do(func() {
		deleteProductUseCaseInstance = DeleteProductUseCase{
			productPersistencePort: productPersistencePort,
		}
	})
	return deleteProductUseCaseInstance
}
