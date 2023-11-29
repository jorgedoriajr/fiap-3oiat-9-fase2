package domain

import "github.com/google/uuid"

type IngredientType struct {
	Name                    string
	ConfigByProductCategory []IngredientTypeProductCategory
}

type IngredientTypeProductCategory struct {
	Id              uuid.UUID
	IngredientType  string
	Optional        bool
	MaxQtd          int
	ProductCategory string
}
