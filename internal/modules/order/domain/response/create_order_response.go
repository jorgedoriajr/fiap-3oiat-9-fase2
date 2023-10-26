package response

type OrderResponse struct {
	Amount      int64 `json:"amount"`
	PaymentData string
}
