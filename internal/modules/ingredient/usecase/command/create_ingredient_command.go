package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type CreateIngredientCommand struct {
	Name   string
	Amount int
	Type   valueobject.IngredientType
}

func (c CreateIngredientCommand) ToIngredientType() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     uuid.New(),
		Name:   c.Name,
		Amount: c.Amount,
		Type:   c.Type,
	}
}

func NewCreateIngredientCommand(
	Name string,
	Amount int,
	Type valueobject.IngredientType,
) *CreateIngredientCommand {
	return &CreateIngredientCommand{
		Name:   Name,
		Amount: Amount,
		Type:   Type,
	}
}
