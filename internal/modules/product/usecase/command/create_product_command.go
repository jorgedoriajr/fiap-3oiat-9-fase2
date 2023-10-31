package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
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
	ID       string
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

func (cmd CreateProductCommand) ToProductEntity(
	productId uuid.UUID,
	ingredients []entity.ProductIngredientEntity,
	amount int,
) entity.ProductEntity {
	return entity.ProductEntity{
		ID:          productId,
		Name:        cmd.Name,
		Amount:      amount,
		Description: cmd.Description,
		Category:    cmd.Category,
		Menu:        cmd.Menu,
		ImgPath:     cmd.ImgPath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Ingredients: ingredients,
	}
}
