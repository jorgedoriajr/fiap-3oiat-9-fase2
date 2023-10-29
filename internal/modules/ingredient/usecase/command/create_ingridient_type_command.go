package command

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type CreateIngredientTypeCommand struct {
	Name            string
	Optional        string
	Max_QTD         string
	ProductCategory string
}

func (cmd CreateIngredientTypeCommand) ToIngredientTypeEntity() entity.IngredientType {
	return entity.IngredientType{
		Name:            cmd.Name,
		Optional:        cmd.Optional,
		Max_QTD:         cmd.Max_QTD,
		ProductCategory: cmd.ProductCategory,
	}
}
