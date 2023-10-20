package entity

import (
	"hamburgueria/internal/modules/order/domain/valueobjects"
)

type OrderEntity struct {
	OrderID  string
	products []valueobjects.Product
}
