package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type FindIngredientResult struct {
	ID     uuid.UUID
	Number int
	Name   string
	Amount int
	Type   string
}

func FromEntity(ingredient entity.IngredientEntity) FindIngredientResult {
	return FindIngredientResult{
		ID:     ingredient.ID,
		Number: ingredient.Number,
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   ingredient.Type,
	}
}
