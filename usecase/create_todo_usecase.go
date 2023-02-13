package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

func (n *newUseCase) CreateTodoUseCase(ctx context.Context, entity domain.CreateTodoInputEntity) (result *domain.CreateTodoEntity, err error) {

	dbTx, err := n.repo.Begin()
	if err != nil {

		return nil, errorLib.WrapError(common.ErrBeginTodo.Error(), err)
	}
	result, err = n.repo.CreateTodoRepository(ctx, dbTx, domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       entity.Title,
		Description: entity.Description,
		Status:      entity.Status,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, errorLib.WrapError(common.ErrUseCaseCreateTodo.Error(), err)
	}

	err = n.repo.Commit(dbTx)
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {

			return nil, errorLib.WrapError(common.ErrRollbackTodo.Error(), rollbackErr)
		}
		return nil, errorLib.WrapError(common.ErrCommitTodo.Error(), err)
	}

	return result, nil
}
