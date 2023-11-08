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
	productId uuid.UUID,
	ingredients []domain.ProductIngredient,
	amount int,
) domain.Product {
	return domain.Product{
		ID:          productId,
		Name:        cmd.Name,
		Amount:      amount,
		Description: cmd.Description,
		Category:    domain.ProductCategory{Name: cmd.Category},
		Menu:        cmd.Menu,
		ImgPath:     cmd.ImgPath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Ingredients: ingredients,
	}
}
