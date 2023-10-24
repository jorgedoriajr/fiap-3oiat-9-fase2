// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CreateOrderUseCasePort is an autogenerated mock type for the CreateOrderUseCasePort type
type CreateOrderUseCasePort struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: createOrderCommand
func (_m *CreateOrderUseCasePort) CreateOrder(createOrderCommand interface{}) (interface{}, error) {
	ret := _m.Called(createOrderCommand)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(createOrderCommand)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(createOrderCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(createOrderCommand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateOrderUseCasePort creates a new instance of CreateOrderUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateOrderUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateOrderUseCasePort {
	mock := &CreateOrderUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}