package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/product/usecase/result"
)

type FindIngredientQueryResult struct {
	ID     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	Amount int       `db:"amount"`
	Type   string    `db:"type"`
}

type FindProductIngredientQueryResult struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Amount   int       `db:"total_amount"`
	Type     string    `db:"type"`
	Quantity int       `db:"quantity,omitempty"`
}

func (fc FindIngredientQueryResult) ToEntity() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     fc.ID,
		Name:   fc.Name,
		Amount: fc.Amount,
		Type:   fc.Type,
	}
}

func (fc FindProductIngredientQueryResult) ToResult() result.FindProductsIngredientsResult {
	return result.FindProductsIngredientsResult{
		ID:       fc.ID,
		Name:     fc.Name,
		Amount:   fc.Amount,
		Type:     fc.Type,
		Quantity: fc.Quantity,
	}
}
