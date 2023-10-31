package input

import (
	"context"
	"hamburgueria/internal/modules/order/usecase/result"
)

type ListOrderPort interface {
	FindAllOrders(ctx context.Context) ([]result.ListOrderResult, error)
	FindByStatus(ctx context.Context, status string) ([]result.ListOrderResult, error)
}
