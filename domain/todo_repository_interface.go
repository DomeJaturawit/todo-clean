package domain

import "gorm.io/gorm"

type TodoRepositoryInterface interface {
	CreateTodoRepository(db *gorm.DB, todo CreateTodoEntity) (resp *CreateTodoEntity, err error)
}
