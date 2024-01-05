// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	result "hamburgueria/internal/modules/ingredient/usecase/result"
)

// FindIngredientTypeUseCasePort is an autogenerated mock type for the FindIngredientTypeUseCasePort type
type FindIngredientTypeUseCasePort struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: ctx
func (_m *FindIngredientTypeUseCasePort) FindAll(ctx context.Context) ([]result.IngredientTypeResult, error) {
	ret := _m.Called(ctx)

	var r0 []result.IngredientTypeResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]result.IngredientTypeResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []result.IngredientTypeResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]result.IngredientTypeResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFindIngredientTypeUseCasePort creates a new instance of FindIngredientTypeUseCasePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFindIngredientTypeUseCasePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *FindIngredientTypeUseCasePort {
	mock := &FindIngredientTypeUseCasePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
