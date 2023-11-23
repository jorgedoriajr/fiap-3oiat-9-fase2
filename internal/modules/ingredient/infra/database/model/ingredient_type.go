package model

import (
	"github.com/google/uuid"
	"hamburgueria/internal/modules/ingredient/domain"
)

type IngredientType struct {
	Name                    string                          `gorm:"primarykey"`
	ConfigByProductCategory []IngredientTypeProductCategory `gorm:"foreignKey:IngredientType"`
}

type IngredientTypeProductCategory struct {
	ID              uuid.UUID
	IngredientType  string
	Optional        string
	MaxQtd          string
	ProductCategory string
}

func (i IngredientType) ToDomain() *domain.IngredientType {
	var configByProductCategory []domain.IngredientTypeProductCategory
	for _, config := range i.ConfigByProductCategory {
		configByProductCategory = append(configByProductCategory, domain.IngredientTypeProductCategory{
			IngredientType:  config.IngredientType,
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

func FromIngredientTypeDomain(ingredientType domain.IngredientType) IngredientType {
	var configByProductCategory []IngredientTypeProductCategory

	for _, config := range ingredientType.ConfigByProductCategory {
		configByProductCategory = append(configByProductCategory, IngredientTypeProductCategory{
			ID:              config.Id,
			IngredientType:  config.IngredientType,
			Optional:        config.Optional,
			MaxQtd:          config.MaxQtd,
			ProductCategory: config.ProductCategory,
		})
	}

	return IngredientType{
		Name:                    ingredientType.Name,
		ConfigByProductCategory: configByProductCategory,
	}
}
