package domain

import (
	"context"
	"github.com/google/uuid"
)

type TodoUseCase interface {
	CreateTodoUseCase(ctx context.Context, entity CreateTodoInputEntity) (result *CreateTodoEntity, err error)
	GetTodoUseCase(ctx context.Context, key *uuid.UUID) (result []GetTodoEntity, err error)
}
