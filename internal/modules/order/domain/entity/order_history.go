package entity

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"time"
)

type OrderHistory struct {
	Id        uuid.UUID
	OrderId   uuid.UUID
	Status    valueobject.OrderStatus
	ChangeBy  string
	CreatedAt time.Time
}
