package response

type ProductCategoryResponse struct {
	Name         string `json:"name"`
	AcceptCustom bool   `json:"acceptCustom"`
}
