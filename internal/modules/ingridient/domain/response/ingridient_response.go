package response

import (
	"hamburgueria/internal/modules/ingridient/domain/valueobjects"

	"github.com/google/uuid"
)

type IngridientResponse struct {
	Id     uuid.UUID           `json:"id"`
	Name   valueobjects.Name   `json:"name"`
	Amount valueobjects.Amount `json:"amount"`
	Type   valueobjects.Type   `json:"type"`
}
