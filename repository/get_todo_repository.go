package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetTodoRepository(ctx context.Context, key uuid.UUID) (resp []domain.GetTodoEntity, err error) {
	var todos []model.TbTodoRepositoryModel
	db := repo.db

	db = getTodoQueryCondition(db, key)
	if err = db.WithContext(ctx).Find(&todos).Error; err != nil {

		return nil, errorLib.WrapError(common.ErrDBGetTodo.Error(), err)
	}

	if err := copier.Copy(&resp, &todos); err != nil {
		return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
	}

	return resp, err
}

func getTodoQueryCondition(db *gorm.DB, key uuid.UUID) *gorm.DB {
	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), key)
	return db
}
