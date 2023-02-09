package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetTodoRepository(ctx context.Context, key uuid.UUID) (resp []domain.GetTodoEntity, err error) {
	var todos []model.TbTodoRepositoryModel
	db := repo.db

	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), key)
	if err = db.WithContext(ctx).Find(&todos).Error; err != nil {

		return nil, error_lib.WrapError(common.ErrDBGetTodo.Error(), err)
	}

	if err := copier.Copy(&resp, &todos); err != nil {
		return nil, error_lib.WrapError(common.ErrCopierCopy.Error(), err)
	}

	return resp, nil
}
