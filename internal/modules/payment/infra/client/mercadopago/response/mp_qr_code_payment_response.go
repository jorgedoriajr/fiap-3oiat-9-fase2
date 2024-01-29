package response

import (
	"hamburgueria/internal/modules/payment/domain"
	"time"

	"github.com/google/uuid"
)

type QrCodePaymentResponse struct {
	QrData         string    `json:"qr_data"`
	InStoreOrderId uuid.UUID `json:"in_store_order_id"`
}

func (q QrCodePaymentResponse) MpQrCodeResponseToPaymentEntity(orderId uuid.UUID) domain.Payment {
	return domain.Payment{
		Id:        q.InStoreOrderId,
		OrderId:   orderId,
		Data:      q.QrData,
		CreatedAt: time.Now(),
	}
}
