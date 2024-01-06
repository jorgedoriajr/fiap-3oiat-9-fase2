// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "hamburgueria/internal/modules/product/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductPersistencePort is an autogenerated mock type for the ProductPersistencePort type
type ProductPersistencePort struct {
	mock.Mock
}

// CheckProductExists provides a mock function with given fields: ctx, product
func (_m *ProductPersistencePort) CheckProductExists(ctx context.Context, product domain.Product) (*domain.Product, error) {
	ret := _m.Called(ctx, product)

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) (*domain.Product, error)); ok {
		return rf(ctx, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) *domain.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, product
func (_m *ProductPersistencePort) Create(ctx context.Context, product domain.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *ProductPersistencePort) GetAll(ctx context.Context) ([]domain.Product, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCategory provides a mock function with given fields: ctx, productID
func (_m *ProductPersistencePort) GetByCategory(ctx context.Context, productID string) ([]domain.Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, productID
func (_m *ProductPersistencePort) GetByID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNumber provides a mock function with given fields: ctx, productNumber
func (_m *ProductPersistencePort) GetByNumber(ctx context.Context, productNumber int) (*domain.Product, error) {
	ret := _m.Called(ctx, productNumber)

	var r0 *domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*domain.Product, error)); ok {
		return rf(ctx, productNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *domain.Product); ok {
		r0 = rf(ctx, productNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, productNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, product
func (_m *ProductPersistencePort) Update(ctx context.Context, product domain.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductPersistencePort creates a new instance of ProductPersistencePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductPersistencePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductPersistencePort {
	mock := &ProductPersistencePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
