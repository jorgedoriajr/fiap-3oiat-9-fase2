package write

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type InsertIngredientTypeRWQueryMapper struct {
	Name            string `position:"0"`
	Optional        string `position:"1"`
	Max_QTD         string `position:"2"`
	ProductCategory string `position:"3"`
}

func ToInsertIngredientTypeQueryMapper(ingredientType entity.IngredientType) InsertIngredientTypeRWQueryMapper {
	return InsertIngredientTypeRWQueryMapper{
		Name:            ingredientType.Name,
		Optional:        ingredientType.Optional,
		Max_QTD:         ingredientType.Max_QTD,
		ProductCategory: ingredientType.ProductCategory,
	}
}
