package subscriber

import (
	"context"
	"fmt"
	"hamburgueria/internal/modules/common/events"
	"hamburgueria/internal/modules/order/port/output"
	"hamburgueria/pkg/domainevent"
)

type PaymentCreatedEventHandler struct {
	orderPersistence output.OrderPersistencePort
}

func (p PaymentCreatedEventHandler) Handle(ctx context.Context, event domainevent.Event[events.PaymentCreatedDomainEvent]) error {
	fmt.Println(event.GetPayload().Status)
	return nil
}
