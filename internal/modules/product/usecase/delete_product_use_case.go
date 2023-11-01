package usecase

import (
	"context"
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
	return d.productPersistencePort.InactiveByNumber(ctx, number)
}

func GetDeleteProductUseCase(productPersistencePort output.ProductPersistencePort) input.DeleteProductUseCasePort {
	deleteProductUseCaseOnce.Do(func() {
		deleteProductUseCaseInstance = DeleteProductUseCase{
			productPersistencePort: productPersistencePort,
		}
	})
	return deleteProductUseCaseInstance
}
