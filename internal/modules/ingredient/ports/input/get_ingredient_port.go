package input

import (
	"context"
	"hamburgueria/internal/modules/ingredient/domain/entity"
)

type GetIngredientPort interface {
	GetIngredientById(ctx context.Context, id string) (*entity.Ingredient, error)
	GetIngredientByName(ctx context.Context, name string) (*entity.Ingredient, error)
	GetIngredientAll(ctx context.Context) ([]*entity.Ingredient, error)
}
