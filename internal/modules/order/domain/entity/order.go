package entity

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"time"
)

type Order struct {
	Id         uuid.UUID
	CustomerId string
	Products   []OrderProduct
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     valueobject.OrderStatus
	Amount     int
}
