// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CreatePaymentPort is an autogenerated mock type for the CreatePaymentPort type
type CreatePaymentPort struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: createPaymentCommand
func (_m *CreatePaymentPort) CreatePayment(createPaymentCommand interface{}) (interface{}, error) {
	ret := _m.Called(createPaymentCommand)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(createPaymentCommand)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(createPaymentCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(createPaymentCommand)
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