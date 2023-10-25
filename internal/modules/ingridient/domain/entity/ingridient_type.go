package entity

import (
	"hamburgueria/internal/modules/ingridient/domain/valueobjects"

	"github.com/google/uuid"
)

type IngridientType struct {
	Id              uuid.UUID                    `json:"id"`
	Name            valueobjects.Name            `json:"name"`
	Optional        valueobjects.Optional        `json:"optional"`
	Max_QTD         valueobjects.Max_QTD         `json:"max_qtd"`
	ProductCategory valueobjects.ProductCategory `json:"product_category"`
}
