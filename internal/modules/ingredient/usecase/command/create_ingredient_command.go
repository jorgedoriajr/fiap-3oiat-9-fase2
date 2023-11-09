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

func (c CreateIngredientCommand) ToIngredientEntity() *domain.Ingredient {
	return &domain.Ingredient{
		ID:     uuid.New(),
		Name:   c.Name,
		Amount: c.Amount,
		Type:   c.Type,
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
