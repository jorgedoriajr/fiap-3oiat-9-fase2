package input

type CreatePaymentPort interface {
	CreatePayment(createPaymentCommand any) (createPaymentResult any, error error)
}
