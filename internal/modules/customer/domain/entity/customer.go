package entity

import "time"

type Customer struct {
	Document       string    `json:"document"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	OptInPromotion bool      `json:"optInPromotion"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
