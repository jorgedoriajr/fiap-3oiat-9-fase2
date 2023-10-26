package entity

import (
	"hamburgueria/internal/modules/ingridient/domain/valueobjects"

	"github.com/google/uuid"
)

type IngridientType struct {
	Id              uuid.UUID
	Name            valueobjects.Name
	Optional        valueobjects.Optional
	Max_QTD         valueobjects.Max_QTD
	ProductCategory valueobjects.ProductCategory
}
