package output

import (
	"context"
	"hamburgueria/internal/modules/customer/domain"
)

type CustomerPersistencePort interface {
	Create(ctx context.Context, customer domain.Customer) error
	Get(ctx context.Context, document string) (*domain.Customer, error)
}
