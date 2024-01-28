// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/product/usecase/result"
)

// GetProductCategoryUseCasePort is an autogenerated mock type for the GetProductCategoryUseCasePort type
type GetProductCategoryUseCasePort struct {
	mock.Mock
}

type GetProductCategoryUseCasePort_Expecter struct {
	mock *mock.Mock
}

func (_m *GetProductCategoryUseCasePort) EXPECT() *GetProductCategoryUseCasePort_Expecter {
	return &GetProductCategoryUseCasePort_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields: ctx
func (_m *GetProductCategoryUseCasePort) FindAll(ctx context.Context) ([]result.FindProductCategoryResult, error) {
	ret := _m.Called(ctx)

	var r0 []result.FindProductCategoryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]result.FindProductCategoryResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []result.FindProductCategoryResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.FindProductCategoryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductCategoryUseCasePort_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type GetProductCategoryUseCasePort_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *GetProductCategoryUseCasePort_Expecter) FindAll(ctx interface{}) *GetProductCategoryUseCasePort_FindAll_Call {
	return &GetProductCategoryUseCasePort_FindAll_Call{Call: _e.mock.On("FindAll", ctx)}
}

func (_c *GetProductCategoryUseCasePort_FindAll_Call) Run(run func(ctx context.Context)) *GetProductCategoryUseCasePort_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GetProductCategoryUseCasePort_FindAll_Call) Return(_a0 []result.FindProductCategoryResult, _a1 error) *GetProductCategoryUseCasePort_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetProductCategoryUseCasePort_FindAll_Call) RunAndReturn(run func(context.Context) ([]result.FindProductCategoryResult, error)) *GetProductCategoryUseCasePort_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetProductCategoryUseCasePort creates a new instance of GetProductCategoryUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetProductCategoryUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetProductCategoryUseCasePort {
	mock := &GetProductCategoryUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
