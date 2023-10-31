// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	command "hamburgueria/internal/modules/customer/usecase/command"

	mock "github.com/stretchr/testify/mock"
)

// CreateCustomerPort is an autogenerated mock type for the CreateCustomerPort type
type CreateCustomerPort struct {
	mock.Mock
}

// AddCustomer provides a mock function with given fields: ctx, createCustomerCommand
func (_m *CreateCustomerPort) AddCustomer(ctx context.Context, createCustomerCommand command.CreateCustomerCommand) error {
	ret := _m.Called(ctx, createCustomerCommand)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.CreateCustomerCommand) error); ok {
		r0 = rf(ctx, createCustomerCommand)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCreateCustomerPort creates a new instance of CreateCustomerPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateCustomerPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateCustomerPort {
	mock := &CreateCustomerPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
