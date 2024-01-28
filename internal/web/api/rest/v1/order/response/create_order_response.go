package response

type OrderResponse struct {
	Amount      int    `json:"amount"`
	PaymentData []byte `json:"paymentData"`
}
