package write

import "time"

type InsertCustomerRWQueryMapper struct {
	Document       string    `position:"0"`
	Phone          string    `position:"1"`
	Name           string    `position:"2"`
	Email          string    `position:"3"`
	OptInPromotion bool      `position:"4"`
	CreatedAt      time.Time `position:"5"`
	UpdatedAt      time.Time `position:"6"`
}
