// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	command "hamburgueria/internal/modules/product/usecase/command"

	mock "github.com/stretchr/testify/mock"
)

// UpdateProductUseCasePort is an autogenerated mock type for the UpdateProductUseCasePort type
type UpdateProductUseCasePort struct {
	mock.Mock
}

// UpdateProduct provides a mock function with given fields: ctx, _a1
func (_m *UpdateProductUseCasePort) UpdateProduct(ctx context.Context, _a1 command.UpdateProductCommand) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.UpdateProductCommand) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUpdateProductUseCasePort creates a new instance of UpdateProductUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUpdateProductUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *UpdateProductUseCasePort {
	mock := &UpdateProductUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}