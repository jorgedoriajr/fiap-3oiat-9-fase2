package input

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/order/domain/valueobject"
)

type UpdateOrderPort interface {
	Update(
		ctx context.Context,
		orderId uuid.UUID,
		status valueobject.OrderStatus,
		paymentId *uuid.UUID,
	) error
}
