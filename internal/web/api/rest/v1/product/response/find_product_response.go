package response

type FindProductWithIngredients struct {
	Name        string                    `json:"name"`
	Number      int                       `json:"number"`
	Amount      int                       `json:"amount"`
	Description string                    `json:"description"`
	Category    string                    `json:"category"`
	ImgPath     string                    `json:"imgPath"`
	Ingredients []FindProductsIngredients `json:"ingredients"`
}

type FindProductsIngredients struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity"`
}
