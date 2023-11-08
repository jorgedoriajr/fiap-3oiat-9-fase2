package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain"
)

type InsertOrderProductRWQueryMapper struct {
	Id        uuid.UUID `position:"0"`
	OrderId   uuid.UUID `position:"1"`
	ProductId uuid.UUID `position:"2"`
	Quantity  int       `position:"3"`
	Amount    int       `position:"4"`
}

func EntityToInsertOrderProductQueryMapper(orderProduct domain.OrderProduct) InsertOrderProductRWQueryMapper {
	return InsertOrderProductRWQueryMapper{
		Id:        orderProduct.Id,
		OrderId:   orderProduct.OrderId,
		ProductId: orderProduct.ProductId,
		Quantity:  orderProduct.Quantity,
		Amount:    orderProduct.Amount,
	}
}
