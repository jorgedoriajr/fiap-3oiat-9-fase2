package command

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type CreateProductCommand struct {
	Name        string
	Amount      int
	Description string
	Category    string
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
	Category string,
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

func (cmd CreateProductCommand) ToProductEntity() entity.ProductEntity {
	return entity.ProductEntity{
		ID:          uuid.New(),
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
