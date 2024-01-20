package command

import "github.com/google/uuid"

type CreatePaymentCommand struct {
	Amount  int
	OrderId uuid.UUID
}
