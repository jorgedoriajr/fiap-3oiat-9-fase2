package response

import "hamburgueria/internal/modules/payment/domain"

type QrCodePaymentResponse struct {
	QrData         string `json:"qr_data"`
	InStoreOrderId string `json:"in_store_order_id"`
}

func (qcpr QrCodePaymentResponse) MpQrCodeResponseToPaymentEntity() *domain.Payment {
	return &domain.Payment{
		Id:     qcpr.InStoreOrderId,
		QrCode: qcpr.QrData,
	}
}
