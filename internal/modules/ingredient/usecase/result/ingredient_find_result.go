package result

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
	"time"
)

type FindProductResult struct {
	ID        uuid.UUID
	Name      string
	Amount    int
	Type      valueobject.IngredientType
	CreatedAt time.Time
	UpdatedAt time.Time
}
