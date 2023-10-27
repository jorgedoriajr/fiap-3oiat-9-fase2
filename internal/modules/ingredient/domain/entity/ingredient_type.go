package entity

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"
)

type IngredientType struct {
	Name            valueobjects.Name
	Optional        valueobjects.Optional
	Max_QTD         valueobjects.Max_QTD
	ProductCategory valueobjects.ProductCategory
}
