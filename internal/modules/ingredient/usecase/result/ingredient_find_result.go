package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type FindIngredientResult struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   string
}

func FromEntity(ingredient entity.IngredientEntity) FindIngredientResult {
	return FindIngredientResult{
		ID:     ingredient.ID,
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   ingredient.Type,
	}
}
