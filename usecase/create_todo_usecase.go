package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
)

func (n *newUseCase) CreateTodoUseCase(ctx context.Context, todo domain.CreateTodoEntityRequest) (result *domain.CreateTodoEntity, err error) {

	dbTx, err := n.repo.Begin(ctx)
	if err != nil {

		return nil, error_lib.WrapError(common.ErrBeginCreateTodo.Error(), err)
	}
	result, err = n.repo.CreateTodoRepository(ctx, dbTx, domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, error_lib.WrapError(common.ErrUseCaseCreateTodo.Error(), err)
	}

	err = n.repo.Commit(dbTx)
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {

			return nil, error_lib.WrapError(common.ErrRollbackCreateTodo.Error(), rollbackErr)
		}
		return nil, error_lib.WrapError(common.ErrCommitCreateTodo.Error(), err)
	}

	return result, nil
}
