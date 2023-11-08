package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type ProductUpdatedResponse struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Number      int       `json:"number"`
	Amount      *int      `json:"amount,omitempty"`
	Description *string   `json:"description,omitempty"`
	Category    *string   `json:"category,omitempty"`
	Menu        *bool     `json:"menu,omitempty"`
	ImgPath     *string   `json:"img_path,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ProductUpdatedResponseFromResult(result result.UpdateProductResult) ProductUpdatedResponse {
	return ProductUpdatedResponse{
		Id:          result.Id,
		Number:      result.Number,
		Name:        result.Name,
		Amount:      result.Amount,
		Description: result.Description,
		Category:    result.Category,
		Menu:        result.Menu,
		ImgPath:     result.ImgPath,
		UpdatedAt:   result.UpdatedAt,
	}
}
