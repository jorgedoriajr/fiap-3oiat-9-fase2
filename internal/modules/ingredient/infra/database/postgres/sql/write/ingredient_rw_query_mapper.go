package write

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type InsertIngredientRWQueryMapper struct {
	ID     uuid.UUID `position:"0"`
	Name   string    `position:"1"`
	Amount int       `position:"2"`
	Type   string    `position:"3"`
}

func ToInsertIngredientQueryMapper(ingredient entity.IngredientEntity) InsertIngredientRWQueryMapper {
	return InsertIngredientRWQueryMapper{
		ID:     ingredient.ID,
		Name:   ingredient.Name,
		Amount: ingredient.Amount,
		Type:   string(ingredient.Type),
	}
}
