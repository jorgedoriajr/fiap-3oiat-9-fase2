package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type ListOrderResult struct {
	OrderId    uuid.UUID
	Status     string
	Amount     int
	CustomerId string
	CreatedAt  time.Time
	Products   []*result.FindProductWithIngredientsResult
}
