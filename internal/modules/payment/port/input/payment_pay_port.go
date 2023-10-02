package input

type PaymentPayPort interface {
	Pay(paymentPayCommand any) (paymentPayResult any, err error)
}
