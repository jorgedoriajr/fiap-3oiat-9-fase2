package service

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/ports/output"
	"sync"
)

var (
	productFinderServiceInstance *ProductFinderService
	productFinderServiceOnce     sync.Once
)

type ProductFinderService struct {
	productPersistencePort output.ProductPersistencePort
}

func (p ProductFinderService) FindAllProducts(ctx context.Context) ([]entity.ProductEntity, error) {
	return p.productPersistencePort.GetAll(ctx)
}

func (p ProductFinderService) FindByCategory(ctx context.Context, category string) ([]entity.ProductEntity, error) {
	return p.productPersistencePort.GetByCategory(ctx, category)
}

func (p ProductFinderService) FindByID(ctx context.Context, id int) (*entity.ProductEntity, error) {
	return p.productPersistencePort.GetByID(ctx, id)
}

func NewProductFinderService(productPersistence output.ProductPersistencePort) *ProductFinderService {
	productFinderServiceOnce.Do(func() {
		productFinderServiceInstance = &ProductFinderService{
			productPersistencePort: productPersistence,
		}
	})
	return productFinderServiceInstance
}
