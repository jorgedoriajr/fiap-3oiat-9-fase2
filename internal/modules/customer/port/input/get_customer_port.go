package input

type GetCustomerPort interface {
	GetCustomer(document string) (customerResult any, error error)
}
