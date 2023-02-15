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

func (repo newRepo) UpdateTodoRepository(ctx context.Context, db *gorm.DB, queryEntity domain.UpdateTodoQueryEntity, entity domain.UpdateTodoEntity) (result *domain.UpdateTodoEntity, err error) {

	todoModel := new(model.TbTodoRepositoryUpdateModel)

	if err := copier.Copy(todoModel, &entity); err != nil {
		//TODO: Check wrapper
		return nil, errorLib.WrapError(common.ErrCopierCopy.Error(), err)
	}

	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), queryEntity.ID)

	if err := db.WithContext(ctx).Updates(todoModel).Error; err != nil {
		return nil, errorLib.WrapError(common.ErrDBUpdateTodo.Error(), err)
	}

	return &entity, err
}
