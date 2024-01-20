package usecase

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	findProductUseCaseInstance FindProductUseCase
	findProductUseCaseOnce     sync.Once
)

type FindProductUseCase struct {
	productPersistenceGateway output.ProductPersistencePort
}

func (f FindProductUseCase) FindByID(ctx context.Context, id uuid.UUID) (*result.FindProductResult, error) {
	product, err := f.productPersistenceGateway.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	productResult := result.FromProductDomain(*product)
	return &productResult, nil
}
func (f FindProductUseCase) FindByNumber(ctx context.Context, number int) (*result.FindProductResult, error) {
	product, err := f.productPersistenceGateway.GetByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	productResult := result.FromProductDomain(*product)
	return &productResult, nil
}
func (f FindProductUseCase) FindByCategory(ctx context.Context, category string) ([]result.FindProductResult, error) {
	products, err := f.productPersistenceGateway.GetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	var productsResult []result.FindProductResult
	for _, product := range products {
		productsResult = append(productsResult, result.FromProductDomain(product))
	}

	return productsResult, nil
}
func (f FindProductUseCase) FindAllProducts(ctx context.Context) ([]result.FindProductResult, error) {
	products, err := f.productPersistenceGateway.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var productsResult []result.FindProductResult
	for _, product := range products {
		productsResult = append(productsResult, result.FromProductDomain(product))
	}

	return productsResult, nil
}

func NewFindProductUseCase(
	productPersistenceGateway output.ProductPersistencePort,
) FindProductUseCase {
	findProductUseCaseOnce.Do(func() {
		findProductUseCaseInstance = FindProductUseCase{
			productPersistenceGateway: productPersistenceGateway,
		}
	})
	return findProductUseCaseInstance
}
