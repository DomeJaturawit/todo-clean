package domain

import (
	"context"
	"gorm.io/gorm"
)

type TodoRepositoryInterface interface {
	GetTodoRepository(ctx context.Context) (resp []GetTodoEntity, err error)
	CreateTodoRepository(ctx context.Context, db *gorm.DB, todo CreateTodoEntity) (resp *CreateTodoEntity, err error)
	Begin() (tx *gorm.DB, err error)
	RollBack(tx *gorm.DB) (err error)
	Commit(tx *gorm.DB) (err error)
}
