package output

type OrderPersistencePort interface {
	Create(createOrderPersistenceCommand any) (createOrderPersistenceResult any, err error)
}
