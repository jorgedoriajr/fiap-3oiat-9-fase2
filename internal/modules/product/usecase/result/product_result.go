package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain"
	"time"
)

type ProductResult struct {
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

func FromDomain(entity domain.Product) *ProductResult {
	return &ProductResult{
		Id:          entity.ID,
		Name:        entity.Name,
		Amount:      entity.Amount,
		Description: entity.Description,
		Category:    entity.Category.Name,
		Menu:        entity.Menu,
		ImgPath:     entity.ImgPath,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
