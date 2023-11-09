package input

import (
	"context"
	"hamburgueria/internal/modules/product/usecase/command"
)

type UpdateProductUseCasePort interface {
	UpdateProduct(ctx context.Context, command command.UpdateProductCommand) error
}
