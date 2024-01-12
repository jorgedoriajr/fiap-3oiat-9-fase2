package presenter

import (
	"hamburgueria/internal/modules/order/usecase/result"
	"hamburgueria/internal/server/api/rest/v1/order/response"
)

func OrderResponseFromResult(result result.CreateOrderResult) response.OrderResponse {
	return response.OrderResponse{
		Amount:      result.Amount,
		PaymentData: result.PaymentData,
	}
}

func ListOrderResponseFromResult(resultOrders []result.ListOrderResult) []response.ListOrderResponse {
	var ordersResponse []response.ListOrderResponse
	for _, order := range resultOrders {
		var productsResponse []response.ListOrderProducts
		for _, product := range order.Products {
			var ingredientsResponse []response.ListOrderProductsIngredients
			for _, ingredient := range product.Ingredients {
				ingredientsResponse = append(ingredientsResponse, response.ListOrderProductsIngredients{
					Name:     ingredient.Name,
					Amount:   ingredient.Amount,
					Quantity: ingredient.Quantity,
				})
			}
			productsResponse = append(productsResponse, response.ListOrderProducts{
				Name:        product.Name,
				Number:      product.Number,
				Amount:      product.Amount,
				Quantity:    product.Quantity,
				Ingredients: ingredientsResponse,
			})
		}

		ordersResponse = append(ordersResponse, response.ListOrderResponse{
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
