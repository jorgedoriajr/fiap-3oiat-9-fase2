package response

type IngredientResponse struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}
