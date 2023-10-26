package entity

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"

	"github.com/google/uuid"
)

type Ingredient struct {
	Id     uuid.UUID
	Name   valueobjects.Name
	Amount valueobjects.Amount
	Type   valueobjects.Type
}
