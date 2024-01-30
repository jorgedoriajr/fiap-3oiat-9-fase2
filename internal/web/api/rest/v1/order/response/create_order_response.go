package response

type OrderResponse struct {
	Number      int    `json:"number"`
	Amount      int    `json:"amount"`
	PaymentData string `json:"paymentData"`
}
