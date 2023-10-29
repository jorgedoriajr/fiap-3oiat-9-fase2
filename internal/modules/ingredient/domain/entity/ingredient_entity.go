package entity

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type IngredientEntity struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   valueobject.IngredientType
}
