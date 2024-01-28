package request

import (
	"hamburgueria/internal/modules/payment/domain/valueobjects"
	"hamburgueria/internal/modules/payment/usecase/command"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CreatePaymentStatusRequest struct {
	Id          uuid.UUID  `json:"id"`
	LiveMode    bool       `json:"live_mode"`
	Type        string     `json:"type"`
	DateCreated time.Time  `json:"date_created"`
	UserID      int        `json:"user_id"`
	APIVersion  string     `json:"api_version"`
	Action      string     `json:"action"`
	Data        DataStatus `json:"data"`
}
type DataStatus struct {
	Id uuid.UUID `json:"id"`
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
		Id:                uuid.New(),
		ExternalReference: c.Data.Id,
		PaymentId:         c.Id,
		Status:            valueobjects.Status(c.GetPaymentStatus()),
	}
}
