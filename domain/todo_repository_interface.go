package domain

import (
	"context"
	"gorm.io/gorm"
)

type TodoRepositoryInterface interface {
	CreateTodoRepository(ctx context.Context, db *gorm.DB, todo CreateTodoEntity) (resp *CreateTodoEntity, err error)
	Begin(ctx context.Context) (tx *gorm.DB, err error)
	RollBack(tx *gorm.DB) (err error)
	Commit() (tx *gorm.DB, err error)
}
