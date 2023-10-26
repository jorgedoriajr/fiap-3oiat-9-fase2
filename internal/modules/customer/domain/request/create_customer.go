package request

import "hamburgueria/internal/modules/customer/usecase/command"

type CreateCustomer struct {
	Document       string `json:"document"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	OptInPromotion bool   `json:"optInPromotion"`
}

func (c CreateCustomer) ToCommand() command.CreateCustomerCommand {
	return command.CreateCustomerCommand{
		Document:       c.Document,
		Name:           c.Name,
		Phone:          c.Phone,
		Email:          c.Email,
		OptInPromotion: c.OptInPromotion,
	}
}
