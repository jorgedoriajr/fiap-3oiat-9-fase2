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
	productCategoryPersistencePort output.ProductCategoryPersistencePort
}

func (c GetProductCategoryUseCase) FindAll(ctx context.Context) ([]result.FindProductCategoryResult, error) {
	productCategories, err := c.productCategoryPersistencePort.GetAllProductCategories(ctx)
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

func NewGetProductCategoryUseCase(productCategoryPersistence output.ProductCategoryPersistencePort) *GetProductCategoryUseCase {
	getProductCategoryUseCaseOnce.Do(func() {
		getProductCategoryUseCaseInstance = &GetProductCategoryUseCase{
			productCategoryPersistencePort: productCategoryPersistence,
		}
	})
	return getProductCategoryUseCaseInstance
}
