package response

import (
	"hamburgueria/internal/modules/payment/domain"

	"github.com/google/uuid"
)

type QrCodePaymentResponse struct {
	QrData         string    `json:"qr_data"`
	InStoreOrderId uuid.UUID `json:"in_store_order_id"`
}

func (qcpr QrCodePaymentResponse) MpQrCodeResponseToPaymentEntity() *domain.Payment {
	return &domain.Payment{
		Id:   qcpr.InStoreOrderId,
		Data: qcpr.QrData,
	}
}
