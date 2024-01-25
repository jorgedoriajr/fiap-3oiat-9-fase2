package domain

import "github.com/google/uuid"

type PaymentIntegrationLog struct {
	Id                   uuid.UUID
	PaymentIntegrationId uuid.UUID
	PaymentStatus        Status
}

type Status string

const (
	Created  Status = "created"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
