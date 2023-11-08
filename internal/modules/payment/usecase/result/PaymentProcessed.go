package result

import "github.com/google/uuid"

type PaymentProcessed struct {
	OrderReference uuid.UUID
	PaymentId      uuid.UUID
	PaymentData    string
}
