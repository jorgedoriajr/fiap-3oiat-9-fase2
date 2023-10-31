package result

import (
	"github.com/google/uuid"
	"time"
)

type FindIngredientResult struct {
	ID        uuid.UUID
	Name      string
	Amount    int
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
