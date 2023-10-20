package input

type CreateCustomerPort interface {
	CreateCustomer(createCustomerCommand any) (createCustomerResult any, error error)
}
