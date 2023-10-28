package input

import (
	"context"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductFinderServicePort interface {
	FindByID(ctx context.Context, id int) (*entity.ProductEntity, error)

	FindByCategory(ctx context.Context, category string) ([]entity.ProductEntity, error)

	FindAllProducts(ctx context.Context) ([]entity.ProductEntity, error)
}
