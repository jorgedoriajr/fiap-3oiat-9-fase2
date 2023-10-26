package request

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"

	"github.com/google/uuid"
)

type CreateingredientCommand struct {
	Id     uuid.UUID           `json:"id"`
	Name   valueobjects.Name   `json:"name"`
	Amount valueobjects.Amount `json:"amount"`
	Type   valueobjects.Type   `json:"type"`
}
