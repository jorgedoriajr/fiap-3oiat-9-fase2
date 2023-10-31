package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type GetIngredientPort interface {
	GetIngredientById(ctx context.Context, id string) (*entity.IngredientEntity, error)
	GetIngredientByName(ctx context.Context, name string) (*entity.IngredientEntity, error)
	GetIngredientAll(ctx context.Context) ([]*entity.IngredientEntity, error)
}
