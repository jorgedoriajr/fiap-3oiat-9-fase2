package entity

import (
	"github.com/google/uuid"
	"time"
)

type ProductEntity struct {
	ID          uuid.UUID
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	Ingredients []uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
