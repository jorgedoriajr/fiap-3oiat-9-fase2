package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/product/domain"
)

type ProductCategory struct {
	Name         string `gorm:"primarykey"`
	AcceptCustom bool
}

func (p ProductCategory) ToDomain() *domain.ProductCategory {
	return &domain.ProductCategory{
		Name:         p.Name,
		AcceptCustom: p.AcceptCustom,
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
