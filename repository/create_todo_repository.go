package repository

import (
	"context"
	"gorm.io/gorm"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
	"todo-clean/repository/model"
)

// TODO Add input context
func (repo newRepo) CreateTodoRepository(ctx context.Context, db *gorm.DB, todo domain.CreateTodoEntity) (resp *domain.CreateTodoEntity, err error) {

	input := model.TbTodoRepositoryModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
	}

	if err = db.WithContext(ctx).Create(input).Error; err != nil {

		return resp, error_lib.WrapError("repository", err)
	}

	resp = &domain.CreateTodoEntity{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   input.CreatedAt,
	}

	return resp, err
}
