package input

import (
	"context"
	"hamburgueria/internal/modules/product/usecase/command"
)

type CreateProductCategoryUseCasePort interface {
	AddProductCategory(ctx context.Context, command command.CreateProductCategoryCommand) (any, error)
}
