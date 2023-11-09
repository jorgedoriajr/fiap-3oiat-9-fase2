package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain"
	"time"
)

type CreateProductCommand struct {
	Name        string
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	Ingredients []Ingredient
}

type Ingredient struct {
	Number   int
	Quantity int
}

func NewCreateProductCommand(
	Name string,
	Description string,
	Category string,
	Menu bool,
	Ingredients []Ingredient,
	ImgPath string,
) *CreateProductCommand {

	cmd := &CreateProductCommand{
		Name:        Name,
		Description: Description,
		Category:    Category,
		Menu:        Menu,
		ImgPath:     ImgPath,
		Ingredients: Ingredients,
	}
	return cmd
}

func (cmd CreateProductCommand) ToProductDomain(
	ingredients []domain.ProductIngredient,
	amount int,
	productID uuid.UUID,
	category domain.ProductCategory,
) domain.Product {
	return domain.Product{
		ID:          productID,
		Name:        cmd.Name,
		Amount:      amount,
		Description: cmd.Description,
		Category:    category,
		Menu:        cmd.Menu,
		ImgPath:     cmd.ImgPath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Ingredients: ingredients,
	}
}
