package command

import "github.com/google/uuid"

type CreateOrderCommand struct {
	CustomerDocument string                       `json:"document"`
	Products         []CreateOrderProductsCommand `json:"products"`
}

type CreateOrderProductsCommand struct {
	Id          uuid.UUID                      `json:"id"`
	Quantity    int                            `json:"quantity"`
	Ingredients []CreateOrderIngredientCommand `json:"ingredients"`
	Type        string                         `json:"type"`
}

type CreateOrderIngredientCommand struct {
	Id       uuid.UUID `json:"id"`
	Quantity int       `json:"quantity"`
}
