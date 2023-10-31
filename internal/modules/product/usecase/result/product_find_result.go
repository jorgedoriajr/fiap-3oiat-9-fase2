package result

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
	ImgPath     string
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
	ImgPath     string
	Ingredients []FindProductsIngredientsResult
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FindProductsIngredientsResult struct {
	ID       uuid.UUID
	Name     string
	Amount   int
	Type     string
	Quantity int
}
