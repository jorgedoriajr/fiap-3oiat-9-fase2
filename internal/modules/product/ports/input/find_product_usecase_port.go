package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
)

type FindProductUseCasePort interface {
	FindByID(ctx context.Context, id uuid.UUID) (*result.FindProductResult, error)

	FindByNumber(ctx context.Context, number int) (*result.FindProductResult, error)

	FindByCategory(ctx context.Context, category string) ([]result.FindProductResult, error)

	FindAllProducts(ctx context.Context) ([]result.FindProductResult, error)
}
