package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
)

//type testS struct {
//	Title string
//}
//func (n *newUseCase)test()(){
//
//	test := testS{}
//
//	result,err := n.CreateTodoUseCase(domain.CreateTodoEntity{
//
//		Title:       test.Title,
//
//	})
//
//}

const (
	StatusTypeTodo = "TODO"
)

func (n *newUseCase) CreateTodoUseCase(ctx context.Context, todo domain.CreateTodoEntityRequest) (result *domain.CreateTodoEntity, err error) {

	dbTx, err := n.repo.Begin(ctx)
	if err != nil {

		err = errors.New("Begin UseCase Error")
		return nil, err
	}
	result, err = n.repo.CreateTodoRepository(dbTx, domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		//rollback
		_, err := n.repo.RollBack()
		err = errors.New("UseCaseError")
		return nil, err
	}
	dbTx, err = n.repo.Commit()
	//do commit

	return result, err
}
