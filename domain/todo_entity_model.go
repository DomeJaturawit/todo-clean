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

func NewCreateTodoEntity(ID uuid.UUID, title string, description string, status string, createdAt time.Time) *CreateTodoEntity {
	return &CreateTodoEntity{ID: ID, Title: title, Description: description, Status: status, CreatedAt: createdAt}
}

type CreateTodoEntityRequest struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func NewCreateTodoEntityRequest(title string, description string, status string) *CreateTodoEntityRequest {
	return &CreateTodoEntityRequest{Title: title, Description: description, Status: status}
}

type GetTodoEntity struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:createdAt"`
}

func NewGetTodoEntity(ID uuid.UUID, title string, description string, status string, createdAt time.Time) *GetTodoEntity {
	return &GetTodoEntity{ID: ID, Title: title, Description: description, Status: status, CreatedAt: createdAt}
}
