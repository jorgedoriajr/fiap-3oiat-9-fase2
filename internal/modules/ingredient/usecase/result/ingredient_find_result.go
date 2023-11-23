package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type FindIngredientResult struct {
	ID     uuid.UUID
	Number int
	Name   string
	Amount int
	Type   string
}

func FromDomain(ingredient domain.Ingredient) FindIngredientResult {
	return FindIngredientResult{
		ID:     ingredient.ID,
		Number: ingredient.Number,
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   ingredient.Type.Name,
	}
}
