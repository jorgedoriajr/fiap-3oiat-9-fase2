package result

import (
	"hamburgueria/internal/modules/payment/domain/valueobjects"

	"github.com/google/uuid"
)

type PaymentStatusProcessed struct {
	Id                   uuid.UUID
	PaymentId            uuid.UUID
	PaymentIntegrationId uuid.UUID
	Status               valueobjects.Status
}
