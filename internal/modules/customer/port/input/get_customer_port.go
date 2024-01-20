package input

import (
	"context"
	"hamburgueria/internal/modules/customer/usecase/result"
)

type GetCustomerPort interface {
	GetCustomer(ctx context.Context, document string) (*result.CustomerCreated, error)
}
