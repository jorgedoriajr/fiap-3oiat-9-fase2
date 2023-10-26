package entity

import (
	"hamburgueria/internal/modules/ingredient/domain/valueobjects"

	"github.com/google/uuid"
)

type IngredientType struct {
	Id              uuid.UUID
	Name            valueobjects.Name
	Optional        valueobjects.Optional
	Max_QTD         valueobjects.Max_QTD
	ProductCategory valueobjects.ProductCategory
}
