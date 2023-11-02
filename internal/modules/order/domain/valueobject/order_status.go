package valueobject

type OrderStatus string

const (
	Created        OrderStatus = "CREATED"
	PaymentCreated OrderStatus = "WAITING_PAYMENT"
	Started        OrderStatus = "IN_PREPARATION"
	Ready          OrderStatus = "READY"
	Completed      OrderStatus = "COMPLETED"
)
