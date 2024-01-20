// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/order/usecase/result"
)

// ListOrderPort is an autogenerated mock type for the ListOrderPort type
type ListOrderPort struct {
	mock.Mock
}

type ListOrderPort_Expecter struct {
	mock *mock.Mock
}

func (_m *ListOrderPort) EXPECT() *ListOrderPort_Expecter {
	return &ListOrderPort_Expecter{mock: &_m.Mock}
}

// FindAllOrders provides a mock function with given fields: ctx
func (_m *ListOrderPort) FindAllOrders(ctx context.Context) ([]result.ListOrderResult, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAllOrders")
	}

	var r0 []result.ListOrderResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]result.ListOrderResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []result.ListOrderResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.ListOrderResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrderPort_FindAllOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllOrders'
type ListOrderPort_FindAllOrders_Call struct {
	*mock.Call
}

// FindAllOrders is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ListOrderPort_Expecter) FindAllOrders(ctx interface{}) *ListOrderPort_FindAllOrders_Call {
	return &ListOrderPort_FindAllOrders_Call{Call: _e.mock.On("FindAllOrders", ctx)}
}

func (_c *ListOrderPort_FindAllOrders_Call) Run(run func(ctx context.Context)) *ListOrderPort_FindAllOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ListOrderPort_FindAllOrders_Call) Return(_a0 []result.ListOrderResult, _a1 error) *ListOrderPort_FindAllOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListOrderPort_FindAllOrders_Call) RunAndReturn(run func(context.Context) ([]result.ListOrderResult, error)) *ListOrderPort_FindAllOrders_Call {
	_c.Call.Return(run)
	return _c
}

// FindByStatus provides a mock function with given fields: ctx, status
func (_m *ListOrderPort) FindByStatus(ctx context.Context, status string) ([]result.ListOrderResult, error) {
	ret := _m.Called(ctx, status)

	if len(ret) == 0 {
		panic("no return value specified for FindByStatus")
	}

	var r0 []result.ListOrderResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]result.ListOrderResult, error)); ok {
		return rf(ctx, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []result.ListOrderResult); ok {
		r0 = rf(ctx, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.ListOrderResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrderPort_FindByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByStatus'
type ListOrderPort_FindByStatus_Call struct {
	*mock.Call
}

// FindByStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - status string
func (_e *ListOrderPort_Expecter) FindByStatus(ctx interface{}, status interface{}) *ListOrderPort_FindByStatus_Call {
	return &ListOrderPort_FindByStatus_Call{Call: _e.mock.On("FindByStatus", ctx, status)}
}

func (_c *ListOrderPort_FindByStatus_Call) Run(run func(ctx context.Context, status string)) *ListOrderPort_FindByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ListOrderPort_FindByStatus_Call) Return(_a0 []result.ListOrderResult, _a1 error) *ListOrderPort_FindByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ListOrderPort_FindByStatus_Call) RunAndReturn(run func(context.Context, string) ([]result.ListOrderResult, error)) *ListOrderPort_FindByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewListOrderPort creates a new instance of ListOrderPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListOrderPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *ListOrderPort {
	mock := &ListOrderPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}