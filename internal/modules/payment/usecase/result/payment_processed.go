package result

import "github.com/google/uuid"

type PaymentProcessed struct {
	PaymentId   uuid.UUID
	OrderId     uuid.UUID
	PaymentData string
}
