package repository

import (
	"context"
	"gorm.io/gorm"
	"todo-clean/domain"
	"todo-clean/repository/model"
)

func (repo newRepo) UpdateTodoRepository(ctx context.Context, db *gorm.DB, todo domain.UpdateTodoEntity) (result domain.UpdateTodoEntity, err error) {

	input := model.TbTodoRepositoryModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		UpdatedAt:   todo.UpdatedAt,
	}
	
	response := db.WithContext(ctx).First(&input, input.ID)
	if len(response) == 0 {

	}

	return
}
