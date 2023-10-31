package entity

import (
	"github.com/google/uuid"
	"time"
)

type OrderHistory struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	Status    string
	ChangeBy  string
	CreatedAt time.Time
}
