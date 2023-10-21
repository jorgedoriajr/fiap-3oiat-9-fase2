package request

type CreateCustomerCommand struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}
