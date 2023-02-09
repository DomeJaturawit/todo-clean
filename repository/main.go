package repository

import (
	"gorm.io/gorm"
	"todo-clean/domain"
)

type newRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.TodoRepositoryInterface {
	return &newRepo{db: db}

}
