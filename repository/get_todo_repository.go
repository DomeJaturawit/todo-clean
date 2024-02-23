package repository

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
	"todo-clean/repository/model"
)

func (repo newRepo) GetTodoRepository(ctx context.Context, key string) (result []domain.GetTodoEntity, err error) {
	var todos []model.TbTodoRepositoryGetModel
	db := repo.db
	if key != "" {
		db = getTodoQueryCondition(db, key)
		if err = db.WithContext(ctx).Find(&todos).Error; err != nil {

			return nil, errorLib.WrapError(common.ErrDBGetTodo.Error(), err)
		}

		if err := copier.Copy(&result, &todos); err != nil {
			return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
		}
	} else {
		if err = repo.db.WithContext(ctx).Find(&todos).Error; err != nil {
			return result, errorLib.WrapError(common.ErrDBGetTodo.Error(), err)
		}
		if err := copier.Copy(&result, &todos); err != nil {
			return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
		}
	}

	if len(result) == 0 {

		return nil, common.ErrDataNotFound
	}

	return result, nil
}

func getTodoQueryCondition(db *gorm.DB, key string) *gorm.DB {
	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), key)
	return db
}
