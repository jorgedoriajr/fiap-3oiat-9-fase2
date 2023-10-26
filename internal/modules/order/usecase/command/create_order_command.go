package command

import "github.com/google/uuid"

type CreateOrderCommand struct {
	CustomerDocument string
	Products         []CreateOrderProductsCommand
}

type CreateOrderProductsCommand struct {
	Id          uuid.UUID
	Quantity    int
	Ingredients []CreateOrderIngredientCommand
	Type        string
}

type CreateOrderIngredientCommand struct {
	Id       uuid.UUID
	Quantity int
}
