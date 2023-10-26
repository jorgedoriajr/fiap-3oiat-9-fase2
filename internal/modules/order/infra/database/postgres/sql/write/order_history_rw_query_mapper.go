package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"time"
)

type InsertOrderHistoryRWQueryMapper struct {
	Id        uuid.UUID `position:"0"`
	OrderId   uuid.UUID `position:"1"`
	Status    string    `position:"2"`
	CreatedAt time.Time `position:"3"`
	ChangeBy  string    `position:"4"`
}

func EntityToInsertOrderHistoryQueryMapper(orderHistory entity.OrderHistory) InsertOrderHistoryRWQueryMapper {
	return InsertOrderHistoryRWQueryMapper{
		Id:        orderHistory.Id,
		OrderId:   orderHistory.OrderId,
		Status:    orderHistory.Status,
		CreatedAt: orderHistory.CreatedAt,
		ChangeBy:  orderHistory.ChangeBy,
	}
}
