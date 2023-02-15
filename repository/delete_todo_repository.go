package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"todo-clean/common"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
	"todo-clean/repository/model"
)

func (repo newRepo) DeleteTodoRepository(ctx context.Context, db *gorm.DB, queryEntity domain.DeleteTodoQueryEntity) (result *domain.DeleteTodoQueryEntity, err error) {
	todoModel := new(model.TbTodoRepositoryDeleteModel)
	db = db.Where(fmt.Sprintf(`%s = ?`, common.TodoIDCol), queryEntity.ID)
	if err := db.WithContext(ctx).Delete(todoModel).Error; err != nil {
		return nil, errorLib.WrapError(common.ErrDBDeleteTodo.Error(), err)
	}

	//result = &domain.DeleteTodoEntity{
	//	Title:       todoModel.Title,
	//	Description: todoModel.Description,
	//	Status:      todoModel.Status,
	//	DeletedAt:   todoModel.DeletedAt,
	//}

	return
}
