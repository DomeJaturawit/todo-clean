package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
)

//	type testS struct {
//		Title string
//	}
//
// func (n *newUseCase)test()(){
//
//	test := testS{}
//
//	result,err := n.CreateTodoUseCase(domain.CreateTodoEntity{
//
//		Title:       test.Title,
//
//	})
//
// }
// TODO Use Func Wrapper Error!
func (n *newUseCase) CreateTodoUseCase(ctx context.Context, todo domain.CreateTodoEntityRequest) (result *domain.CreateTodoEntity, err error) {

	dbTx, err := n.repo.Begin(ctx)
	if err != nil {

		return nil, fmt.Errorf("create todo usecase error: %w", err)
	}
	result, err = n.repo.CreateTodoRepository(dbTx, domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("create todo usecase error: %w", err)
	}
	dbTx, err = n.repo.Commit()
	if err != nil {
		if rollbackErr := n.repo.RollBack(dbTx); rollbackErr != nil {

			return nil, fmt.Errorf("create todo usecase error: %w", rollbackErr)
		}
		return nil, fmt.Errorf("create todo usecase error: %w", err)
	}

	return result, nil
}
