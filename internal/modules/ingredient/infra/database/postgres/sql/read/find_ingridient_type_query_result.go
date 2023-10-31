package read

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type FindIngredientTypeQueryResult struct {
	Name                              string `db:"name"`
	IngredientsTypePerProductCategory []FindIngredientsTypeProductCategory
}

type FindIngredientsTypeProductCategory struct {
	Optional        string `db:"optional"`
	MaxQtd          string `db:"max_qtd"`
	ProductCategory string `db:"product_category"`
}

func (fc FindIngredientTypeQueryResult) ToEntity() entity.IngredientType {
	var ingredientsTypePerProductCategory []entity.IngredientTypeProductCategory
	for _, ingredientTypePerProductCategory := range fc.IngredientsTypePerProductCategory {
		ingredientsTypePerProductCategory = append(ingredientsTypePerProductCategory, entity.IngredientTypeProductCategory{
			Optional:        ingredientTypePerProductCategory.Optional,
			MaxQtd:          ingredientTypePerProductCategory.MaxQtd,
			ProductCategory: ingredientTypePerProductCategory.ProductCategory,
		})
	}

	return entity.IngredientType{
		Name:                    fc.Name,
		ConfigByProductCategory: ingredientsTypePerProductCategory,
	}
}

func ToIngredientTypeEntityList(results []FindIngredientTypeQueryResult) []entity.IngredientType {
	entities := make([]entity.IngredientType, len(results))

	for i, result := range results {
		entities[i] = result.ToEntity()
	}

	return entities
}
