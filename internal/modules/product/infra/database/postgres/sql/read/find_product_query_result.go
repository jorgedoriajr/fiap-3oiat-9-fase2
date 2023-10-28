package read

import (
	"hamburgueria/internal/modules/product/domain/entity"
	"hamburgueria/internal/modules/product/domain/valueobject"
	"time"
)

type FindProductQueryResult struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Amount      int       `db:"amount"`
	Description string    `db:"description"`
	Category    string    `db:"category"`
	Menu        bool      `db:"menu"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (fc FindProductQueryResult) ToEntity() *entity.ProductEntity {
	return &entity.ProductEntity{
		ID:          fc.ID,
		Name:        fc.Name,
		Amount:      fc.Amount,
		Description: fc.Description,
		Category:    valueobject.ProductCategory(fc.Category),
		Menu:        fc.Menu,
		CreatedAt:   fc.CreatedAt,
		UpdatedAt:   fc.UpdatedAt,
	}
}
