package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
)

// TODO Use Func Wrapper Error!
func (n *newUseCase) CreateTodoUseCase(ctx context.Context, todo domain.CreateTodoEntityRequest) (result *domain.CreateTodoEntity, err error) {

	dbTx, err := n.repo.Begin(ctx)
	if err != nil {

		return nil, error_lib.WrapError("create todo usecase", err)
	}
	result, err = n.repo.CreateTodoRepository(dbTx, domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, error_lib.WrapError("create todo usecase", err)
	}
	dbTx, err = n.repo.Commit()
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {

			return nil, error_lib.WrapError("rollback create todo usecase", err)
		}
		return nil, error_lib.WrapError("commit create todo usecase", err)
	}

	return result, nil
}
