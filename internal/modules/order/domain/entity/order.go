package entity

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id         uuid.UUID
	CustomerId string
	Products   []OrderProduct
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
	Amount     int64
}
