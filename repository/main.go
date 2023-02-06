package repository

import (
	"gorm.io/gorm"
	"todo-clean/domain"
)

type newRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.TodoRepositoryInterface {
	return &newRepository{db: db}

}
