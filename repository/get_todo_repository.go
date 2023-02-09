package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetTodoRepository(ctx context.Context) (resp []domain.GetTodoEntity, err error) {

	var todos []model.TbTodoRepositoryModel

	if err = repo.db.WithContext(ctx).Find(&todos).Error; err != nil {

		return resp, error_lib.WrapError(common.ErrDBGetTodo.Error(), err)
	}

	//resp = make([]domain.GetTodoEntity, len(todos))
	//for i, todo := range todos {
	//	resp[i] = domain.GetTodoEntity{
	//		ID:          todo.ID,
	//		Title:       todo.Title,
	//		Description: todo.Description,
	//		Status:      todo.Status,
	//		CreatedAt:   todo.CreatedAt,
	//	}
	//}

	if err := copier.Copy(&resp, &todos); err != nil {
		return nil, err //TODO: wrap Error
	}

	return resp, nil
}
