// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	command "hamburgueria/internal/modules/order/usecase/command"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/order/usecase/result"
)

// CreateOrderPort is an autogenerated mock type for the CreateOrderPort type
type CreateOrderPort struct {
	mock.Mock
}

// AddOrder provides a mock function with given fields: ctx, createOrderCommand
func (_m *CreateOrderPort) AddOrder(ctx context.Context, createOrderCommand command.CreateOrderCommand) (*result.CreateOrderResult, error) {
	ret := _m.Called(ctx, createOrderCommand)

	var r0 *result.CreateOrderResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, command.CreateOrderCommand) (*result.CreateOrderResult, error)); ok {
		return rf(ctx, createOrderCommand)
	}
	if rf, ok := ret.Get(0).(func(context.Context, command.CreateOrderCommand) *result.CreateOrderResult); ok {
		r0 = rf(ctx, createOrderCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*result.CreateOrderResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, command.CreateOrderCommand) error); ok {
		r1 = rf(ctx, createOrderCommand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateOrderPort creates a new instance of CreateOrderPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateOrderPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateOrderPort {
	mock := &CreateOrderPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
