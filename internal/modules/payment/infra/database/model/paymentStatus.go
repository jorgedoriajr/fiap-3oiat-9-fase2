package model

import (
	"hamburgueria/internal/modules/payment/domain"
	"hamburgueria/internal/modules/payment/domain/valueobjects"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaymentStatus struct {
	Id                   uuid.UUID
	PaymentIntegrationId uuid.UUID
	PaymentStatus        string
}

func (ps PaymentStatus) ToDomain() *domain.PaymentStatus {
	return &domain.PaymentStatus{
		Id:                   ps.Id,
		PaymentIntegrationId: ps.PaymentIntegrationId,
		PaymentStatus:        valueobjects.Status(ps.PaymentStatus),
	}
}

func (ps PaymentStatus) BeforeCreate(tx *gorm.DB) (err error) {
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
