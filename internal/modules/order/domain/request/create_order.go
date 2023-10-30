package request

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/usecase/command"
)

type CreateOrder struct {
	CustomerDocument string                `json:"document"`
	Products         []CreateOrderProducts `json:"products"`
}

type CreateOrderProducts struct {
	Id          uuid.UUID               `json:"id"`
	Quantity    int                     `json:"quantity"`
	Ingredients []CreateOrderIngredient `json:"ingredients"`
	Type        string                  `json:"type"`
}

type CreateOrderIngredient struct {
	Id       uuid.UUID `json:"id"`
	Quantity int       `json:"quantity"`
}

func (c CreateOrder) ToCommand() command.CreateOrderCommand {
	var productsCommand []command.CreateOrderProductsCommand

	for _, product := range c.Products {
		var ingredientsCommand []command.CreateOrderIngredientCommand
		for _, ingredient := range product.Ingredients {
			ingredientsCommand = append(ingredientsCommand, command.CreateOrderIngredientCommand{
				Id:       ingredient.Id,
				Quantity: ingredient.Quantity,
			})
		}

		productsCommand = append(productsCommand, command.CreateOrderProductsCommand{
			Id:          product.Id,
			Quantity:    product.Quantity,
			Ingredients: ingredientsCommand,
			Type:        product.Type,
		})
	}

	return command.CreateOrderCommand{
		CustomerDocument: c.CustomerDocument,
		Products:         productsCommand,
	}
}