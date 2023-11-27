package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"time"
)

type OrderHistory struct {
	ID        uuid.UUID
	OrderId   uuid.UUID
	Status    string
	ChangeBy  string
	CreatedAt time.Time
}

func (o OrderHistory) ToDomain() domain.OrderHistory {
	return domain.OrderHistory{
		Id:        o.ID,
		OrderId:   o.OrderId,
		Status:    valueobject.OrderStatus(o.Status),
		ChangeBy:  o.ChangeBy,
		CreatedAt: o.CreatedAt,
	}
}

func OrderHistoryFromDomain(history domain.OrderHistory) OrderHistory {
	return OrderHistory{
		ID:        history.Id,
		OrderId:   history.OrderId,
		Status:    string(history.Status),
		ChangeBy:  history.ChangeBy,
		CreatedAt: history.CreatedAt,
	}
}

func (o OrderHistory) BeforeCreate(tx *gorm.DB) (err error) {
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
