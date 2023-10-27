package result

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"
)

type CreateIngredientTypeResult struct {
	Name            valueobjects.Name
	Optional        valueobjects.Optional
	Max_QTD         valueobjects.Max_QTD
	ProductCategory valueobjects.ProductCategory
}
