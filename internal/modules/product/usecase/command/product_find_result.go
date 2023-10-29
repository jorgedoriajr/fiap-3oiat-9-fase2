package command

import (
	"github.com/google/uuid"
	"time"
)

type FindProductResult struct {
	ID          uuid.UUID
	Name        string
	Number      int
	Amount      int
	Description string
	Category    string
	Menu        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FindProductWithIngredientsResult struct {
	ID          uuid.UUID
	Name        string
	Number      int
	Amount      int
	Description string
	Category    string
	Menu        bool
	Ingredients []any
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
