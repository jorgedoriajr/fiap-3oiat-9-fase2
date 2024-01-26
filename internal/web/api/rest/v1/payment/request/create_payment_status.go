package request

import (
	"hamburgueria/internal/modules/payment/usecase/command"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CreatePaymentStatusRequest struct {
	ID          uuid.UUID `json:"id"`
	LiveMode    bool      `json:"live_mode"`
	Type        string    `json:"type"`
	DateCreated time.Time `json:"date_created"`
	UserID      int       `json:"user_id"`
	APIVersion  string    `json:"api_version"`
	Action      string    `json:"action"`
	Data        Data      `json:"data"`
}

type Data struct {
	ID uuid.UUID `json:"id"`
}

func (c *CreatePaymentStatusRequest) GetPaymentStatus() string {
	const prefix = "payment."
	if strings.HasPrefix(c.Action, prefix) {
		return strings.TrimPrefix(c.Action, prefix)
	}
	return ""
}

func (c CreatePaymentStatusRequest) ToCommand() command.CreatePaymentStatusCommand {
	return command.CreatePaymentStatusCommand{
		Id:                c.ID,
		ExternalReference: c.Data.ID,
		Status:            command.Status(c.GetPaymentStatus()),
	}
}
