package result

import (
	"github.com/google/uuid"
)

type CreateIngredientResult struct {
	Id     uuid.UUID
	Name   string
	Amount string
	Type   string
}
