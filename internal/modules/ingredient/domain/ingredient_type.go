package domain

type IngredientType struct {
	Name                    string
	ConfigByProductCategory []IngredientTypeProductCategory
}

type IngredientTypeProductCategory struct {
	Optional        string
	MaxQtd          string
	ProductCategory string
}
