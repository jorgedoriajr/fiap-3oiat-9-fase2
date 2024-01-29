package command

import "github.com/google/uuid"

type CreatePaymentCommand struct {
	Amount     int
	OrderId    uuid.UUID
	OrderItems []OrderItem
}

type OrderItem struct {
	Name        string
	Amount      int
	Quantity    int
	TotalAmount int
}
