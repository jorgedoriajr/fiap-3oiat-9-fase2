package response

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/domain/valueobject"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type ProductCreatedResponse struct {
	ID          uuid.UUID                   `json:"id,omitempty"`
	Number      int                         `json:"number,omitempty"`
	Name        string                      `json:"name"`
	Amount      int                         `json:"amount"`
	Description string                      `json:"description"`
	Category    valueobject.ProductCategory `json:"category"`
	Menu        bool                        `json:"menu"`
	CreatedAt   time.Time                   `json:"created_at"`
	UpdatedAt   time.Time                   `json:"updated_at"`
}

func (pcr ProductCreatedResponse) From(entity entity.ProductEntity) ProductCreatedResponse {
	return ProductCreatedResponse{
		ID:          entity.ID,
		Number:      entity.Number,
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    entity.Category,
		Menu:        entity.Menu,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (pcr ProductCreatedResponse) FromResult(result result.CreateProductResult) ProductCreatedResponse {
	return ProductCreatedResponse{
		Name:        result.Name,
		Amount:      result.Amount,
		Description: result.Description,
		Category:    result.Category,
		Menu:        result.Menu,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}
}
