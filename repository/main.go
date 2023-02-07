package repository

import (
	"gorm.io/gorm"
	"todo-clean/domain"
)

type NewRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.TodoRepositoryInterface {
	return &NewRepo{db: db}

}
