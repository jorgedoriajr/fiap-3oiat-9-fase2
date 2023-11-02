package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"time"
)

type FindOrderQueryResult struct {
	Id         uuid.UUID
	CustomerId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
	Amount     int
}

func (fc FindOrderQueryResult) ToEntity() entity.Order {
	return entity.Order{
		Id:         fc.Id,
		CustomerId: fc.CustomerId,
		CreatedAt:  fc.CreatedAt,
		UpdatedAt:  fc.UpdatedAt,
		Status:     valueobject.OrderStatus(fc.Status),
		Amount:     fc.Amount,
	}
}
