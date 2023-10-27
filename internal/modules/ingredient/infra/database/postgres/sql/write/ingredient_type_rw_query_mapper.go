package write

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"
)

type InsertIngredientTypeRWQueryMapper struct {
	Name            valueobjects.Name            `position:"0"`
	Optional        valueobjects.Optional        `position:"1"`
	Max_QTD         valueobjects.Max_QTD         `position:"2"`
	ProductCategory valueobjects.ProductCategory `position:"3"`
}

func ToInsertIngredientTypeQueryMapper(ingredientType entity.IngredientType) InsertIngredientTypeRWQueryMapper {
	return InsertIngredientTypeRWQueryMapper{
		Name:            ingredientType.Name,
		Optional:        ingredientType.Optional,
		Max_QTD:         ingredientType.Max_QTD,
		ProductCategory: ingredientType.ProductCategory,
	}
}
