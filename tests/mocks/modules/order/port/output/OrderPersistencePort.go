// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "hamburgueria/internal/modules/order/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// OrderPersistencePort is an autogenerated mock type for the OrderPersistencePort type
type OrderPersistencePort struct {
	mock.Mock
}

type OrderPersistencePort_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderPersistencePort) EXPECT() *OrderPersistencePort_Expecter {
	return &OrderPersistencePort_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, order
func (_m *OrderPersistencePort) Create(ctx context.Context, order domain.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderPersistencePort_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type OrderPersistencePort_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - order domain.Order
func (_e *OrderPersistencePort_Expecter) Create(ctx interface{}, order interface{}) *OrderPersistencePort_Create_Call {
	return &OrderPersistencePort_Create_Call{Call: _e.mock.On("Create", ctx, order)}
}

func (_c *OrderPersistencePort_Create_Call) Run(run func(ctx context.Context, order domain.Order)) *OrderPersistencePort_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Order))
	})
	return _c
}

func (_c *OrderPersistencePort_Create_Call) Return(_a0 error) *OrderPersistencePort_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderPersistencePort_Create_Call) RunAndReturn(run func(context.Context, domain.Order) error) *OrderPersistencePort_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: ctx
func (_m *OrderPersistencePort) FindAll(ctx context.Context) ([]domain.Order, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Order, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Order); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderPersistencePort_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type OrderPersistencePort_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OrderPersistencePort_Expecter) FindAll(ctx interface{}) *OrderPersistencePort_FindAll_Call {
	return &OrderPersistencePort_FindAll_Call{Call: _e.mock.On("FindAll", ctx)}
}

func (_c *OrderPersistencePort_FindAll_Call) Run(run func(ctx context.Context)) *OrderPersistencePort_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OrderPersistencePort_FindAll_Call) Return(_a0 []domain.Order, _a1 error) *OrderPersistencePort_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderPersistencePort_FindAll_Call) RunAndReturn(run func(context.Context) ([]domain.Order, error)) *OrderPersistencePort_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: ctx, orderId
func (_m *OrderPersistencePort) FindById(ctx context.Context, orderId uuid.UUID) (*domain.Order, error) {
	ret := _m.Called(ctx, orderId)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Order, error)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Order); ok {
		r0 = rf(ctx, orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderPersistencePort_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type OrderPersistencePort_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - ctx context.Context
//   - orderId uuid.UUID
func (_e *OrderPersistencePort_Expecter) FindById(ctx interface{}, orderId interface{}) *OrderPersistencePort_FindById_Call {
	return &OrderPersistencePort_FindById_Call{Call: _e.mock.On("FindById", ctx, orderId)}
}

func (_c *OrderPersistencePort_FindById_Call) Run(run func(ctx context.Context, orderId uuid.UUID)) *OrderPersistencePort_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *OrderPersistencePort_FindById_Call) Return(_a0 *domain.Order, _a1 error) *OrderPersistencePort_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderPersistencePort_FindById_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*domain.Order, error)) *OrderPersistencePort_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// FindByStatus provides a mock function with given fields: ctx, status
func (_m *OrderPersistencePort) FindByStatus(ctx context.Context, status string) ([]domain.Order, error) {
	ret := _m.Called(ctx, status)

	if len(ret) == 0 {
		panic("no return value specified for FindByStatus")
	}

	var r0 []domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Order, error)); ok {
		return rf(ctx, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Order); ok {
		r0 = rf(ctx, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderPersistencePort_FindByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByStatus'
type OrderPersistencePort_FindByStatus_Call struct {
	*mock.Call
}

// FindByStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - status string
func (_e *OrderPersistencePort_Expecter) FindByStatus(ctx interface{}, status interface{}) *OrderPersistencePort_FindByStatus_Call {
	return &OrderPersistencePort_FindByStatus_Call{Call: _e.mock.On("FindByStatus", ctx, status)}
}

func (_c *OrderPersistencePort_FindByStatus_Call) Run(run func(ctx context.Context, status string)) *OrderPersistencePort_FindByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderPersistencePort_FindByStatus_Call) Return(_a0 []domain.Order, _a1 error) *OrderPersistencePort_FindByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderPersistencePort_FindByStatus_Call) RunAndReturn(run func(context.Context, string) ([]domain.Order, error)) *OrderPersistencePort_FindByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, order
func (_m *OrderPersistencePort) Update(ctx context.Context, order domain.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderPersistencePort_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type OrderPersistencePort_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - order domain.Order
func (_e *OrderPersistencePort_Expecter) Update(ctx interface{}, order interface{}) *OrderPersistencePort_Update_Call {
	return &OrderPersistencePort_Update_Call{Call: _e.mock.On("Update", ctx, order)}
}

func (_c *OrderPersistencePort_Update_Call) Run(run func(ctx context.Context, order domain.Order)) *OrderPersistencePort_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Order))
	})
	return _c
}

func (_c *OrderPersistencePort_Update_Call) Return(_a0 error) *OrderPersistencePort_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderPersistencePort_Update_Call) RunAndReturn(run func(context.Context, domain.Order) error) *OrderPersistencePort_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderPersistencePort creates a new instance of OrderPersistencePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderPersistencePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderPersistencePort {
	mock := &OrderPersistencePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
