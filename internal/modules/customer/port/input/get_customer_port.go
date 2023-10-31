package input

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/response"
)

type GetCustomerPort interface {
	GetCustomer(ctx context.Context, document string) (*response.Customer, error)
}
