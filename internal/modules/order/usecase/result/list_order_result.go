package result

import (
	"github.com/google/uuid"
	"time"
)

type ListOrderResult struct {
	OrderId    uuid.UUID
	Status     string
	Amount     int
	CustomerId string
	CreatedAt  time.Time
}
