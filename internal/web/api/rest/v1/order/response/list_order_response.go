package response

import (
	"time"
)

type ListOrderResponse struct {
	OrderNumber int                 `json:"orderNumber"`
	Status      string              `json:"status"`
	Amount      int                 `json:"amount"`
	CustomerId  string              `json:"customerId"`
	CreatedAt   time.Time           `json:"createdAt"`
	Products    []ListOrderProducts `json:"products"`
}

type ListOrderProducts struct {
	Name        string                         `json:"name"`
	Number      int                            `json:"number"`
	Amount      int                            `json:"amount"`
	Quantity    int                            `json:"quantity"`
	Ingredients []ListOrderProductsIngredients `json:"ingredients"`
}

type ListOrderProductsIngredients struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity"`
}
