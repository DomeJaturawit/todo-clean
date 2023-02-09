package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetTodoRepository(ctx context.Context, key uuid.UUID) (result []domain.GetTodoEntity, err error) {
	var todos []model.TbTodoRepositoryModel
	db := repo.db

	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), key)
	if err = db.WithContext(ctx).Find(&todos).Error; err != nil {

		return nil, errorLib.WrapError(common.ErrDBGetTodo.Error(), err)
	}

	if err := copier.Copy(&result, &todos); err != nil {
		return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
	}

	return result, nil
}
