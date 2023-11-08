package entity

import "time"

type Customer struct {
	Document       string
	Name           string
	Phone          string
	Email          string
	OptInPromotion bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
