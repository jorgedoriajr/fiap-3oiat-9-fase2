package input

import (
	"context"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
)

type UpdateProductUseCasePort interface {
	UpdateProduct(ctx context.Context, command command.UpdateProductCommand) (result.UpdateProductResult, error)
}
