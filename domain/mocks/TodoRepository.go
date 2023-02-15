// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "todo-clean/domain"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *TodoRepository) Begin() (*gorm.DB, error) {
	ret := _m.Called()

	var r0 *gorm.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gorm.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: tx
func (_m *TodoRepository) Commit(tx *gorm.DB) error {
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
func (_m *TodoRepository) CreateTodoRepository(ctx context.Context, db *gorm.DB, todo domain.CreateTodoEntity) (*domain.CreateTodoEntity, error) {
	ret := _m.Called(ctx, db, todo)

	var r0 *domain.CreateTodoEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.CreateTodoEntity) (*domain.CreateTodoEntity, error)); ok {
		return rf(ctx, db, todo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.CreateTodoEntity) *domain.CreateTodoEntity); ok {
		r0 = rf(ctx, db, todo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CreateTodoEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, domain.CreateTodoEntity) error); ok {
		r1 = rf(ctx, db, todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoRepository provides a mock function with given fields: ctx, key
func (_m *TodoRepository) GetTodoRepository(ctx context.Context, key *uuid.UUID) ([]domain.GetTodoEntity, error) {
	ret := _m.Called(ctx, key)

	var r0 []domain.GetTodoEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID) ([]domain.GetTodoEntity, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *uuid.UUID) []domain.GetTodoEntity); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.GetTodoEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *uuid.UUID) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollBack provides a mock function with given fields: tx
func (_m *TodoRepository) RollBack(tx *gorm.DB) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTodoRepository provides a mock function with given fields: ctx, db, query, entity
func (_m *TodoRepository) UpdateTodoRepository(ctx context.Context, db *gorm.DB, query domain.UpdateTodoQueryEntity, entity domain.UpdateTodoEntity) (*domain.UpdateTodoEntity, error) {
	ret := _m.Called(ctx, db, query, entity)

	var r0 *domain.UpdateTodoEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.UpdateTodoQueryEntity, domain.UpdateTodoEntity) (*domain.UpdateTodoEntity, error)); ok {
		return rf(ctx, db, query, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.UpdateTodoQueryEntity, domain.UpdateTodoEntity) *domain.UpdateTodoEntity); ok {
		r0 = rf(ctx, db, query, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.UpdateTodoEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, domain.UpdateTodoQueryEntity, domain.UpdateTodoEntity) error); ok {
		r1 = rf(ctx, db, query, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepository(t mockConstructorTestingTNewTodoRepository) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
