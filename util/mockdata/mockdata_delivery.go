package mockdata

import (
	"todo-clean/delivery/model"
)

func CreateTodoDeliveryRequestMockData() *model.CreateTodoDeliveryRequest {
	return model.NewCreateTodoDeliveryRequest(
		mockTitle,
		mockDescription,
		mockStatus,
	)
}
