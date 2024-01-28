package result

import (
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type ListOrderResult struct {
	OrderNumber int
	Status      string
	Amount      int
	CustomerId  string
	CreatedAt   time.Time
	Products    []OrderProductResult
}

type OrderProductResult struct {
	Name        string
	Number      int
	Amount      int
	Quantity    int
	Ingredients []result.ProductIngredientsResult
}

func ListOrderResultFromDomain(order domain.Order) *ListOrderResult {
	var products []OrderProductResult
	for _, product := range order.Products {
		products = append(products, OrderProductResultFromDomain(product))
	}

	return &ListOrderResult{
		OrderNumber: order.Number,
		Status:      string(order.Status),
		Amount:      order.Amount,
		CustomerId:  order.CustomerId,
		CreatedAt:   order.CreatedAt,
		Products:    products,
	}
}

func OrderProductResultFromDomain(orderProduct domain.OrderProduct) OrderProductResult {
	var ingredients []result.ProductIngredientsResult

	for _, productIngredient := range orderProduct.Product.Ingredients {
		ingredients = append(ingredients, result.FromProductIngredientDomain(productIngredient))
	}
	return OrderProductResult{
		Name:        orderProduct.Product.Name,
		Number:      orderProduct.Product.Number,
		Amount:      orderProduct.Amount,
		Quantity:    orderProduct.Quantity,
		Ingredients: ingredients,
	}
}
