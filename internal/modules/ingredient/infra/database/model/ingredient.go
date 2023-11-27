package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/ingredient/domain"
)

type Ingredient struct {
	ID             uuid.UUID
	Number         int    `gorm:"autoIncrement:true;unique"`
	Name           string `gorm:"unique"`
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
		Type:   *i.IngredientType.ToDomain(),
	}
}

func (i Ingredient) BeforeCreate(tx *gorm.DB) (err error) {
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
