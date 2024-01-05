// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	command "hamburgueria/internal/modules/payment/usecase/command"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/payment/usecase/result"
)

// CreatePaymentPort is an autogenerated mock type for the CreatePaymentPort type
type CreatePaymentPort struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: ctx, _a1
func (_m *CreatePaymentPort) CreatePayment(ctx context.Context, _a1 command.CreatePaymentCommand) (*result.PaymentProcessed, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *result.PaymentProcessed
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, command.CreatePaymentCommand) (*result.PaymentProcessed, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, command.CreatePaymentCommand) *result.PaymentProcessed); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*result.PaymentProcessed)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, command.CreatePaymentCommand) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreatePaymentPort creates a new instance of CreatePaymentPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreatePaymentPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreatePaymentPort {
	mock := &CreatePaymentPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
