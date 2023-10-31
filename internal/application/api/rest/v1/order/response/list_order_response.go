package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type ListOrderResponse struct {
	OrderId    uuid.UUID                                  `json:"orderId"`
	Status     string                                     `json:"status"`
	Amount     int                                        `json:"amount"`
	CustomerId string                                     `json:"customerId"`
	CreatedAt  time.Time                                  `json:"createdAt"`
	Products   []*result.FindProductWithIngredientsResult `json:"products"`
}
