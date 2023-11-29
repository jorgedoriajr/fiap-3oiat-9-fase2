package domain

import (
	"hamburgueria/internal/modules/ingredient/domain"
)

type ProductCategory struct {
	Name                    string
	AcceptCustom            bool
	ConfigByProductCategory []domain.IngredientTypeProductCategory
}
