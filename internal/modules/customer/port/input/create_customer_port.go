package input

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/request"
)

type CreateCustomerPort interface {
	AddCustomer(ctx context.Context, createCustomerCommand request.CreateCustomerCommand) error
}
