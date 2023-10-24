// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PaymentPayPort is an autogenerated mock type for the PaymentPayPort type
type PaymentPayPort struct {
	mock.Mock
}

// Pay provides a mock function with given fields: paymentPayCommand
func (_m *PaymentPayPort) Pay(paymentPayCommand interface{}) (interface{}, error) {
	ret := _m.Called(paymentPayCommand)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(paymentPayCommand)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(paymentPayCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(paymentPayCommand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPaymentPayPort creates a new instance of PaymentPayPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentPayPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentPayPort {
	mock := &PaymentPayPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}