package request

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
	"hamburgueria/internal/modules/ingredient/usecase/command"
)

type CreateIngredientRequest struct {
	Name   string `json:"name" validator:"required"`
	Amount int    `json:"amount" validator:"required"`
	Type   string `json:"type" validator:"required,ingredientType"`
}

func (cp CreateIngredientRequest) ToCommand() command.CreateIngredientCommand {
	return *command.NewCreateIngredientCommand(
		cp.Name, cp.Amount, valueobject.IngredientType(cp.Type),
	)
}
