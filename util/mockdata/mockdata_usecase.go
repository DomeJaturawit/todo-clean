package mockdata

import (
	"github.com/google/uuid"
	"time"
	"todo-clean/domain"
)

const ()

var (
	mockTitle       = NewString()
	mockDescription = NewString()
	mockStatus      = NewString()
	mockID          = uuid.New()
	mockCreatedAt   = time.Now()
)

// Mock Request  usecase
func CreateTodoEntityRequestMockData() *domain.CreateTodoEntityRequest {
	return domain.NewCreateTodoEntityRequest(
		mockTitle,
		mockDescription,
		mockStatus,
	)
}

func CreateTodoUseCaseEntityMockData() *domain.CreateTodoEntity {
	return domain.NewCreateTodoEntity(
		mockID,
		mockTitle,
		mockDescription,
		mockStatus,
		mockCreatedAt,
	)
}
