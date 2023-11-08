package model

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type IngredientType struct {
	Name                    string `gorm:"primarykey"`
	ConfigByProductCategory []IngredientTypeProductCategory
}

type IngredientTypeProductCategory struct {
	ID              uuid.UUID
	Optional        string
	MaxQtd          string
	ProductCategory string
}

func (i IngredientType) ToDomain() *domain.IngredientType {
	var configByProductCategory []domain.IngredientTypeProductCategory
	for _, config := range i.ConfigByProductCategory {
		configByProductCategory = append(configByProductCategory, domain.IngredientTypeProductCategory{
			Optional:        config.Optional,
			MaxQtd:          config.MaxQtd,
			ProductCategory: config.ProductCategory,
		})
	}
	return &domain.IngredientType{
		Name:                    i.Name,
		ConfigByProductCategory: configByProductCategory,
	}
}
