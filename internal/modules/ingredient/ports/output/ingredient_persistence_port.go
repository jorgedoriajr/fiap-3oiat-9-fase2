package output

import (
	"context"
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
)

type IngredientPersistencePort interface {
	Create(ctx context.Context, ingredient entity.IngredientEntity) error
	GetAll(ctx context.Context) ([]entity.IngredientEntity, error)
	GetByID(ctx context.Context, ingredientID uuid.UUID) (*entity.IngredientEntity, error)
	GetByProductID(ctx context.Context, productID uuid.UUID) ([]read.FindProductIngredientQueryResult, error)
	GetByType(ctx context.Context, ingredientType string) ([]entity.IngredientEntity, error)
}
