package usecase

import (
	"context"
	"hamburgueria/internal/modules/product/ports/output"
	"hamburgueria/internal/modules/product/usecase/result"
	"sync"
)

var (
	getProductCategoryUseCaseInstance *GetProductCategoryUseCase
	getProductCategoryUseCaseOnce     sync.Once
)

type GetProductCategoryUseCase struct {
	productCategoryPersistenceGateway output.ProductCategoryPersistencePort
}

func (c GetProductCategoryUseCase) FindAll(ctx context.Context) ([]result.FindProductCategoryResult, error) {
	productCategories, err := c.productCategoryPersistenceGateway.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var categoriesResult []result.FindProductCategoryResult
	for _, productCategory := range productCategories {
		categoriesResult = append(categoriesResult, result.FindProductCategoryResult{
			Name:         productCategory.Name,
			AcceptCustom: productCategory.AcceptCustom,
		})
	}
	return categoriesResult, nil
}

func NewGetProductCategoryUseCase(productCategoryPersistenceGateway output.ProductCategoryPersistencePort) *GetProductCategoryUseCase {
	getProductCategoryUseCaseOnce.Do(func() {
		getProductCategoryUseCaseInstance = &GetProductCategoryUseCase{
			productCategoryPersistenceGateway: productCategoryPersistenceGateway,
		}
	})
	return getProductCategoryUseCaseInstance
}
