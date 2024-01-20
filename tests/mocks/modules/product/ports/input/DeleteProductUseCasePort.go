// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DeleteProductUseCasePort is an autogenerated mock type for the DeleteProductUseCasePort type
type DeleteProductUseCasePort struct {
	mock.Mock
}

type DeleteProductUseCasePort_Expecter struct {
	mock *mock.Mock
}

func (_m *DeleteProductUseCasePort) EXPECT() *DeleteProductUseCasePort_Expecter {
	return &DeleteProductUseCasePort_Expecter{mock: &_m.Mock}
}

// Inactive provides a mock function with given fields: ctx, number
func (_m *DeleteProductUseCasePort) Inactive(ctx context.Context, number int) error {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for Inactive")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProductUseCasePort_Inactive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Inactive'
type DeleteProductUseCasePort_Inactive_Call struct {
	*mock.Call
}

// Inactive is a helper method to define mock.On call
//   - ctx context.Context
//   - number int
func (_e *DeleteProductUseCasePort_Expecter) Inactive(ctx interface{}, number interface{}) *DeleteProductUseCasePort_Inactive_Call {
	return &DeleteProductUseCasePort_Inactive_Call{Call: _e.mock.On("Inactive", ctx, number)}
}

func (_c *DeleteProductUseCasePort_Inactive_Call) Run(run func(ctx context.Context, number int)) *DeleteProductUseCasePort_Inactive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *DeleteProductUseCasePort_Inactive_Call) Return(_a0 error) *DeleteProductUseCasePort_Inactive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteProductUseCasePort_Inactive_Call) RunAndReturn(run func(context.Context, int) error) *DeleteProductUseCasePort_Inactive_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeleteProductUseCasePort creates a new instance of DeleteProductUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeleteProductUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeleteProductUseCasePort {
	mock := &DeleteProductUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}