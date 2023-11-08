package model

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"time"
)

type OrderHistory struct {
	ID        uuid.UUID
	OrderId   uuid.UUID
	Status    string
	ChangeBy  string
	CreatedAt time.Time
}

func (o OrderHistory) ToDomain() domain.OrderHistory {
	return domain.OrderHistory{
		Id:        o.ID,
		OrderId:   o.OrderId,
		Status:    valueobject.OrderStatus(o.Status),
		ChangeBy:  o.ChangeBy,
		CreatedAt: o.CreatedAt,
	}
}
