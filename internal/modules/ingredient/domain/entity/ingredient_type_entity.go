package entity

import "hamburgueria/internal/modules/ingredient/domain/valueobject"

type IngredientTypeEntity struct {
	ID   int
	Name valueobject.IngredientType
}
