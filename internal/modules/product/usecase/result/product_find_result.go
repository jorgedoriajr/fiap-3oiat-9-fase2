package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
	"sync/atomic"
	"time"
)

type FindProductResult struct {
	ID          uuid.UUID
	Name        string
	Number      int
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FindProductWithIngredientsResult struct {
	ID          uuid.UUID
	Name        string
	Number      int
	Amount      int
	Description string
	Category    string
	Menu        bool
	ImgPath     string
	Ingredients []read.FindIngredientQueryResult
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (f *FindProductWithIngredientsResult) CalculateIngredientsAmount() {

	var totalAmount int64

	for _, ig := range f.Ingredients {
		atomic.AddInt64(&totalAmount, int64(ig.Amount*ig.Quantity))
	}
	f.Amount = int(totalAmount)
}
