package read

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type FindIngredientQueryResult struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Amount   int       `db:"amount"`
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
		Type:   fc.Type,
	}
}

func ToIngredientEntityList(results []FindIngredientQueryResult) []*entity.IngredientEntity {
	entities := make([]*entity.IngredientEntity, len(results))

	for i, result := range results {
		entities[i] = result.ToEntity()
	}

	return entities
}
