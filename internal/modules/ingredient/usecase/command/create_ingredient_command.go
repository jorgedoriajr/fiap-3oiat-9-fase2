package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type CreateIngredientCommand struct {
	Name   string
	Amount int
	Type   string
}

func (c CreateIngredientCommand) ToIngredientEntity(ingredientType domain.IngredientType) *domain.Ingredient {
	return &domain.Ingredient{
		ID:     uuid.New(),
		Name:   c.Name,
		Amount: c.Amount,
		Type:   ingredientType,
	}
}

func NewCreateIngredientCommand(
	Name string,
	Amount int,
	Type string,
) *CreateIngredientCommand {
	return &CreateIngredientCommand{
		Name:   Name,
		Amount: Amount,
		Type:   Type,
	}
}
