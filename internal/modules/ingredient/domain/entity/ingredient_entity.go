package entity

import (
	"github.com/google/uuid"
)

type IngredientEntity struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   string
}
