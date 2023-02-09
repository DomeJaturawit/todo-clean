// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "todo-clean/domain"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// TodoRepositoryInterface is an autogenerated mock type for the TodoRepositoryInterface type
type TodoRepositoryInterface struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *TodoRepositoryInterface) Begin() (*gorm.DB, error) {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: tx
func (_m *TodoRepositoryInterface) Commit(tx *gorm.DB) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTodoRepository provides a mock function with given fields: ctx, db, todo
func (_m *TodoRepositoryInterface) CreateTodoRepository(ctx context.Context, db *gorm.DB, todo domain.CreateTodoEntity) (*domain.CreateTodoEntity, error) {
	ret := _m.Called(ctx, db, todo)

	var r0 *domain.CreateTodoEntity
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.CreateTodoEntity) *domain.CreateTodoEntity); ok {
		r0 = rf(ctx, db, todo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CreateTodoEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, domain.CreateTodoEntity) error); ok {
		r1 = rf(ctx, db, todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTodoRepository provides a mock function with given fields: ctx
func (_m *TodoRepositoryInterface) GetAllTodoRepository(ctx context.Context) ([]domain.GetTodoEntity, error) {
	ret := _m.Called(ctx)

	var r0 []domain.GetTodoEntity
	if rf, ok := ret.Get(0).(func(context.Context) []domain.GetTodoEntity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.GetTodoEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollBack provides a mock function with given fields: tx
func (_m *TodoRepositoryInterface) RollBack(tx *gorm.DB) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTodoRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepositoryInterface creates a new instance of TodoRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepositoryInterface(t mockConstructorTestingTNewTodoRepositoryInterface) *TodoRepositoryInterface {
	mock := &TodoRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
