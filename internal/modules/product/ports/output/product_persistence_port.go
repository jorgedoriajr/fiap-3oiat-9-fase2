package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain"
)

type ProductPersistencePort interface {
	Create(ctx context.Context, product domain.Product) error
	Update(ctx context.Context, product domain.Product) error
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error)
	GetByNumber(ctx context.Context, productNumber int) (*domain.Product, error)
	GetByCategory(ctx context.Context, productID string) ([]domain.Product, error)
	ProductAlreadyExists(ctx context.Context, product domain.Product) (bool, error)
}
