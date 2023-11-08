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
	Products   []OrderProduct
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     string
	Amount     int
}

type OrderProduct struct {
	ID       uuid.UUID
	Product  model.Product
	OrderId  uuid.UUID
	Quantity int
	Amount   int
}

func (o Order) ToDomain() *domain.Order {
	var products []domain.OrderProduct
	for _, orderProduct := range o.Products {
		products = append(products, *orderProduct.ToDomain())
	}
	return &domain.Order{
		Id:         o.ID,
		CustomerId: o.CustomerId,
		Products:   products,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
		Status:     valueobject.OrderStatus(o.Status),
		Amount:     o.Amount,
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
