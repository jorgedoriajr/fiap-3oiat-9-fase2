package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	Data      string
	CreatedAt time.Time
}
