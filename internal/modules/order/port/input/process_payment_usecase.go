package input

import (
	"context"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/usecase/result"
)

type ProcessPaymentUseCasePort interface {
	ProcessPayment(ctx context.Context, order domain.Order) (*result.PaymentCreatedResult, error)
}
