package entity

import (
	"hamburgueria/internal/modules/ingridient/domain/valueobjects"

	"github.com/google/uuid"
)

type Ingridients struct {
	Id     uuid.UUID
	Name   valueobjects.Name
	Amount valueobjects.Amount
	Type   valueobjects.Type
}
