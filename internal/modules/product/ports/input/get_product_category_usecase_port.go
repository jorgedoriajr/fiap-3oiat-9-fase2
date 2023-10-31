package input

import (
	"context"
	"hamburgueria/internal/modules/product/usecase/result"
)

type GetProductCategoryUseCasePort interface {
	FindAll(ctx context.Context) ([]result.FindProductCategoryResult, error)
}
