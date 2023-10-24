package write

import (
	"hamburgueria/internal/modules/customer/domain/entity"
	"time"
)

type InsertCustomerRWQueryMapper struct {
	Document       string    `position:"0"`
	Phone          string    `position:"1"`
	Name           string    `position:"2"`
	Email          string    `position:"3"`
	OptInPromotion bool      `position:"4"`
	CreatedAt      time.Time `position:"5"`
	UpdatedAt      time.Time `position:"6"`
}

func EntityToInsertCustomerQueryMapper(customer entity.Customer) InsertCustomerRWQueryMapper {
	return InsertCustomerRWQueryMapper{
		Document:       customer.Document,
		Phone:          customer.Phone,
		Name:           customer.Name,
		Email:          customer.Email,
		OptInPromotion: customer.OptInPromotion,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}
}
