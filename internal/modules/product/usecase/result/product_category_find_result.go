package result

import (
	"github.com/google/uuid"
)

type FindProductCategoryResult struct {
	ID   uuid.UUID
	Name string
}
