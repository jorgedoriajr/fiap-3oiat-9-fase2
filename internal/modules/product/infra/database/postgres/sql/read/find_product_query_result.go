package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/usecase/result"
	"time"
)

type FindProductQueryResult struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Number      int       `db:"number"`
	Amount      int       `db:"amount"`
	Description string    `db:"description"`
	Category    string    `db:"category"`
	Menu        bool      `db:"menu"`
	ImgPath     string    `db:"img_path"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type FindProductWithIngredientsQueryResult struct {
	ID          uuid.UUID                        `db:"id"`
	Name        string                           `db:"name"`
	Number      int                              `db:"number"`
	Amount      int                              `db:"amount"`
	Description string                           `db:"description"`
	Category    string                           `db:"category"`
	Menu        bool                             `db:"menu"`
	ImgPath     string                           `db:"img_path"`
	Ingredients []read.FindIngredientQueryResult `db:"ingredients"`
	CreatedAt   time.Time                        `db:"created_at"`
	UpdatedAt   time.Time                        `db:"updated_at"`
}

func (fc FindProductQueryResult) ToEntity() *entity.ProductEntity {
	return &entity.ProductEntity{
		ID:          fc.ID,
		Name:        fc.Name,
		Number:      fc.Number,
		Amount:      fc.Amount,
		Description: fc.Description,
		Category:    fc.Category,
		Menu:        fc.Menu,
		ImgPath:     fc.ImgPath,
		CreatedAt:   fc.CreatedAt,
		UpdatedAt:   fc.UpdatedAt,
	}
}

func (fc FindProductWithIngredientsQueryResult) ToResult() *result.FindProductWithIngredientsResult {
	return &result.FindProductWithIngredientsResult{
		ID:          fc.ID,
		Name:        fc.Name,
		Number:      fc.Number,
		Amount:      fc.Amount,
		Description: fc.Description,
		Category:    fc.Category,
		Menu:        fc.Menu,
		ImgPath:     fc.ImgPath,
		CreatedAt:   fc.CreatedAt,
		UpdatedAt:   fc.UpdatedAt,
	}
}

func (fc FindProductQueryResult) ToCommandResult() *entity.ProductEntity {
	return &entity.ProductEntity{
		ID:          fc.ID,
		Name:        fc.Name,
		Number:      fc.Number,
		Amount:      fc.Amount,
		Description: fc.Description,
		Category:    fc.Category,
		Menu:        fc.Menu,
		ImgPath:     fc.ImgPath,
		CreatedAt:   fc.CreatedAt,
		UpdatedAt:   fc.UpdatedAt,
	}
}
