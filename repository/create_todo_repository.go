package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"todo-clean/domain"
	"todo-clean/repository/model"
)

// TODO add input context
func (repo newRepo) CreateTodoRepository(db *gorm.DB, todo domain.CreateTodoEntity) (resp *domain.CreateTodoEntity, err error) {

	input := model.TbTodoRepositoryModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
	}

	if err = db.Create(input).Error; err != nil {
		err = errors.New("RepositoryError")
		log.Println("ERR >>>>", err.Error())

		return resp, err
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
