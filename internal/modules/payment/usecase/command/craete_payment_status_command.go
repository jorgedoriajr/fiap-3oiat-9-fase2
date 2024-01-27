package command

import (
	"hamburgueria/internal/modules/payment/domain/valueobjects"

	"github.com/google/uuid"
)

type CreatePaymentStatusCommand struct {
	Id                uuid.UUID
	PaymentId         uuid.UUID
	ExternalReference uuid.UUID
	Status            valueobjects.Status
}
