package domain

import (
	"context"
)

type TodoUseCaseInterface interface {
	CreateTodoUseCase(ctx context.Context, todo CreateTodoEntityRequest) (result *CreateTodoEntity, err error)
}
