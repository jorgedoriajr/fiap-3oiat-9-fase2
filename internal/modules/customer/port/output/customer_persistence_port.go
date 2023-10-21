package output

import (
	"context"
	"hamburgueria/internal/modules/customer/domain/entity"
)

type CustomerPersistencePort interface {
	Create(ctx context.Context, customer entity.Customer) error
	Get(ctx context.Context, document string) (customerResult *entity.Customer, err error)
}
