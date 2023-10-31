package input

import (
	"context"
	"hamburgueria/internal/modules/customer/usecase/command"
)

type CreateCustomerPort interface {
	AddCustomer(ctx context.Context, createCustomerCommand command.CreateCustomerCommand) error
}
