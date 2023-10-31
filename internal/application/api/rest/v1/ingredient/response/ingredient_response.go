package response

type ProductCreatedResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}
