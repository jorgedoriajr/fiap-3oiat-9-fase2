// Code generated by mockery v2.39.2. DO NOT EDIT.

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

type UpdateProductUseCasePort_Expecter struct {
	mock *mock.Mock
}

func (_m *UpdateProductUseCasePort) EXPECT() *UpdateProductUseCasePort_Expecter {
	return &UpdateProductUseCasePort_Expecter{mock: &_m.Mock}
}

// UpdateProduct provides a mock function with given fields: ctx, _a1
func (_m *UpdateProductUseCasePort) UpdateProduct(ctx context.Context, _a1 command.UpdateProductCommand) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, command.UpdateProductCommand) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductUseCasePort_UpdateProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProduct'
type UpdateProductUseCasePort_UpdateProduct_Call struct {
	*mock.Call
}

// UpdateProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 command.UpdateProductCommand
func (_e *UpdateProductUseCasePort_Expecter) UpdateProduct(ctx interface{}, _a1 interface{}) *UpdateProductUseCasePort_UpdateProduct_Call {
	return &UpdateProductUseCasePort_UpdateProduct_Call{Call: _e.mock.On("UpdateProduct", ctx, _a1)}
}

func (_c *UpdateProductUseCasePort_UpdateProduct_Call) Run(run func(ctx context.Context, _a1 command.UpdateProductCommand)) *UpdateProductUseCasePort_UpdateProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(command.UpdateProductCommand))
	})
	return _c
}

func (_c *UpdateProductUseCasePort_UpdateProduct_Call) Return(_a0 error) *UpdateProductUseCasePort_UpdateProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UpdateProductUseCasePort_UpdateProduct_Call) RunAndReturn(run func(context.Context, command.UpdateProductCommand) error) *UpdateProductUseCasePort_UpdateProduct_Call {
	_c.Call.Return(run)
	return _c
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
