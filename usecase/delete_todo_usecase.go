package usecase

import (
	"context"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) DeleteTodoUseCase(ctx context.Context, queryEntity domain.DeleteTodoQueryEntity) (result *domain.DeleteTodoQueryEntity, err error) {
	dbTx, err := n.repo.Begin()
	if err != nil {
		return nil, errorLib.WrapError(common.ErrBeginTodo.Error(), err)
	}
	result, err = n.repo.DeleteTodoRepository(ctx, dbTx, queryEntity)
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {
			return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
		}
		return nil, errorLib.WrapError(common.ErrUseCaseDeleteTodo.Error(), err)
	}

	err = n.repo.Commit(dbTx)
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {
			return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
		}
		return nil, errorLib.WrapError(common.ErrCommitTodo.Error(), err)
	}
	return
}
