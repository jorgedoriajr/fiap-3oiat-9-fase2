package model

import (
	"hamburgueria/internal/modules/payment/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Payment struct {
	Id       uuid.UUID `gorm:"primarykey"`
	OrderId  uuid.UUID
	Data     string
	CreateAt time.Time
}

func (p Payment) ToDomain() *domain.Payment {
	return &domain.Payment{
		Id:        p.Id,
		OrderId:   p.OrderId,
		Data:      p.Data,
		CreatedAt: p.CreateAt,
	}
}

func (p Payment) BeforeCreate(tx *gorm.DB) (err error) {
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
