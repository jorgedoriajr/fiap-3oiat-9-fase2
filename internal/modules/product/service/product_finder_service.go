package service

import (
	"context"
	"github.com/google/uuid"
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

func (p ProductFinderService) FindByID(ctx context.Context, id uuid.UUID) (*entity.ProductEntity, error) {
	return p.productPersistencePort.GetByID(ctx, id)
}

func (p ProductFinderService) FindByNumber(ctx context.Context, number int) (*entity.ProductEntity, error) {
	return p.productPersistencePort.GetByNumber(ctx, number)
}

func NewProductFinderService(productPersistence output.ProductPersistencePort) *ProductFinderService {
	productFinderServiceOnce.Do(func() {
		productFinderServiceInstance = &ProductFinderService{
			productPersistencePort: productPersistence,
		}
	})
	return productFinderServiceInstance
}
