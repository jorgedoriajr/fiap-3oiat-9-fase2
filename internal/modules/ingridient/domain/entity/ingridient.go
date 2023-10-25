package entity

import (
	"hamburgueria/internal/modules/ingridient/domain/valueobjects"
)

type Ingridients struct {
	Name   valueobjects.Name   `json:"name"`
	Amount valueobjects.Amount `json:"amount"`
	Type   valueobjects.Type   `json:"type"`
}
