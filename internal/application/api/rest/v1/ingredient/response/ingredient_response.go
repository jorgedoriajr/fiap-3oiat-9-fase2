package response

type IngredientCreatedResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}
