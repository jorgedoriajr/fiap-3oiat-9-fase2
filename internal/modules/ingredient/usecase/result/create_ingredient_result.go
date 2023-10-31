package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/application/api/rest/v1/ingredient/response"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type CreateIngredientResult struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   string
}

func ToCreateIngredientResultFrom(entity entity.IngredientEntity) CreateIngredientResult {
	return CreateIngredientResult{
		ID:     entity.ID,
		Name:   entity.Name,
		Amount: entity.Amount,
		Type:   entity.Type,
	}
}

func (c CreateIngredientResult) ToResponse() response.IngredientCreatedResponse {
	return response.IngredientCreatedResponse{
		Name:   c.Name,
		Amount: c.Amount,
		Type:   c.Type,
	}
}
