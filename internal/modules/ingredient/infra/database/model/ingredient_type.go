package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/ingredient/domain"
)

type IngredientType struct {
	Name                    string                          `gorm:"primarykey"`
	ConfigByProductCategory []IngredientTypeProductCategory `gorm:"foreignKey:IngredientType"`
}

type IngredientTypeProductCategory struct {
	ID              uuid.UUID
	IngredientType  string
	Optional        bool
	MaxQtd          int
	ProductCategory string
}

func (i IngredientType) ToDomain() *domain.IngredientType {
	var configByProductCategory []domain.IngredientTypeProductCategory
	for _, config := range i.ConfigByProductCategory {
		configByProductCategory = append(configByProductCategory, *config.ToDomain())
	}
	return &domain.IngredientType{
		Name:                    i.Name,
		ConfigByProductCategory: configByProductCategory,
	}
}

func (i IngredientTypeProductCategory) ToDomain() *domain.IngredientTypeProductCategory {
	return &domain.IngredientTypeProductCategory{
		IngredientType:  i.IngredientType,
		Optional:        i.Optional,
		MaxQtd:          i.MaxQtd,
		ProductCategory: i.ProductCategory,
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

func (i IngredientType) BeforeCreate(tx *gorm.DB) (err error) {
	var cols []clause.Column
	var colsNames []string
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
		colsNames = append(colsNames, field.DBName)
	}
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoNothing: true,
	})
	return nil
}

func (i IngredientTypeProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	var cols []clause.Column
	var colsNames []string
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
		colsNames = append(colsNames, field.DBName)
	}
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoNothing: true,
	})
	return nil
}
