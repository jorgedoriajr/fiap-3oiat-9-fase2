package command

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"
)

type CreateIngredientTypeCommand struct {
	Name            valueobjects.Name
	Optional        valueobjects.Optional
	Max_QTD         valueobjects.Max_QTD
	ProductCategory valueobjects.ProductCategory
}

func (cmd CreateIngredientTypeCommand) ToIngredientTypeEntity() entity.IngredientType {
	return entity.IngredientType{
		Name:            cmd.Name,
		Optional:        cmd.Optional,
		Max_QTD:         cmd.Max_QTD,
		ProductCategory: cmd.ProductCategory,
	}
}
