package model

import (
	"github.com/google/uuid"
	"time"
	"todo-clean/common"
)

type TbTodoRepositoryModel struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func (TbTodoRepositoryModel) TableName() string {
	return common.TodoTable
}
