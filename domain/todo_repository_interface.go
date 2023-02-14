package domain

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodoRepository(ctx context.Context, key *uuid.UUID) (resp []GetTodoEntity, err error)
	CreateTodoRepository(ctx context.Context, db *gorm.DB, todo CreateTodoEntity) (resp *CreateTodoEntity, err error)
	UpdateTodoRepository(ctx context.Context, db *gorm.DB, query QueryUpdateTodoEntity, entity UpdateTodoEntity) (result *UpdateTodoEntity, err error)
	Begin() (tx *gorm.DB, err error)
	RollBack(tx *gorm.DB) (err error)
	Commit(tx *gorm.DB) (err error)
}
