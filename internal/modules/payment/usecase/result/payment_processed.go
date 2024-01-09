package result

import "github.com/google/uuid"

type PaymentProcessed struct {
	PaymentId   uuid.UUID
	PaymentData string
}
