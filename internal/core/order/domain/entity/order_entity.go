package entity

import "hamburgueria/internal/core/order/domain/valueobjects"

type OrderEntity struct {
	OrderID  string
	products []valueobjects.Product
}
