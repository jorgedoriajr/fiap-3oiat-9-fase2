package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
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
	Ingredients []entity.IngredientEntity
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type IngredientEntity struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   valueobject.IngredientType
}
