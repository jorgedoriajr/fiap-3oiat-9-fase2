package valueobject

type OrderStatus string

const (
	Created        OrderStatus = "CREATED"
	PaymentCreated OrderStatus = "WAITING_PAYMENT"
	PaymentRefused OrderStatus = "REFUSED"
	Started        OrderStatus = "IN_PREPARATION"
	Ready          OrderStatus = "READY"
	Completed      OrderStatus = "COMPLETED"
)

func (os OrderStatus) Order() int {
	switch os {
	case Ready:
		return 1
	case Started:
		return 2
	case PaymentCreated:
		return 3
	case Created:
		return 4
	default:
		return 9999999
	}
}
