package result

import (
	"github.com/google/uuid"
	"time"
)

type CreateProductResult struct {
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
