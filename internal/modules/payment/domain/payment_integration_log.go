package domain

import (
	"hamburgueria/internal/modules/payment/domain/valueobjects"

	"github.com/google/uuid"
)

type PaymentIntegrationLog struct {
	Id                   uuid.UUID
	PaymentId            uuid.UUID
	PaymentIntegrationId uuid.UUID
	PaymentStatus        valueobjects.Status
}
