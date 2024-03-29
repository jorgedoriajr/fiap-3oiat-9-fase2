// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/ingredient/usecase/result"
)

// FindIngredientUseCasePort is an autogenerated mock type for the FindIngredientUseCasePort type
type FindIngredientUseCasePort struct {
	mock.Mock
}

type FindIngredientUseCasePort_Expecter struct {
	mock *mock.Mock
}

func (_m *FindIngredientUseCasePort) EXPECT() *FindIngredientUseCasePort_Expecter {
	return &FindIngredientUseCasePort_Expecter{mock: &_m.Mock}
}

// FindAllIngredients provides a mock function with given fields: ctx
func (_m *FindIngredientUseCasePort) FindAllIngredients(ctx context.Context) ([]result.FindIngredientResult, error) {
	ret := _m.Called(ctx)

	var r0 []result.FindIngredientResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]result.FindIngredientResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []result.FindIngredientResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.FindIngredientResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIngredientUseCasePort_FindAllIngredients_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllIngredients'
type FindIngredientUseCasePort_FindAllIngredients_Call struct {
	*mock.Call
}

// FindAllIngredients is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FindIngredientUseCasePort_Expecter) FindAllIngredients(ctx interface{}) *FindIngredientUseCasePort_FindAllIngredients_Call {
	return &FindIngredientUseCasePort_FindAllIngredients_Call{Call: _e.mock.On("FindAllIngredients", ctx)}
}

func (_c *FindIngredientUseCasePort_FindAllIngredients_Call) Run(run func(ctx context.Context)) *FindIngredientUseCasePort_FindAllIngredients_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FindIngredientUseCasePort_FindAllIngredients_Call) Return(_a0 []result.FindIngredientResult, _a1 error) *FindIngredientUseCasePort_FindAllIngredients_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FindIngredientUseCasePort_FindAllIngredients_Call) RunAndReturn(run func(context.Context) ([]result.FindIngredientResult, error)) *FindIngredientUseCasePort_FindAllIngredients_Call {
	_c.Call.Return(run)
	return _c
}

// FindIngredientByNumber provides a mock function with given fields: ctx, number
func (_m *FindIngredientUseCasePort) FindIngredientByNumber(ctx context.Context, number int) (*result.FindIngredientResult, error) {
	ret := _m.Called(ctx, number)

	var r0 *result.FindIngredientResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*result.FindIngredientResult, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *result.FindIngredientResult); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*result.FindIngredientResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIngredientUseCasePort_FindIngredientByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindIngredientByNumber'
type FindIngredientUseCasePort_FindIngredientByNumber_Call struct {
	*mock.Call
}

// FindIngredientByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number int
func (_e *FindIngredientUseCasePort_Expecter) FindIngredientByNumber(ctx interface{}, number interface{}) *FindIngredientUseCasePort_FindIngredientByNumber_Call {
	return &FindIngredientUseCasePort_FindIngredientByNumber_Call{Call: _e.mock.On("FindIngredientByNumber", ctx, number)}
}

func (_c *FindIngredientUseCasePort_FindIngredientByNumber_Call) Run(run func(ctx context.Context, number int)) *FindIngredientUseCasePort_FindIngredientByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *FindIngredientUseCasePort_FindIngredientByNumber_Call) Return(_a0 *result.FindIngredientResult, _a1 error) *FindIngredientUseCasePort_FindIngredientByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FindIngredientUseCasePort_FindIngredientByNumber_Call) RunAndReturn(run func(context.Context, int) (*result.FindIngredientResult, error)) *FindIngredientUseCasePort_FindIngredientByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// FindIngredientByType provides a mock function with given fields: ctx, ingredientType
func (_m *FindIngredientUseCasePort) FindIngredientByType(ctx context.Context, ingredientType string) ([]result.FindIngredientResult, error) {
	ret := _m.Called(ctx, ingredientType)

	var r0 []result.FindIngredientResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]result.FindIngredientResult, error)); ok {
		return rf(ctx, ingredientType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []result.FindIngredientResult); ok {
		r0 = rf(ctx, ingredientType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.FindIngredientResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ingredientType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIngredientUseCasePort_FindIngredientByType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindIngredientByType'
type FindIngredientUseCasePort_FindIngredientByType_Call struct {
	*mock.Call
}

// FindIngredientByType is a helper method to define mock.On call
//   - ctx context.Context
//   - ingredientType string
func (_e *FindIngredientUseCasePort_Expecter) FindIngredientByType(ctx interface{}, ingredientType interface{}) *FindIngredientUseCasePort_FindIngredientByType_Call {
	return &FindIngredientUseCasePort_FindIngredientByType_Call{Call: _e.mock.On("FindIngredientByType", ctx, ingredientType)}
}

func (_c *FindIngredientUseCasePort_FindIngredientByType_Call) Run(run func(ctx context.Context, ingredientType string)) *FindIngredientUseCasePort_FindIngredientByType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *FindIngredientUseCasePort_FindIngredientByType_Call) Return(_a0 []result.FindIngredientResult, _a1 error) *FindIngredientUseCasePort_FindIngredientByType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FindIngredientUseCasePort_FindIngredientByType_Call) RunAndReturn(run func(context.Context, string) ([]result.FindIngredientResult, error)) *FindIngredientUseCasePort_FindIngredientByType_Call {
	_c.Call.Return(run)
	return _c
}

// NewFindIngredientUseCasePort creates a new instance of FindIngredientUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFindIngredientUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *FindIngredientUseCasePort {
	mock := &FindIngredientUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
