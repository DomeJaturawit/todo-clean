package mockdata

import (
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
)

// Mock Create Repository
func CreateTodoEntityMockData() domain.CreateTodoEntity {
	return domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       NewString(),
		Description: NewString(),
		Status:      NewString(),
		CreatedAt:   time.Now(),
	}
}

func GetTodoEntityMockData() domain.GetTodoEntity {
	return domain.NewGetTodoEntity(
		mockID,
		mockTitle,
		mockDescription,
		mockStatus,
		mockCreatedAt,
	)
}

func UpdateTodoEntityMockData() domain.UpdateTodoEntity {
	return domain.NewUpdateTodoEntity(
		mockTitle,
		mockDescription,
		mockStatus,
		mockUpdateAt,
	)
}


func QueryUpdateTodoEntityMockData() domain.UpdateTodoQueryEntity {
	return domain.NewQueryUpdateTodoEntity(
		mockID,
	)
}

