package entity

import (
	"github.com/google/uuid"
	"time"
)

type ProductEntity struct {
	ID          uuid.UUID
	Number      int
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
