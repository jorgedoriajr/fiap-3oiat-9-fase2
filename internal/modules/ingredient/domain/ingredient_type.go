package domain

type IngredientType struct {
	Name                    string
	ConfigByProductCategory []IngredientTypeProductCategory
}

type IngredientTypeProductCategory struct {
	IngredientType  string
	Optional        string
	MaxQtd          string
	ProductCategory string
}
