package input

import (
	"context"
	"hamburgueria/internal/modules/product/usecase/command"
	"hamburgueria/internal/modules/product/usecase/result"
)

type CreateProductUseCasePort interface {
	AddProduct(ctx context.Context, command command.CreateProductCommand) (*result.ProductResult, error)
}
