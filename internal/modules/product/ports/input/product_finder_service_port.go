package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/usecase/result"
)

type ProductFinderServicePort interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.ProductEntity, error)

	FindByIDWithIngredients(ctx context.Context, id uuid.UUID) (*result.FindProductWithIngredientsResult, error)

	FindByNumber(ctx context.Context, number int) (*result.FindProductWithIngredientsResult, error)

	FindByCategory(ctx context.Context, category string) ([]*result.FindProductWithIngredientsResult, error)

	FindAllProducts(ctx context.Context) ([]*result.FindProductWithIngredientsResult, error)
}
