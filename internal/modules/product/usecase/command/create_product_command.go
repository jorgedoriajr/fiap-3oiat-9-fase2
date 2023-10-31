package command

import (
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/domain/valueobject"
	"time"
)

type CreateProductCommand struct {
	Name        string
	Amount      int
	Description string
	Category    valueobject.ProductCategory
	Menu        bool
	Ingredients []Ingredient
}

type Ingredient struct {
	ID       string
	Quantity int
}

func NewCreateProductCommand(
	Name string,
	Description string,
	Category valueobject.ProductCategory,
	Menu bool,
	Ingredients []Ingredient,
) *CreateProductCommand {

	cmd := &CreateProductCommand{
		Name:        Name,
		Description: Description,
		Category:    Category,
		Menu:        Menu,
		Ingredients: Ingredients,
	}
	//cmd.calculateAmountFromIngredients()
	return cmd
}

//func (c CreateProductCommand) calculateAmountFromIngredients() {
//	var total int
//	for _, ingredient := range c.Ingredients {
//		atomic.Addint(&total, int(ingredient.Amount))
//	}
//	c.Amount = int(total)
//}

type IngredientType string

func GetIngredientTypeByName(name string) IngredientType {
	if t, ok := types[name]; ok {
		return t
	}
	return ""
}

var types = map[string]IngredientType{
	"Protein":           Protein,
	"VegetableAndSalad": VegetableAndSalad,
	"Sauce":             Sauces,
	"Cheese":            Cheeses,
}

const (
	Protein           IngredientType = "Protein"
	VegetableAndSalad IngredientType = "VegetableAndSalad"
	Sauces            IngredientType = "Sauce"
	Cheeses           IngredientType = "Cheese"
)

func (cmd CreateProductCommand) ToProductEntity() entity.ProductEntity {
	return entity.ProductEntity{
		Name:        cmd.Name,
		Amount:      cmd.Amount,
		Description: cmd.Description,
		Category:    cmd.Category,
		Menu:        cmd.Menu,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToIngredientEntity() {

}
