package model

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type Ingredient struct {
	ID             uuid.UUID
	Number         int `gorm:"autoIncrement:true;unique"`
	Name           string
	Amount         int
	Type           string
	IngredientType IngredientType `gorm:"foreignKey:Type;references:Name"`
}

func (i Ingredient) ToDomain() *domain.Ingredient {
	return &domain.Ingredient{
		ID:     i.ID,
		Number: i.Number,
		Name:   i.Name,
		Amount: i.Amount,
		Type:   i.IngredientType.Name,
	}
}
