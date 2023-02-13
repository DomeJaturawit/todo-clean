package usecase

import (
	"context"
	"log"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) GetAllTodoUseCase(ctx context.Context) (result []domain.GetTodoEntity, err error) {
	result, err = n.repo.GetAllTodoRepository(ctx)
	log.Println("result =>>", result)
	if err != nil {
		return nil, errorLib.WrapError(common.ErrGetAllTodo.Error(), err)
	}

	return result, err

}
