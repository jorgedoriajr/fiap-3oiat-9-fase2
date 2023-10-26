package entity

import (
	"github.com/google/uuid"
)

type OrderProduct struct {
	Id        uuid.UUID
	ProductId uuid.UUID
	OrderId   uuid.UUID
	Quantity  int
	Amount    int64
}
