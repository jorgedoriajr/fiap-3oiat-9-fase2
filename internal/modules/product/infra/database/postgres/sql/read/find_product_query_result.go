package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
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

type FindProductOrderQueryResult struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Number      int       `db:"number"`
	Quantity    int       `db:"quantity"`
	Amount      int       `db:"total_amount"`
	Description string    `db:"description"`
	Category    string    `db:"category"`
	Menu        bool      `db:"menu"`
	ImgPath     string    `db:"img_path"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
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
