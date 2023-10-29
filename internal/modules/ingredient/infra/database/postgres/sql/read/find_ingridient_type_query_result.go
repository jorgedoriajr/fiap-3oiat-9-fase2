package read

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type FindIngredientTypeQueryResult struct {
	Name            string `db: "name"`
	Optional        string `db: "optional"`
	Max_QTD         string `db: "max_qtd"`
	ProductCategory string `db: "product_category"`
}

func (fc FindIngredientTypeQueryResult) ToEntity() *entity.IngredientType {
	return &entity.IngredientType{
		Name:            fc.Name,
		Optional:        fc.Optional,
		Max_QTD:         fc.Max_QTD,
		ProductCategory: fc.ProductCategory,
	}
}

func ToIngredientTypeEntityList(results []FindIngredientTypeQueryResult) []*entity.IngredientType {
	entities := make([]*entity.IngredientType, len(results))

	for i, result := range results {
		entities[i] = result.ToEntity()
	}

	return entities
}
