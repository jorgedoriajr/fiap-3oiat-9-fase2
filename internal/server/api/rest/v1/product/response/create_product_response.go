package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
)

type ProductCreatedResponse struct {
	Id          uuid.UUID
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
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
