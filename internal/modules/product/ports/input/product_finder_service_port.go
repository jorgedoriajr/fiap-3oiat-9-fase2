package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
)

type ProductFinderServicePort interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.ProductEntity, error)

	//FindByIDWithIngredients(ctx context.Context, id uuid.UUID) (*result.FindProductWithIngredientsResult, error)

	FindByNumber(ctx context.Context, number int) (*entity.ProductEntity, error)

	FindByCategory(ctx context.Context, category string) ([]entity.ProductEntity, error)

	FindAllProducts(ctx context.Context) ([]entity.ProductEntity, error)
}
