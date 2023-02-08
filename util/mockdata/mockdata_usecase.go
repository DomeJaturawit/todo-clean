package mockdata

import (
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
)

// Mock Request  usecase
func CreateTodoEntityRequestMockData() domain.CreateTodoEntityRequest {
	return domain.CreateTodoEntityRequest{
		Title:       NewString(),
		Description: NewString(),
		Status:      NewString(),
	}
}

func CreateTodoUseCaseEntityMockData(input domain.CreateTodoEntityRequest) domain.CreateTodoEntity {
	return domain.CreateTodoEntity{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   time.Now(),
	}
}
