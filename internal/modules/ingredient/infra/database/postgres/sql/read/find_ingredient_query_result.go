package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type FindIngredientQueryResult struct {
	ID     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	Amount int       `db:"amount"`
	Type   string    `db:"type"`
}

func (fc FindIngredientQueryResult) ToEntity() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     fc.ID,
		Name:   fc.Name,
		Amount: fc.Amount,
		Type:   valueobject.IngredientType(fc.Type),
	}
}

func (fc FindIngredientQueryResult) ToCommandResult() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     fc.ID,
		Name:   fc.Name,
		Amount: fc.Amount,
		Type:   valueobject.IngredientType(fc.Type),
	}
}
