package result

import (
	"time"

	"github.com/google/uuid"
)

type PaymentProcessed struct {
	PaymentId   uuid.UUID
	OrderId     uuid.UUID
	PaymentData string
	CreatedAt   time.Time
}
