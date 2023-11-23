package domain

import (
	"github.com/google/uuid"
)

type Ingredient struct {
	ID     uuid.UUID
	Number int
	Name   string
	Amount int
	Type   IngredientType
}
