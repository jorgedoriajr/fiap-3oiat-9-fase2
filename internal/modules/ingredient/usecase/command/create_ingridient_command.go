package command

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"

	"github.com/google/uuid"
)

type CreateingredientCommand struct {
	Name   valueobjects.Name
	Amount valueobjects.Amount
	Type   valueobjects.Type
}

func (cmd CreateingredientCommand) ToingredientEntity() entity.Ingredient {
	return entity.Ingredient{
		Id:     uuid.New(),
		Name:   cmd.Name,
		Amount: cmd.Amount,
		Type:   cmd.Type,
	}
}
