package read

import (
	"hamburgueria/internal/modules/customer/domain/entity"
	"time"
)

type FindCustomerQueryResult struct {
	Document       string    `db:"cpf"`
	Name           string    `db:"name"`
	Phone          string    `db:"phone"`
	Email          string    `db:"email"`
	OptInPromotion bool      `db:"opt_in_promotion"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (fc FindCustomerQueryResult) ToEntity() *entity.Customer {
	return &entity.Customer{
		Document:       fc.Document,
		Name:           fc.Name,
		Phone:          fc.Phone,
		Email:          fc.Email,
		OptInPromotion: fc.OptInPromotion,
		CreatedAt:      fc.CreatedAt,
		UpdatedAt:      fc.UpdatedAt,
	}
}
