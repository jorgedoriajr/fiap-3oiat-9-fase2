package command

import "github.com/google/uuid"

type CreatePaymentStatusCommand struct {
	Id                uuid.UUID
	ExternalReference uuid.UUID
	Status            Status
}

type Status string

const (
	Created  Status = "created"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
