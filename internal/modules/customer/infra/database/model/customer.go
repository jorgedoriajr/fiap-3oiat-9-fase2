package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/customer/domain"
	"time"
)

type Customer struct {
	Cpf            string `gorm:"primarykey"`
	Name           string
	Phone          string
	Email          string
	OptInPromotion bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c Customer) ToDomain() *domain.Customer {
	return &domain.Customer{
		Document:       c.Cpf,
		Name:           c.Name,
		Phone:          c.Phone,
		Email:          c.Email,
		OptInPromotion: c.OptInPromotion,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}

func (c Customer) BeforeCreate(tx *gorm.DB) (err error) {
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
