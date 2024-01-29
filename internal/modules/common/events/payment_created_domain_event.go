package events

import (
	"github.com/google/uuid"
)

type PaymentCreatedDomainEvent struct {
	OrderID   uuid.UUID
	Amount    int
	Status    string
	PaymentID uuid.UUID
}

func (p PaymentCreatedDomainEvent) IsSuccess() bool {
	return p.Status == "SUCCESS"
}
