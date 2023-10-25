package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/product/domain/entity"
	"time"
)

type FindProductQueryResult struct {
	ID          uuid.UUID   `db:"id"`
	Name        string      `db:"name"`
	Amount      int         `db:"amount"`
	Description string      `db:"description"`
	Category    string      `db:"category"`
	Menu        bool        `db:"menu"`
	Ingredients []uuid.UUID `db:"ingredients"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}

func (fc FindProductQueryResult) ToEntity() *entity.ProductEntity {
	return &entity.ProductEntity{
		ID:          fc.ID,
		Name:        fc.Name,
		Amount:      fc.Amount,
		Description: fc.Description,
		Category:    fc.Category,
		Menu:        fc.Menu,
		Ingredients: fc.Ingredients,
		CreatedAt:   fc.CreatedAt,
		UpdatedAt:   fc.UpdatedAt,
	}
}
