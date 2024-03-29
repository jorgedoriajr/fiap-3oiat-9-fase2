// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	command "hamburgueria/internal/modules/payment/usecase/command"

	mock "github.com/stretchr/testify/mock"
)

// CreatePaymentStatusPort is an autogenerated mock type for the CreatePaymentStatusPort type
type CreatePaymentStatusPort struct {
	mock.Mock
}

type CreatePaymentStatusPort_Expecter struct {
	mock *mock.Mock
}

func (_m *CreatePaymentStatusPort) EXPECT() *CreatePaymentStatusPort_Expecter {
	return &CreatePaymentStatusPort_Expecter{mock: &_m.Mock}
}

// AddPaymentStatus provides a mock function with given fields: ctx, createOrderCommand
func (_m *CreatePaymentStatusPort) AddPaymentStatus(ctx context.Context, createOrderCommand command.CreatePaymentStatusCommand) error {
	ret := _m.Called(ctx, createOrderCommand)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.CreatePaymentStatusCommand) error); ok {
		r0 = rf(ctx, createOrderCommand)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePaymentStatusPort_AddPaymentStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPaymentStatus'
type CreatePaymentStatusPort_AddPaymentStatus_Call struct {
	*mock.Call
}

// AddPaymentStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - createOrderCommand command.CreatePaymentStatusCommand
func (_e *CreatePaymentStatusPort_Expecter) AddPaymentStatus(ctx interface{}, createOrderCommand interface{}) *CreatePaymentStatusPort_AddPaymentStatus_Call {
	return &CreatePaymentStatusPort_AddPaymentStatus_Call{Call: _e.mock.On("AddPaymentStatus", ctx, createOrderCommand)}
}

func (_c *CreatePaymentStatusPort_AddPaymentStatus_Call) Run(run func(ctx context.Context, createOrderCommand command.CreatePaymentStatusCommand)) *CreatePaymentStatusPort_AddPaymentStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(command.CreatePaymentStatusCommand))
	})
	return _c
}

func (_c *CreatePaymentStatusPort_AddPaymentStatus_Call) Return(_a0 error) *CreatePaymentStatusPort_AddPaymentStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CreatePaymentStatusPort_AddPaymentStatus_Call) RunAndReturn(run func(context.Context, command.CreatePaymentStatusCommand) error) *CreatePaymentStatusPort_AddPaymentStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewCreatePaymentStatusPort creates a new instance of CreatePaymentStatusPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreatePaymentStatusPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreatePaymentStatusPort {
	mock := &CreatePaymentStatusPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
