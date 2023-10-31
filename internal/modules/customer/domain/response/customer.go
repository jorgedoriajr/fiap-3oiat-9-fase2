package response

type Customer struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
