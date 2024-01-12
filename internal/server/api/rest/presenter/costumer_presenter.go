package presenter

import "hamburgueria/internal/modules/customer/usecase/result"

type GetCustomerResponse struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func ToGetCustomerResponse(result result.CustomerCreated) GetCustomerResponse {
	return GetCustomerResponse{
		Document: result.Document,
		Name:     result.Name,
		Phone:    result.Phone,
		Email:    result.Email,
	}
}
