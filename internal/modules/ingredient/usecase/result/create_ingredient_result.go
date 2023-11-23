package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type CreateIngredientResult struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   string
}

func ToCreateIngredientResultFrom(entity domain.Ingredient) CreateIngredientResult {
	return CreateIngredientResult{
		ID:     entity.ID,
		Name:   entity.Name,
		Amount: entity.Amount,
		Type:   entity.Type.Name,
	}
}
