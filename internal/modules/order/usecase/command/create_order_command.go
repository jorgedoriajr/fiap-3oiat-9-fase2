package command

import "github.com/google/uuid"

type CreateOrderCommand struct {
	CustomerDocument string
	Products         []CreateOrderProductsCommand
}

type CreateOrderProductsCommand struct {
	Number          int
	Quantity        int
	Ingredients     []CreateOrderIngredientCommand
	Type            string
	ProductCategory string
}

type CreateOrderIngredientCommand struct {
	Id       uuid.UUID
	Quantity int
}
