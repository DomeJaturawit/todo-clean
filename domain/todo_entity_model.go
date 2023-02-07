package domain

import (
	"github.com/google/uuid"
	"time"
)

type CreateTodoEntity struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:createdAt"`
}

type CreateTodoEntityRequest struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}
