package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type FindIngredientQueryResult struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Amount   int       `db:"amount"`
	Type     string    `db:"type"`
	Quantity string    `db:"quantity,omitempty"`
}

func (fc FindIngredientQueryResult) ToEntity() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     fc.ID,
		Name:   fc.Name,
		Amount: fc.Amount,
		Type:   valueobject.IngredientType(fc.Type),
	}
}

func ToIngredientEntities(results []FindIngredientQueryResult) []*entity.IngredientEntity {
	var entities []*entity.IngredientEntity

	for _, r := range results {
		entities = append(entities, r.ToEntity())
	}
	return entities
}

func (fc FindIngredientQueryResult) ToCommandResult() *entity.IngredientEntity {
	return &entity.IngredientEntity{
		ID:     fc.ID,
		Name:   fc.Name,
		Amount: fc.Amount,
		Type:   valueobject.IngredientType(fc.Type),
	}
}
