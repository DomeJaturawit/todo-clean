package usecase

import (
	"context"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) GetTodoUseCase(ctx context.Context, key string) (result []domain.GetTodoEntity, err error) {

	if key != "" {
		result, err = n.repo.GetTodoRepository(ctx, key)

		if err != nil {
			return nil, errorLib.WrapError(common.ErrGetTodo.Error(), err)
		}
	} else {
		result, err = n.repo.GetTodoRepository(ctx, "")
		if err != nil {
			return nil, errorLib.WrapError(common.ErrGetAllTodo.Error(), err)
		}
	}

	return
}
