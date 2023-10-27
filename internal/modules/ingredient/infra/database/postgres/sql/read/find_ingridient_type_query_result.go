package read

import (
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"
)

type FindIngredientTypeQueryResult struct {
	Name            valueobjects.Name            `db: "name"`
	Optional        valueobjects.Optional        `db: "optional"`
	Max_QTD         valueobjects.Max_QTD         `db: "max_qtd"`
	ProductCategory valueobjects.ProductCategory `db: "product_category"`
}

func (fc FindIngredientTypeQueryResult) ToEntity() *entity.IngredientType {
	return &entity.IngredientType{
		Name:            fc.Name,
		Optional:        fc.Optional,
		Max_QTD:         fc.Max_QTD,
		ProductCategory: fc.ProductCategory,
	}
}
