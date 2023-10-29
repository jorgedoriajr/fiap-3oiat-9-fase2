package command

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"

	"github.com/google/uuid"
)

type CreateIngredientCommand struct {
	Name   string
	Amount string
	Type   string
}

func (cmd CreateIngredientCommand) ToIngredientEntity() entity.Ingredient {
	return entity.Ingredient{
		Id:     uuid.New(),
		Name:   cmd.Name,
		Amount: cmd.Amount,
		Type:   cmd.Type,
	}
}
