package domain

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/product/domain"
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

type OrderProduct struct {
	Id       uuid.UUID
	Product  domain.Product
	OrderId  uuid.UUID
	Quantity int
	Amount   int
}
