package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/application/api/rest/v1/product/response"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type CreateProductResult struct {
	Id          uuid.UUID
	Name        string
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromEntity(entity entity.ProductEntity) CreateProductResult {
	return CreateProductResult{
		Id:          entity.ID,
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    entity.Category,
		Menu:        entity.Menu,
		ImgPath:     entity.ImgPath,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (c CreateProductResult) ToResponse() response.ProductCreatedResponse {
	return response.ProductCreatedResponse{
		Id:          c.Id,
		Name:        c.Name,
		Amount:      c.Amount,
		Description: c.Description,
		Category:    c.Category,
		Menu:        c.Menu,
		ImgPath:     c.ImgPath,
	}
}
