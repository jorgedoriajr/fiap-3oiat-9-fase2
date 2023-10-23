package entity

import "time"

type Customer struct {
	Document       string    `json:"document" db:"cpf"`
	Name           string    `json:"name" db:"name"`
	Phone          string    `json:"phone" db:"phone"`
	Email          string    `json:"email" db:"email"`
	OptInPromotion bool      `json:"optInPromotion" db:"opt_in_promotion"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}
