package model

import (
	"github.com/google/uuid"
	"time"
	"todo-clean/common"
)

type TbTodoTableSchema struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
type TbTodoRepositoryCreateModel struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TbTodoRepositoryUpdateModel struct {
	Title       *string    `json:"title" gorm:"column:title"`
	Description *string    `json:"description" gorm:"column:description"`
	Status      *string    `json:"status" gorm:"column:status"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type TbTodoRepositoryGetModel struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TbTodoRepositoryDeleteModel struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func (TbTodoRepositoryCreateModel) TableName() string {
	return common.TodoTable
}
func (TbTodoRepositoryUpdateModel) TableName() string {
	return common.TodoTable
}
func (TbTodoRepositoryGetModel) TableName() string {
	return common.TodoTable
}

func (TbTodoRepositoryDeleteModel) TableName() string {
	return common.TodoTable
}

func (TbTodoTableSchema) TableName() string {
	return common.TodoTable
}
