package input

type CreateOrderUseCasePort interface {
	CreateOrder(createOrderCommand any) (createOrderResult any, err error)
}
