package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	Data      []byte
	CreatedAt time.Time
}
