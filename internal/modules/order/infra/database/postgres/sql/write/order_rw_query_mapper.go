package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/entity"
	"time"
)

type InsertOrderRWQueryMapper struct {
	Id               uuid.UUID `position:"0"`
	CustomerDocument string    `position:"1"`
	Amount           int       `position:"2"`
	Status           string    `position:"3"`
	CreatedAt        time.Time `position:"4"`
	UpdatedAt        time.Time `position:"5"`
}

func EntityToInsertOrderQueryMapper(order entity.Order) InsertOrderRWQueryMapper {
	return InsertOrderRWQueryMapper{
		Id:               order.Id,
		CustomerDocument: order.CustomerId,
		Amount:           order.Amount,
		Status:           order.Status,
		CreatedAt:        order.CreatedAt,
		UpdatedAt:        order.UpdatedAt,
	}
}
