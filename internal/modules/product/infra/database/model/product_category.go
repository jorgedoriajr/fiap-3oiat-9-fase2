package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	ingredientDomain "hamburgueria/internal/modules/ingredient/domain"
	"hamburgueria/internal/modules/ingredient/infra/database/model"
	"hamburgueria/internal/modules/product/domain"
)

type ProductCategory struct {
	Name                    string `gorm:"primarykey"`
	AcceptCustom            bool
	ConfigByProductCategory []model.IngredientTypeProductCategory `gorm:"foreignKey:ProductCategory"`
}

func (p ProductCategory) ToDomain() *domain.ProductCategory {
	var configs []ingredientDomain.IngredientTypeProductCategory
	for _, config := range p.ConfigByProductCategory {
		configs = append(configs, *config.ToDomain())
	}

	return &domain.ProductCategory{
		Name:                    p.Name,
		AcceptCustom:            p.AcceptCustom,
		ConfigByProductCategory: configs,
	}
}

func (p ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
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
