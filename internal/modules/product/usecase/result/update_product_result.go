package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type UpdateProductResult struct {
	Id          uuid.UUID
	Number      int
	Name        *string
	Description *string
	Category    *string
	Amount      *int
	Menu        *bool
	ImgPath     *string
	UpdatedAt   time.Time
}

func UpdateProductResultFromEntity(entity entity.ProductEntity) UpdateProductResult {
	return UpdateProductResult{
		Id:          entity.ID,
		Number:      entity.Number,
		Name:        &entity.Name,
		Description: &entity.Description,
		Category:    &entity.Category,
		Menu:        &entity.Menu,
		ImgPath:     &entity.ImgPath,
		UpdatedAt:   time.Now(),
	}
}
