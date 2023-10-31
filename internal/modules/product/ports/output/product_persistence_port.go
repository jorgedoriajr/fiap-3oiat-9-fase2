package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductPersistencePort interface {
	Create(ctx context.Context, product entity.ProductEntity) error
	GetAll(ctx context.Context) ([]entity.ProductEntity, error)
	GetByID(ctx context.Context, productID uuid.UUID) (*entity.ProductEntity, error)
	GetByOrderID(ctx context.Context, productID uuid.UUID) ([]entity.ProductEntity, error)
	GetByNumber(ctx context.Context, productNumber int) (*entity.ProductEntity, error)
	GetByCategory(ctx context.Context, productID string) ([]entity.ProductEntity, error)
}
