package model

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/product/infra/database/model"
	"time"
)

type Order struct {
	ID         uuid.UUID
	CustomerId string
	Products   []OrderProduct `gorm:"foreignKey:OrderId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
	Amount     int
	History    []OrderHistory `gorm:"foreignKey:OrderId"`
	PaymentId  uuid.UUID
}

type OrderProduct struct {
	ID       uuid.UUID
	Product  model.Product `gorm:"foreignKey:ID"`
	OrderId  uuid.UUID
	Quantity int
	Amount   int
}

func (o Order) ToDomain() *domain.Order {
	var products []domain.OrderProduct
	for _, orderProduct := range o.Products {
		products = append(products, *orderProduct.ToDomain())
	}
	var orderHistory []domain.OrderHistory
	for _, history := range o.History {
		orderHistory = append(orderHistory, history.ToDomain())
	}
	return &domain.Order{
		Id:         o.ID,
		CustomerId: o.CustomerId,
		Products:   products,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
		Status:     valueobject.OrderStatus(o.Status),
		Amount:     o.Amount,
		History:    orderHistory,
		PaymentId:  o.PaymentId,
	}
}

func (o OrderProduct) ToDomain() *domain.OrderProduct {
	return &domain.OrderProduct{
		Id:       o.ID,
		Product:  *o.Product.ToDomain(),
		OrderId:  o.OrderId,
		Quantity: o.Quantity,
		Amount:   o.Amount,
	}
}

func FromDomain(order domain.Order) *Order {
	var orderProducts []OrderProduct
	for _, product := range order.Products {
		orderProducts = append(orderProducts, OrderProductFromDomain(product))
	}
	return &Order{
		ID:         order.Id,
		CustomerId: order.CustomerId,
		Products:   orderProducts,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Status:     string(order.Status),
		Amount:     order.Amount,
		PaymentId:  order.PaymentId,
	}
}

func OrderProductFromDomain(orderProduct domain.OrderProduct) OrderProduct {
	return OrderProduct{
		ID:       orderProduct.Id,
		Product:  model.ProductFromDomain(orderProduct.Product),
		OrderId:  orderProduct.OrderId,
		Quantity: orderProduct.Quantity,
		Amount:   orderProduct.Amount,
	}
}
