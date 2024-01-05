// Code generated by mockery v2.32.3. DO NOT EDIT.

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

// GetAll provides a mock function with given fields: ctx
func (_m *IngredientTypePersistencePort) GetAll(ctx context.Context) ([]domain.IngredientType, error) {
	ret := _m.Called(ctx)

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

// GetByName provides a mock function with given fields: ctx, name
func (_m *IngredientTypePersistencePort) GetByName(ctx context.Context, name string) (*domain.IngredientType, error) {
	ret := _m.Called(ctx, name)

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