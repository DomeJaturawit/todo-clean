package usecase

import (
	"context"
	"github.com/google/uuid"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) GetTodoUseCase(ctx context.Context, key uuid.UUID) (result []domain.GetTodoEntity, err error) {

	result, err = n.repo.GetTodoRepository(ctx, key)
	if err != nil {
		return nil, errorLib.WrapError(common.ErrGetAllTodo.Error(), err)
	}

	return
}
