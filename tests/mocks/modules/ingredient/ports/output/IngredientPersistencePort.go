// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "hamburgueria/internal/modules/ingredient/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IngredientPersistencePort is an autogenerated mock type for the IngredientPersistencePort type
type IngredientPersistencePort struct {
	mock.Mock
}

type IngredientPersistencePort_Expecter struct {
	mock *mock.Mock
}

func (_m *IngredientPersistencePort) EXPECT() *IngredientPersistencePort_Expecter {
	return &IngredientPersistencePort_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, ingredient
func (_m *IngredientPersistencePort) Create(ctx context.Context, ingredient domain.Ingredient) error {
	ret := _m.Called(ctx, ingredient)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Ingredient) error); ok {
		r0 = rf(ctx, ingredient)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IngredientPersistencePort_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type IngredientPersistencePort_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - ingredient domain.Ingredient
func (_e *IngredientPersistencePort_Expecter) Create(ctx interface{}, ingredient interface{}) *IngredientPersistencePort_Create_Call {
	return &IngredientPersistencePort_Create_Call{Call: _e.mock.On("Create", ctx, ingredient)}
}

func (_c *IngredientPersistencePort_Create_Call) Run(run func(ctx context.Context, ingredient domain.Ingredient)) *IngredientPersistencePort_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Ingredient))
	})
	return _c
}

func (_c *IngredientPersistencePort_Create_Call) Return(_a0 error) *IngredientPersistencePort_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IngredientPersistencePort_Create_Call) RunAndReturn(run func(context.Context, domain.Ingredient) error) *IngredientPersistencePort_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx
func (_m *IngredientPersistencePort) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Ingredient, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Ingredient); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientPersistencePort_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type IngredientPersistencePort_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *IngredientPersistencePort_Expecter) GetAll(ctx interface{}) *IngredientPersistencePort_GetAll_Call {
	return &IngredientPersistencePort_GetAll_Call{Call: _e.mock.On("GetAll", ctx)}
}

func (_c *IngredientPersistencePort_GetAll_Call) Run(run func(ctx context.Context)) *IngredientPersistencePort_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *IngredientPersistencePort_GetAll_Call) Return(_a0 []domain.Ingredient, _a1 error) *IngredientPersistencePort_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientPersistencePort_GetAll_Call) RunAndReturn(run func(context.Context) ([]domain.Ingredient, error)) *IngredientPersistencePort_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, ingredientID
func (_m *IngredientPersistencePort) GetByID(ctx context.Context, ingredientID uuid.UUID) (*domain.Ingredient, error) {
	ret := _m.Called(ctx, ingredientID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Ingredient, error)); ok {
		return rf(ctx, ingredientID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Ingredient); ok {
		r0 = rf(ctx, ingredientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, ingredientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientPersistencePort_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type IngredientPersistencePort_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - ingredientID uuid.UUID
func (_e *IngredientPersistencePort_Expecter) GetByID(ctx interface{}, ingredientID interface{}) *IngredientPersistencePort_GetByID_Call {
	return &IngredientPersistencePort_GetByID_Call{Call: _e.mock.On("GetByID", ctx, ingredientID)}
}

func (_c *IngredientPersistencePort_GetByID_Call) Run(run func(ctx context.Context, ingredientID uuid.UUID)) *IngredientPersistencePort_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *IngredientPersistencePort_GetByID_Call) Return(_a0 *domain.Ingredient, _a1 error) *IngredientPersistencePort_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientPersistencePort_GetByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*domain.Ingredient, error)) *IngredientPersistencePort_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByNumber provides a mock function with given fields: ctx, number
func (_m *IngredientPersistencePort) GetByNumber(ctx context.Context, number int) (*domain.Ingredient, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for GetByNumber")
	}

	var r0 *domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*domain.Ingredient, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Ingredient); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientPersistencePort_GetByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByNumber'
type IngredientPersistencePort_GetByNumber_Call struct {
	*mock.Call
}

// GetByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number int
func (_e *IngredientPersistencePort_Expecter) GetByNumber(ctx interface{}, number interface{}) *IngredientPersistencePort_GetByNumber_Call {
	return &IngredientPersistencePort_GetByNumber_Call{Call: _e.mock.On("GetByNumber", ctx, number)}
}

func (_c *IngredientPersistencePort_GetByNumber_Call) Run(run func(ctx context.Context, number int)) *IngredientPersistencePort_GetByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *IngredientPersistencePort_GetByNumber_Call) Return(_a0 *domain.Ingredient, _a1 error) *IngredientPersistencePort_GetByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientPersistencePort_GetByNumber_Call) RunAndReturn(run func(context.Context, int) (*domain.Ingredient, error)) *IngredientPersistencePort_GetByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetByType provides a mock function with given fields: ctx, ingredientType
func (_m *IngredientPersistencePort) GetByType(ctx context.Context, ingredientType string) ([]domain.Ingredient, error) {
	ret := _m.Called(ctx, ingredientType)

	if len(ret) == 0 {
		panic("no return value specified for GetByType")
	}

	var r0 []domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Ingredient, error)); ok {
		return rf(ctx, ingredientType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Ingredient); ok {
		r0 = rf(ctx, ingredientType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ingredientType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IngredientPersistencePort_GetByType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByType'
type IngredientPersistencePort_GetByType_Call struct {
	*mock.Call
}

// GetByType is a helper method to define mock.On call
//   - ctx context.Context
//   - ingredientType string
func (_e *IngredientPersistencePort_Expecter) GetByType(ctx interface{}, ingredientType interface{}) *IngredientPersistencePort_GetByType_Call {
	return &IngredientPersistencePort_GetByType_Call{Call: _e.mock.On("GetByType", ctx, ingredientType)}
}

func (_c *IngredientPersistencePort_GetByType_Call) Run(run func(ctx context.Context, ingredientType string)) *IngredientPersistencePort_GetByType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IngredientPersistencePort_GetByType_Call) Return(_a0 []domain.Ingredient, _a1 error) *IngredientPersistencePort_GetByType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IngredientPersistencePort_GetByType_Call) RunAndReturn(run func(context.Context, string) ([]domain.Ingredient, error)) *IngredientPersistencePort_GetByType_Call {
	_c.Call.Return(run)
	return _c
}

// NewIngredientPersistencePort creates a new instance of IngredientPersistencePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIngredientPersistencePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *IngredientPersistencePort {
	mock := &IngredientPersistencePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
