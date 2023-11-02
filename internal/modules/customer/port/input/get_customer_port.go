package input

import (
	"context"
	"hamburgueria/internal/server/api/rest/v1/customer/response"
)

type GetCustomerPort interface {
	GetCustomer(ctx context.Context, document string) (*response.Customer, error)
}
