package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/usecase/result"
)

type ProductCreatedResponse struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Name   string    `json:"name"`
	Amount int       `json:"amount"`
	Type   string    `json:"type"`
}

func (pcr ProductCreatedResponse) From(entity entity.IngredientEntity) ProductCreatedResponse {
	return ProductCreatedResponse{
		ID:     entity.ID,
		Name:   entity.Name,
		Amount: entity.Amount,
		Type:   entity.Type,
	}
}

func (pcr ProductCreatedResponse) FromResult(result result.CreateIngredientResult) ProductCreatedResponse {
	return ProductCreatedResponse{
		Name:   result.Name,
		Amount: result.Amount,
		Type:   result.Type,
	}
}
