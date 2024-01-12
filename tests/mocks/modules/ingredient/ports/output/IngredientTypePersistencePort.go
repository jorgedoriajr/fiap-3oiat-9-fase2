// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "hamburgueria/internal/modules/ingredient/domain"

	mock "github.com/stretchr/testify/mock"
)

// IngredientTypePersistencePort is an autogenerated mock type for the IngredientTypePersistencePort type
type IngredientTypePersistencePort struct {
	mock.Mock
}

type IngredientTypePersistencePort_Expecter struct {
	mock *mock.Mock
}

func (_m *IngredientTypePersistencePort) EXPECT() *IngredientTypePersistencePort_Expecter {
	return &IngredientTypePersistencePort_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields: ctx
func (_m *IngredientTypePersistencePort) GetAll(ctx context.Context) ([]domain.IngredientType, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.IngredientType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.IngredientType, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.IngredientType); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IngredientType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientTypePersistencePort_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type IngredientTypePersistencePort_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *IngredientTypePersistencePort_Expecter) GetAll(ctx interface{}) *IngredientTypePersistencePort_GetAll_Call {
	return &IngredientTypePersistencePort_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *IngredientTypePersistencePort_GetAll_Call) Run(run func(ctx context.Context)) *IngredientTypePersistencePort_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *IngredientTypePersistencePort_GetAll_Call) Return(_a0 []domain.IngredientType, _a1 error) *IngredientTypePersistencePort_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientTypePersistencePort_GetAll_Call) RunAndReturn(run func(context.Context) ([]domain.IngredientType, error)) *IngredientTypePersistencePort_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *IngredientTypePersistencePort) GetByName(ctx context.Context, name string) (*domain.IngredientType, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *domain.IngredientType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.IngredientType, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.IngredientType); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.IngredientType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientTypePersistencePort_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type IngredientTypePersistencePort_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *IngredientTypePersistencePort_Expecter) GetByName(ctx interface{}, name interface{}) *IngredientTypePersistencePort_GetByName_Call {
	return &IngredientTypePersistencePort_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *IngredientTypePersistencePort_GetByName_Call) Run(run func(ctx context.Context, name string)) *IngredientTypePersistencePort_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IngredientTypePersistencePort_GetByName_Call) Return(_a0 *domain.IngredientType, _a1 error) *IngredientTypePersistencePort_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientTypePersistencePort_GetByName_Call) RunAndReturn(run func(context.Context, string) (*domain.IngredientType, error)) *IngredientTypePersistencePort_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// NewIngredientTypePersistencePort creates a new instance of IngredientTypePersistencePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIngredientTypePersistencePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *IngredientTypePersistencePort {
	mock := &IngredientTypePersistencePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
