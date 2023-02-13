// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "todo-clean/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// TodoUseCase is an autogenerated mock type for the TodoUseCase type
type TodoUseCase struct {
	mock.Mock
}

// CreateTodoUseCase provides a mock function with given fields: ctx, entity
func (_m *TodoUseCase) CreateTodoUseCase(ctx context.Context, entity domain.CreateTodoInputEntity) (*domain.CreateTodoEntity, error) {
	ret := _m.Called(ctx, entity)

	var r0 *domain.CreateTodoEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateTodoInputEntity) (*domain.CreateTodoEntity, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateTodoInputEntity) *domain.CreateTodoEntity); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CreateTodoEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateTodoInputEntity) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoUseCase provides a mock function with given fields: ctx, key
func (_m *TodoUseCase) GetTodoUseCase(ctx context.Context, key *uuid.UUID) ([]domain.GetTodoEntity, error) {
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

type mockConstructorTestingTNewTodoUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoUseCase creates a new instance of TodoUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoUseCase(t mockConstructorTestingTNewTodoUseCase) *TodoUseCase {
	mock := &TodoUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
