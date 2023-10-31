package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/usecase/result"
	"time"
)

type ListOrderResponse struct {
	OrderId    uuid.UUID           `json:"orderId"`
	Status     string              `json:"status"`
	Amount     int                 `json:"amount"`
	CustomerId string              `json:"customerId"`
	CreatedAt  time.Time           `json:"createdAt"`
	Products   []ListOrderProducts `json:"products"`
}

type ListOrderProducts struct {
	Name        string                         `json:"name"`
	Number      int                            `json:"number"`
	Amount      int                            `json:"amount"`
	Ingredients []ListOrderProductsIngredients `json:"ingredients"`
}

type ListOrderProductsIngredients struct {
	Name        string `json:"name"`
	TotalAmount int    `json:"totalAmount"`
	Quantity    int    `json:"quantity"`
}

func FromResult(resultOrders []result.ListOrderResult) []ListOrderResponse {
	var ordersResponse []ListOrderResponse
	for _, order := range resultOrders {
		var productsResponse []ListOrderProducts
		for _, product := range order.Products {
			var ingredientsResponse []ListOrderProductsIngredients
			for _, ingredient := range product.Ingredients {
				ingredientsResponse = append(ingredientsResponse, ListOrderProductsIngredients{
					Name:        ingredient.Name,
					TotalAmount: ingredient.TotalAmount,
					Quantity:    ingredient.Quantity,
				})
			}
			productsResponse = append(productsResponse, ListOrderProducts{
				Name:        product.Name,
				Number:      product.Number,
				Amount:      product.Amount,
				Ingredients: ingredientsResponse,
			})
		}

		ordersResponse = append(ordersResponse, ListOrderResponse{
			OrderId:    order.OrderId,
			Status:     order.Status,
			Amount:     order.Amount,
			CustomerId: order.CustomerId,
			CreatedAt:  order.CreatedAt,
			Products:   productsResponse,
		})
	}
	return ordersResponse
}
