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

type CreateTodoInputEntity struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Status      string `json:"status" gorm:"column:status"`
}

func NewCreateTodoInputEntity(title string, description string, status string) *CreateTodoInputEntity {
	return &CreateTodoInputEntity{Title: title, Description: description, Status: status}
}

type GetTodoEntity struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:createdAt"`
}

type GetOneTodoEntity struct {
	ID uuid.UUID `json:"id" gorm:"column:id"`
}

func NewGetTodoEntity(ID uuid.UUID, title string, description string, status string, createdAt time.Time) GetTodoEntity {
	return GetTodoEntity{ID: ID, Title: title, Description: description, Status: status, CreatedAt: createdAt}
}

type UpdateTodoEntity struct {
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Status      string    `json:"status" gorm:"column:status"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func NewUpdateTodoEntity(title string, description string, status string, updatedAt time.Time) UpdateTodoEntity {
	return UpdateTodoEntity{Title: title, Description: description, Status: status, UpdatedAt: updatedAt}
}

type QueryUpdateTodoEntity struct {
	ID uuid.UUID `json:"id" gorm:"primary_key"`
}

func NewQueryUpdateTodoEntity(ID uuid.UUID) QueryUpdateTodoEntity {
	return QueryUpdateTodoEntity{ID: ID}
}
