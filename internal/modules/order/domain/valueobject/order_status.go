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
