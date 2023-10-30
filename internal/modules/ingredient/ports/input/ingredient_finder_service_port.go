package input

import (
	"context"
	"github.com/google/uuid"
	entity2 "hamburgueria/internal/modules/ingredient/domain/entity"
	"hamburgueria/internal/modules/ingredient/domain/valueobject"
	"hamburgueria/internal/modules/ingredient/infra/database/postgres/sql/read"
)

type IngredientFinderServicePort interface {
	FindIngredientByID(ctx context.Context, id uuid.UUID) (*entity2.IngredientEntity, error)

	FindIngredientsByProductId(ctx context.Context, productID uuid.UUID) ([]read.FindIngredientQueryResult, error)

	FindIngredientByType(ctx context.Context, ingredientType valueobject.IngredientType) ([]entity2.IngredientEntity, error)

	FindAllIngredients(ctx context.Context) ([]entity2.IngredientEntity, error)
}
