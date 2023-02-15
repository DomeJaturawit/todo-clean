package usecase

import (
	"context"
	"time"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) UpdateTodoUseCase(ctx context.Context, queryEntity domain.UpdateTodoQueryEntity, entity *domain.UpdateTodoEntity) (result *domain.UpdateTodoEntity, err error) {

	response, err := n.repo.GetTodoRepository(ctx, &queryEntity.ID)
	if len(response) == 0 {
		return nil, errorLib.WrapError(common.ErrDataNotFound.Error(), err)
	}

	dbTx, err := n.repo.Begin()
	if err != nil {
		return nil, errorLib.WrapError(common.ErrBeginTodo.Error(), err)
	}
	result, err = n.repo.UpdateTodoRepository(ctx, dbTx, queryEntity, domain.UpdateTodoEntity{
		Title:       entity.Title,
		Description: entity.Description,
		Status:      entity.Status,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {
			return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
		}
		return nil, errorLib.WrapError(common.ErrUseCaseUpdateTodo.Error(), err)
	}
	err = n.repo.Commit(dbTx)
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {
			return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
		}
		return nil, errorLib.WrapError(common.ErrCommitTodo.Error(), err)
	}

	return result, err
}
