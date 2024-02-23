package domain

import (
	"context"
)

type TodoUseCase interface {
	CreateTodoUseCase(ctx context.Context, entity CreateTodoInputEntity) (result *CreateTodoEntity, err error)
	GetTodoUseCase(ctx context.Context, key string) (result []GetTodoEntity, err error)
	DeleteTodoUseCase(ctx context.Context, queryEntity DeleteTodoQueryEntity) (result *DeleteTodoQueryEntity, err error)
	UpdateTodoUseCase(ctx context.Context, queryEntity UpdateTodoQueryEntity, entity *UpdateTodoEntity) (result *UpdateTodoEntity, err error)
}
