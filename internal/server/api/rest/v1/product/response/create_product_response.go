package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
)

type ProductCreatedResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Menu        bool      `json:"menu"`
	ImgPath     string    `json:"imgPath"`
}

func ProductCreatedResponseFromResult(result result.ProductResult) ProductCreatedResponse {
	return ProductCreatedResponse{
		Id:          result.Id,
		Name:        result.Name,
		Amount:      result.Amount,
		Description: result.Description,
		Category:    result.Category,
		Menu:        result.Menu,
		ImgPath:     result.ImgPath,
	}
}
