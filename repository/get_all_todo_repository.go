package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetAllTodoRepository(ctx context.Context) (result []domain.GetTodoEntity, err error) {

	var todos []model.TbTodoRepositoryModel

	if err = repo.db.WithContext(ctx).Find(&todos).Error; err != nil {
		return result, errorLib.WrapError(common.ErrDBGetTodo.Error(), err)
	}

	//I want keep this comment, i think it's help when i comeback to read,it will explain copier.Copy below
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

	if err := copier.Copy(&result, &todos); err != nil {
		return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
	}

	return result, nil
}
