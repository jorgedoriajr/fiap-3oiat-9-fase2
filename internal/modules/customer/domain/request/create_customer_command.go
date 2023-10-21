package request

type CreateCustomerCommand struct {
	Cpf   string `json:"cpf"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
