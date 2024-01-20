package response

type QrCodePaymentResponse struct {
	QrData         string `json:"qr_data"`
	InStoreOrderId string `json:"in_store_order_id"`
}
