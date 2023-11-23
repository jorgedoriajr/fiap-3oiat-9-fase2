package model

import (
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
