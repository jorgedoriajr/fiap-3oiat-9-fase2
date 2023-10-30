package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
)

type CreateIngredientResult struct {
	ID     uuid.UUID
	Name   string
	Amount int
	Type   valueobject.IngredientType
}

func ToCreateIngredientResultFrom(entity entity.IngredientEntity) CreateIngredientResult {
	return CreateIngredientResult{
		ID:     entity.ID,
		Name:   entity.Name,
		Amount: entity.Amount,
		Type:   entity.Type,
	}
}
